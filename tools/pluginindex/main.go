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
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/QuantumNous/new-api/pkg/jsplugin"
)

// kindDirs maps a repository layout directory to the plugin kind it declares.
// The directory is organization, the compiled meta is the declaration; task
// plugins never write meta.kind, which is exactly the "tasks" directory.
var kindDirs = map[string]string{"tasks": "task"}

type indexVersion struct {
	Version       string   `json:"version"`
	Path          string   `json:"path"`
	Sha256        string   `json:"sha256"`
	MinApiVersion int      `json:"minApiVersion"`
	Kind          string   `json:"kind"`
	AllowedHosts  []string `json:"allowedHosts,omitempty"`
	Auth          string   `json:"auth,omitempty"`
}

type indexPlugin struct {
	displayVersion string                   `json:"-"`
	Key            string                   `json:"key"`
	Name           string                   `json:"name"`
	Icon           string                   `json:"icon,omitempty"`
	Description    jsplugin.LocalizedText   `json:"description,omitempty"`
	Protocols      []jsplugin.ProtocolClaim `json:"protocols,omitempty"`
	ChannelTypes   []int                    `json:"channelTypes,omitempty"`
	Models         []string                 `json:"models,omitempty"`
	Latest         string                   `json:"latest"`
	Versions       []indexVersion           `json:"versions"`
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
	for _, versionEntry := range versionEntries {
		if !versionEntry.IsDir() {
			return fmt.Errorf("unexpected non-directory %s under plugins/%s/%s/", versionEntry.Name(), kindDir, key)
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
			entry.Description = loaded.Meta.Description
			entry.Protocols = loaded.Meta.Protocols
			entry.ChannelTypes = loaded.Meta.ChannelTypes
			entry.Models = loaded.Meta.Models
		}
		entry.Versions = append(entry.Versions, indexVersion{
			Version:       version,
			Path:          relativePath,
			Sha256:        fmt.Sprintf("%x", digest),
			MinApiVersion: loaded.Meta.APIVersion,
			Kind:          kind,
			AllowedHosts:  loaded.Meta.AllowedHosts,
			Auth:          loaded.Meta.Auth.Type,
		})
	}
	if len(byKey[key].Versions) == 0 {
		return fmt.Errorf("plugin %s has no versions", key)
	}
	return nil
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
