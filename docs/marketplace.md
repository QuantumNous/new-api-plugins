# Plugin Marketplace Index v1

A marketplace source is any HTTP(S)-reachable repository that serves a root `index.json` in this format plus the plugin files it references. The official source is `QuantumNous/new-api-plugins`; any repository following this contract can be added as a third-party source.

## Format

```json
{
  "indexVersion": 1,
  "name": "new-api Official Plugins",
  "plugins": [
    {
      "key": "doubao",
      "name": "doubao-video",
      "iconFile": {"path": "plugins/tasks/doubao/icon.svg", "sha256": "<hex sha256 of the icon bytes>"},
      "channelTypes": [54, 45],
      "models": ["doubao-seedance-1-0-pro-250528"],
      "latest": "1.2.0",
      "versions": [
        {
          "version": "1.2.0",
          "path": "plugins/tasks/doubao/1.2.0/plugin.js",
          "sha256": "<hex sha256 of the plugin.js bytes>",
          "minApiVersion": 1,
          "kind": "task",
          "allowedHosts": ["example.com"],
          "baseUrl": "https://api.example.com",
          "auth": "api_key"
        }
      ]
    }
  ]
}
```

- `path` is resolved relative to the index URL, so the same repository works through any raw prefix (GitHub raw, jsDelivr, mirrors).
- `sha256` is verified at install time before the source enters the upload pipeline.
- `kind` discriminates plugin kinds; `"task"` is the only kind today. Marketplace UIs filter to kinds the gateway supports.
- `minApiVersion` is a UI hint only. Admission is always decided by the gateway's own manifest validation.
- `iconFile` points at the plugin's sidecar logo, `plugins/<kind>/<key>/icon.svg` or `icon.png`, stored once per key beside the version directories. The gateway loads it from the index origin at install time and stores it apart from the plugin source; the source itself never carries image data. `meta.icon` stays a LobeHub name or `text` fallback.
- `allowedHosts`, `baseUrl` and `auth` are display hints for the install confirmation page, derived from compiled meta like every other index field. `baseUrl` is the default upstream address a gateway copies onto a Task Plugin channel when the administrator leaves Base URL empty; showing it before install lets the administrator see where the channel key will be sent.

## Trust model

The index is a **derived cache, never a trust anchor**:

- Display fields (`name`, `channelTypes`, `models`) are generated from the compiled plugin meta by `tools/pluginindex`; hand-edited indexes will fail the official repository's CI.
- Installation fetches the plugin source, verifies `sha256`, and then submits it through the standard upload pipeline — compile, `ValidateV1Meta`, pre-flight conflict check. Whatever the index claimed, the compiled meta decides.
- The index carries the three things a single plugin file cannot: the catalog (raw hosting has no directory listing), the cross-version table, and the integrity hash (a hash cannot live inside the file it hashes).

## Repository layout convention

```
plugins/<kind>/<key>/<version>/plugin.js
plugins/<kind>/<key>/icon.svg        # optional sidecar logo (or icon.png)
index.json
```

Directory names are organization; the compiled meta is the declaration. `tools/pluginindex` enforces that directory `<key>`/`<version>` agree with `meta.key`/`meta.version`, and that the kind directory matches the plugin kind (task plugins never declare `meta.kind`; the default is exactly the `tasks/` directory). Published version directories are immutable — ship changes as a new version.

## Release notes

Every new or unpublished plugin version has an English `CHANGELOG.md` next to
`plugin.js`. The path is derived from the version's `path` in `index.json` by
replacing the final `plugin.js` filename with `CHANGELOG.md`. Historical published
versions may have no changelog; consumers must tolerate a missing file for them.

Use UTF-8 Markdown with YAML front matter:

```markdown
---
changelogVersion: 1
plugin: "example"
version: "1.2.0"
locale: "en"
---
# Changelog

## [1.2.0]

### Fixed

- Preserve an explicitly supplied zero seed in upstream requests.
```

The front matter starts at the first line with `---` and ends at the next
standalone `---`. Parse it as YAML data, without custom tags or executable
constructors. `changelogVersion: 1` identifies this format. Reject unsupported
format versions rather than guessing their structure.

| Field | Type | Requirement |
| --- | --- | --- |
| `changelogVersion` | integer | Required; `1`. |
| `plugin` | string | Required; matches the plugin directory and `meta.key`. |
| `version` | string | Required; matches the version directory and `meta.version`. |
| `locale` | string | Required BCP 47 language tag; `en` for the canonical file. |
| `translations` | object | Optional, canonical English file only; maps language tags to sibling filenames, such as `zh-CN: CHANGELOG.zh-CN.md`. |

The body has exactly one `# Changelog` heading and one release heading:
`## [<version>]`, matching the metadata. Categories use these exact level-three headings, in this order
when present: `Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, `Security`, `Migration`. Omit empty categories. Each category contains an
unordered list of concrete entries. Inline Markdown and wrapped entry text are
allowed; do not use tables, nested lists, or additional headings for entries.
At least one of `Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, or `Security`
must contain a change. Entries describe actual changes and their compatibility,
default-parameter, or billing impacts only. Do not include validation or test
reports, verification limits, generic upgrade instructions, or lists of unrelated unchanged
behavior. Put those details in commit or PR descriptions.

Include `Migration` whenever a change affects pricing, billing calculations, or
required configuration. Start with the pricing impact: state whether prices or
billing calculations change and whether administrators must reconfigure model
prices, resolution tiers, billing mode, or billing expressions. Describe any
automatic migration and give the concrete action for affected models/settings.
If no price reconfiguration is required, say so explicitly only when supported
by the change. This targeted compatibility statement is allowed. Put other
required compatibility steps after pricing guidance. Do not add validation
reports or generic release procedures.

Consumers can read the metadata first, then use a Markdown parser to extract the
release heading and category list items. Structural keys and headings are always
English. Optional translated prose belongs in `CHANGELOG.<locale>.md` in the
same directory and must be discoverable through the canonical file's
`translations` map. Translations use the same plugin/version and
categories, with their own `locale`. Translation filenames must be basenames
matching `CHANGELOG.<locale>.md`, not URLs or paths outside the version directory.
English remains the fallback when a translation is absent or unavailable.

New plugin versions, their changelogs, and the generated index are committed
together. Do not rewrite an immutable published release solely to add or translate
release notes. The current `pluginindex check` validates the plugin catalog;
changelog metadata and structure require a separate validation step.

## Tooling

From this repository, with a new-api checkout as a sibling directory (`../new-api`, required by the module replace directive):

```bash
cd tools/pluginindex
go run . generate ../..   # (re)build index.json
go run . check ../..      # CI gate: recompile all, diff index
```
