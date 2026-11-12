// Command pluginindex generates and verifies the marketplace index.json for a
// new-api plugin repository laid out as plugins/<kind>/<key>/<version>/plugin.js.
//
// The index is a derived artifact: every display field comes from the compiled
// plugin meta, never from hand-edited JSON, so the index can repeat meta
// content without ever drifting from it. Admission on the gateway still trusts
// only the compiled meta; the index exists for the three things a single
// plugin file cannot carry — the repository catalog, the cross-version table,
// and the source sha256.
//
// Usage:
//
//	pluginindex generate <marketplace-root>
//	pluginindex check <marketplace-root>
//
// generate writes index.json at the repository root. check recompiles
// everything and fails when the committed index.json differs byte-for-byte.
package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/QuantumNous/new-api/pkg/jsplugin"
	"gopkg.in/yaml.v3"
)

// kindDirs maps a repository layout directory to the plugin kind it declares.
// The directory is organization, the compiled meta is the declaration; task
// plugins never write meta.kind, which is exactly the "tasks" directory.
var kindDirs = map[string]string{"tasks": "task"}

type indexVersion struct {
	Version       string        `json:"version"`
	Path          string        `json:"path"`
	Sha256        string        `json:"sha256"`
	MinApiVersion int           `json:"minApiVersion"`
	Kind          string        `json:"kind"`
	AllowedHosts  []string      `json:"allowedHosts,omitempty"`
	BaseUrl       string        `json:"baseUrl,omitempty"`
	Auth          string        `json:"auth,omitempty"`
	Readme        *indexDocFile `json:"readme,omitempty"`
	Changelog     *indexDocFile `json:"changelog,omitempty"`
}

// indexDocFile points at a version's usage guide (README.md) or release notes
// (CHANGELOG.md) and the translations its front matter lists. Paths are
// repository-relative, like the version path.
type indexDocFile struct {
	Path         string                       `json:"path"`
	Sha256       string                       `json:"sha256"`
	Translations map[string]indexDocFileEntry `json:"translations,omitempty"`
}

type indexDocFileEntry struct {
	Path   string `json:"path"`
	Sha256 string `json:"sha256"`
}

// maxDocTranslations bounds the translations map of a README or CHANGELOG.
const maxDocTranslations = 10

// localeTagPattern matches the locale tags plugin metadata accepts.
var localeTagPattern = regexp.MustCompile(`^[a-zA-Z]{2,3}(-[a-zA-Z0-9]{2,8})*$`)

// pluginReadmePattern matches the plugin-level usage guide and its translations.
var pluginReadmePattern = regexp.MustCompile(`^README(\.[a-zA-Z]{2,3}(-[a-zA-Z0-9]{2,8})*)?\.md$`)

// indexIconFile points at the sidecar logo (icon.svg / icon.png) stored once
// per plugin key, next to the version directories. Gateways fetch it from the
// index origin and store it apart from the plugin source.
type indexIconFile struct {
	Path   string `json:"path"`
	Sha256 string `json:"sha256"`
}

type indexPlugin struct {
	displayVersion string                               `json:"-"`
	Key            string                               `json:"key"`
	Name           string                               `json:"name"`
	Icon           string                               `json:"icon,omitempty"`
	IconFile       *indexIconFile                       `json:"iconFile,omitempty"`
	Website        string                               `json:"website,omitempty"`
	SortPriority   int                                  `json:"sortPriority,omitempty"`
	Description    jsplugin.LocalizedText               `json:"description,omitempty"`
	Routes         []jsplugin.Route                     `json:"routes"`
	Protocols      []jsplugin.ProtocolClaim             `json:"protocols,omitempty"`
	UsageSchema    map[string]jsplugin.UsageFieldSchema `json:"usageSchema"`
	UsageExamples  []jsplugin.UsageExample              `json:"usageExamples"`
	UsageProfiles  []jsplugin.UsageProfile              `json:"usageProfiles"`
	ChannelTypes   []int                                `json:"channelTypes,omitempty"`
	Models         []string                             `json:"models,omitempty"`
	Task           bool                                 `json:"task"`
	Provider       bool                                 `json:"provider"`
	Transform      bool                                 `json:"transform"`
	Latest         string                               `json:"latest"`
	Versions       []indexVersion                       `json:"versions"`
}

type indexFile struct {
	IndexVersion int           `json:"indexVersion"`
	Name         string        `json:"name"`
	Plugins      []indexPlugin `json:"plugins"`
}

func main() {
	if len(os.Args) != 3 || (os.Args[1] != "generate" && os.Args[1] != "check") {
		fmt.Fprintln(os.Stderr, "usage: pluginindex generate|check <marketplace-root>")
		os.Exit(2)
	}
	mode, root := os.Args[1], os.Args[2]
	generated, err := buildIndex(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "pluginindex:", err)
		os.Exit(1)
	}
	indexPath := filepath.Join(root, "index.json")
	if mode == "generate" {
		if err = os.WriteFile(indexPath, generated, 0o644); err != nil {
			fmt.Fprintln(os.Stderr, "pluginindex:", err)
			os.Exit(1)
		}
		fmt.Println("wrote", indexPath)
		return
	}
	committed, err := os.ReadFile(indexPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "pluginindex:", err)
		os.Exit(1)
	}
	if !bytes.Equal(committed, generated) {
		fmt.Fprintln(os.Stderr, "pluginindex: index.json is stale; run pluginindex generate")
		os.Exit(1)
	}
	fmt.Println("index.json is up to date")
}

func buildIndex(root string) ([]byte, error) {
	pluginsDir := filepath.Join(root, "plugins")
	kindEntries, err := os.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", pluginsDir, err)
	}
	byKey := make(map[string]*indexPlugin)
	for _, kindEntry := range kindEntries {
		if !kindEntry.IsDir() {
			return nil, fmt.Errorf("unexpected non-directory %s under plugins/", kindEntry.Name())
		}
		kind, known := kindDirs[kindEntry.Name()]
		if !known {
			return nil, fmt.Errorf("unknown plugin kind directory %q", kindEntry.Name())
		}
		keyEntries, readErr := os.ReadDir(filepath.Join(pluginsDir, kindEntry.Name()))
		if readErr != nil {
			return nil, readErr
		}
		for _, keyEntry := range keyEntries {
			if !keyEntry.IsDir() {
				return nil, fmt.Errorf("unexpected non-directory %s under plugins/%s/", keyEntry.Name(), kindEntry.Name())
			}
			if err = collectPluginVersions(root, kindEntry.Name(), kind, keyEntry.Name(), byKey); err != nil {
				return nil, err
			}
		}
	}

	index := indexFile{IndexVersion: 1, Name: "New API Official Plugins", Plugins: make([]indexPlugin, 0, len(byKey))}
	for _, plugin := range byKey {
		sort.Slice(plugin.Versions, func(i, j int) bool {
			return semverLess(plugin.Versions[j].Version, plugin.Versions[i].Version)
		})
		plugin.Latest = plugin.Versions[0].Version
		index.Plugins = append(index.Plugins, *plugin)
	}
	sort.Slice(index.Plugins, func(i, j int) bool { return index.Plugins[i].Key < index.Plugins[j].Key })

	encoded, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(encoded, '\n'), nil
}

func collectPluginVersions(root, kindDir, kind, key string, byKey map[string]*indexPlugin) error {
	keyDir := filepath.Join(root, "plugins", kindDir, key)
	versionEntries, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	// The plugin-level README serves every version without its own README.
	pluginReadme, err := loadDocFile(root, filepath.ToSlash(filepath.Join("plugins", kindDir, key)), "README")
	if err != nil {
		return err
	}
	for _, versionEntry := range versionEntries {
		if !versionEntry.IsDir() {
			// The sidecar logo and the usage guide live once per plugin key,
			// beside the version directories. The logo is picked up after the
			// versions below.
			name := versionEntry.Name()
			if name == "icon.svg" || name == "icon.png" || pluginReadmePattern.MatchString(name) {
				continue
			}
			return fmt.Errorf("unexpected non-directory %s under plugins/%s/%s/", name, kindDir, key)
		}
		version := versionEntry.Name()
		relativePath := filepath.ToSlash(filepath.Join("plugins", kindDir, key, version, "plugin.js"))
		source, readErr := os.ReadFile(filepath.Join(root, relativePath))
		if readErr != nil {
			return fmt.Errorf("plugin %s/%s: %w", key, version, readErr)
		}
		loaded, compileErr := jsplugin.CompilePlugin(string(source), jsplugin.Options{})
		if compileErr != nil {
			return fmt.Errorf("plugin %s/%s does not compile: %w", key, version, compileErr)
		}
		if loaded.Meta.Key != key {
			return fmt.Errorf("plugin %s/%s declares key %q; the directory and meta must agree", key, version, loaded.Meta.Key)
		}
		if loaded.Meta.Version != version {
			return fmt.Errorf("plugin %s/%s declares version %q; the directory and meta must agree", key, version, loaded.Meta.Version)
		}
		versionDir := filepath.ToSlash(filepath.Join("plugins", kindDir, key, version))
		readme, docErr := loadDocFile(root, versionDir, "README")
		if docErr != nil {
			return docErr
		}
		if readme == nil {
			readme = pluginReadme
		}
		changelog, docErr := loadDocFile(root, versionDir, "CHANGELOG")
		if docErr != nil {
			return docErr
		}
		digest := sha256.Sum256(source)
		entry := byKey[key]
		if entry == nil {
			entry = &indexPlugin{Key: key}
			byKey[key] = entry
		}
		// Display fields describe the latest version. ReadDir yields directories in
		// lexical order, so the first one seen is usually the oldest release.
		if entry.displayVersion == "" || semverLess(entry.displayVersion, version) {
			entry.displayVersion = version
			entry.Name = loaded.Meta.Name
			entry.Icon = loaded.Meta.Icon
			entry.Website = loaded.Meta.Website
			entry.SortPriority = loaded.Meta.SortPriority
			entry.Description = loaded.Meta.Description
			entry.Routes = loaded.Meta.Routes
			entry.Protocols = loaded.Meta.Protocols
			entry.UsageSchema = loaded.Meta.UsageSchema
			entry.UsageExamples = loaded.Meta.UsageExamples
			entry.UsageProfiles = loaded.Meta.UsageProfiles
			entry.ChannelTypes = loaded.Meta.ChannelTypes
			entry.Models = loaded.Meta.Models
			// Task-only hosts require buildSubmitRequest during compilation.
			// Use the host's part metadata when it exposes that interface.
			entry.Task, entry.Provider, entry.Transform = true, false, false
			if parts, ok := any(&loaded.Meta).(interface {
				HasTaskPart() bool
				HasProviderPart() bool
				HasTransformPart() bool
			}); ok {
				entry.Task = parts.HasTaskPart()
				entry.Provider = parts.HasProviderPart()
				entry.Transform = parts.HasTransformPart()
			}
		}
		entry.Versions = append(entry.Versions, indexVersion{
			Version:       version,
			Path:          relativePath,
			Sha256:        fmt.Sprintf("%x", digest),
			MinApiVersion: loaded.Meta.APIVersion,
			Kind:          kind,
			AllowedHosts:  loaded.Meta.AllowedHosts,
			BaseUrl:       loaded.Meta.BaseURL,
			Auth:          loaded.Meta.Auth.Type,
			Readme:        readme,
			Changelog:     changelog,
		})
	}
	if len(byKey[key].Versions) == 0 {
		return fmt.Errorf("plugin %s has no versions", key)
	}
	for _, name := range []string{"icon.svg", "icon.png"} {
		iconPath := filepath.Join(keyDir, name)
		data, readErr := os.ReadFile(iconPath)
		if readErr != nil {
			continue
		}
		mediaType := "image/svg+xml"
		if name == "icon.png" {
			mediaType = "image/png"
		}
		if iconErr := jsplugin.ValidateIconImage(mediaType, data); iconErr != nil {
			return fmt.Errorf("plugin %s %s: %w", key, name, iconErr)
		}
		if len(data) > jsplugin.MaxIconDataURIBytes {
			return fmt.Errorf("plugin %s %s exceeds %d bytes", key, name, jsplugin.MaxIconDataURIBytes)
		}
		digest := sha256.Sum256(data)
		byKey[key].IconFile = &indexIconFile{Path: filepath.ToSlash(filepath.Join("plugins", kindDir, key, name)), Sha256: fmt.Sprintf("%x", digest)}
		break
	}
	return nil
}

// loadDocFile describes <dir>/<base>.md and the translations listed in its
// front matter. It returns nil when the file does not exist. The Markdown body
// is not checked here; the spec's README and CHANGELOG rules apply to it.
func loadDocFile(root, dir, base string) (*indexDocFile, error) {
	relativePath := dir + "/" + base + ".md"
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(relativePath)))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	translations, err := frontMatterTranslations(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", relativePath, err)
	}
	if len(translations) > maxDocTranslations {
		return nil, fmt.Errorf("%s: translations lists %d files; at most %d are allowed", relativePath, len(translations), maxDocTranslations)
	}
	doc := &indexDocFile{Path: relativePath, Sha256: fmt.Sprintf("%x", sha256.Sum256(data))}
	tags := make([]string, 0, len(translations))
	for tag := range translations {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	for _, tag := range tags {
		if !localeTagPattern.MatchString(tag) || canonicalLocaleTag(tag) != tag {
			return nil, fmt.Errorf("%s: translation tag %q is not a canonical BCP 47 tag", relativePath, tag)
		}
		if tag == "en" {
			return nil, fmt.Errorf("%s: translation tag \"en\" is not allowed; %s.md is the English file", relativePath, base)
		}
		expected := base + "." + tag + ".md"
		if translations[tag] != expected {
			return nil, fmt.Errorf("%s: translation %q must be the sibling file %s, not %q", relativePath, tag, expected, translations[tag])
		}
		translationPath := dir + "/" + expected
		translated, readErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(translationPath)))
		if errors.Is(readErr, fs.ErrNotExist) {
			return nil, fmt.Errorf("%s: listed translation %s does not exist", relativePath, translationPath)
		}
		if readErr != nil {
			return nil, readErr
		}
		if doc.Translations == nil {
			doc.Translations = make(map[string]indexDocFileEntry, len(tags))
		}
		doc.Translations[tag] = indexDocFileEntry{Path: translationPath, Sha256: fmt.Sprintf("%x", sha256.Sum256(translated))}
	}
	return doc, nil
}

// frontMatterTranslations reads the translations map from optional YAML front
// matter: a first line of exactly "---" up to the next line of exactly "---".
func frontMatterTranslations(data []byte) (map[string]string, error) {
	lines := strings.Split(string(data), "\n")
	if strings.TrimSuffix(lines[0], "\r") != "---" {
		return nil, nil
	}
	for index := 1; index < len(lines); index++ {
		if strings.TrimSuffix(lines[index], "\r") != "---" {
			continue
		}
		var matter struct {
			Translations map[string]string `yaml:"translations"`
		}
		if err := yaml.Unmarshal([]byte(strings.Join(lines[1:index], "\n")), &matter); err != nil {
			return nil, fmt.Errorf("invalid front matter: %w", err)
		}
		return matter.Translations, nil
	}
	return nil, errors.New("front matter has no closing --- line")
}

// canonicalLocaleTag applies the BCP 47 casing plugin metadata uses: language
// lowercase, 2-letter region uppercase, 4-letter script title case.
func canonicalLocaleTag(tag string) string {
	parts := strings.Split(tag, "-")
	parts[0] = strings.ToLower(parts[0])
	for index := 1; index < len(parts); index++ {
		switch len(parts[index]) {
		case 2:
			parts[index] = strings.ToUpper(parts[index])
		case 4:
			lowered := strings.ToLower(parts[index])
			parts[index] = strings.ToUpper(lowered[:1]) + lowered[1:]
		default:
			parts[index] = strings.ToLower(parts[index])
		}
	}
	return strings.Join(parts, "-")
}

// semverLess orders release versions numerically; prerelease/build suffixes
// (already constrained by the manifest version pattern) compare as strings,
// which is sufficient for choosing "latest" in a curated repository.
func semverLess(left, right string) bool {
	leftCore, leftRest, _ := strings.Cut(left, "-")
	rightCore, rightRest, _ := strings.Cut(right, "-")
	leftParts := strings.Split(leftCore, ".")
	rightParts := strings.Split(rightCore, ".")
	for index := 0; index < 3 && index < len(leftParts) && index < len(rightParts); index++ {
		leftNumber, _ := strconv.Atoi(leftParts[index])
		rightNumber, _ := strconv.Atoi(rightParts[index])
		if leftNumber != rightNumber {
			return leftNumber < rightNumber
		}
	}
	return leftRest < rightRest
}
