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
      "task": true,
      "provider": false,
      "transform": false,
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
          "auth": "api_key",
          "readme": {
            "path": "plugins/tasks/doubao/README.md",
            "sha256": "<hex sha256 of the README bytes>",
            "translations": {
              "zh-CN": {"path": "plugins/tasks/doubao/README.zh-CN.md", "sha256": "<hex sha256>"}
            }
          },
          "changelog": {
            "path": "plugins/tasks/doubao/1.2.0/CHANGELOG.md",
            "sha256": "<hex sha256 of the CHANGELOG bytes>",
            "translations": {
              "zh-CN": {"path": "plugins/tasks/doubao/1.2.0/CHANGELOG.zh-CN.md", "sha256": "<hex sha256>"}
            }
          }
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
- `task`, `provider` and `transform` say which parts the latest version has, detected the way the gateway detects them from the compiled manifest and exports: `task` when the plugin exports `buildSubmitRequest`, `provider` when `meta.relay` is not empty, and `transform` (request processing) when it exports `ingress`, `request`, `outbound`, `egress` or `report`. A plugin may have several parts; marketplace UIs list it under each one, and list an entry that names none as a task plugin. Parts are independent of `kind`: a provider or request processing plugin still publishes its versions as `"task"`.
- `iconFile` points at the plugin's sidecar logo, `plugins/<kind>/<key>/icon.svg` or `icon.png`, stored once per key beside the version directories. The gateway loads it from the index origin at install time and stores it apart from the plugin source; the source itself never carries image data. `meta.icon` stays a LobeHub name or `text` fallback.
- `allowedHosts`, `baseUrl` and `auth` are display hints for the install confirmation page, derived from compiled meta like every other index field. `baseUrl` is the default upstream address a gateway copies onto a Task Plugin channel when the administrator leaves Base URL empty; showing it before install lets the administrator see where the channel key will be sent.
- `readme` and `changelog` point at the version's usage guide and release notes. Each has a `path`, resolved like the version `path`, the file's `sha256`, and an optional `translations` map from a language tag to the translated file's `path` and `sha256`. A version without the file has no field. Other marketplaces may add `sourceUrl`, a page where readers can see the file. This repository leaves it out. Gateways that find no `changelog` look for `CHANGELOG.md` beside `plugin.js`.
- `routes`, `protocols`, `usageSchema`, `usageExamples`, and `usageProfiles` carry the latest version's compiled display metadata. Usage fields retain localized `description`, `unitLabel`, and `enumLabels`; profiles retain their model lists, complete schemas, and examples. Consumers should use these fields when present and may fall back to source parsing for older indexes. Empty or null compiled usage collections mean no entries. This preserves metadata defined through JavaScript constants or expressions without requiring consumers to execute plugin source.

## Trust model

The index is a **derived cache, never a trust anchor**:

- Display fields (`name`, `channelTypes`, `models`) are generated from the compiled plugin meta by `tools/pluginindex`; hand-edited indexes will fail the official repository's CI.
- Installation fetches the plugin source, verifies `sha256`, and then submits it through the standard upload pipeline — compile, `ValidateV1Meta`, pre-flight conflict check. Whatever the index claimed, the compiled meta decides.
- The index carries the three things a single plugin file cannot: the catalog (raw hosting has no directory listing), the cross-version table, and the integrity hash (a hash cannot live inside the file it hashes).

## Repository layout convention

```
plugins/<kind>/<key>/<version>/plugin.js
plugins/<kind>/<key>/<version>/CHANGELOG.md        # release notes for this version
plugins/<kind>/<key>/<version>/CHANGELOG.zh-CN.md  # optional translation
plugins/<kind>/<key>/README.md                     # usage guide for every version
plugins/<kind>/<key>/README.zh-CN.md               # optional translation
plugins/<kind>/<key>/icon.svg                      # optional sidecar logo (or icon.png)
index.json
```

Directory names are organization; the compiled meta is the declaration. `tools/pluginindex` enforces that directory `<key>`/`<version>` agree with `meta.key`/`meta.version`, and that the kind directory matches the plugin kind (task plugins never declare `meta.kind`; the default is exactly the `tasks/` directory). Published version directories are immutable — ship changes as a new version.

## Usage guide and release notes

The plugin API spec defines both file formats:

- [Usage guide (README)](https://github.com/QuantumNous/new-api/blob/main/docs/plugin-api/v1.md#usage-guide-readme)
- [Changelog](https://github.com/QuantumNous/new-api/blob/main/docs/plugin-api/v1.md#changelog)

This repository adds three rules:

- Each plugin has one usage guide, `plugins/<kind>/<key>/README.md`, beside the version directories. `tools/pluginindex` attaches it to every version. A `README.md` inside a version directory replaces it for that version.
- Each new version has its own `CHANGELOG.md` beside `plugin.js`. Older versions published without one stay as they are.
- Published version directories are immutable. Update the plugin-level README in place, and ship release notes only with a new version.

`tools/pluginindex` fails when a translation listed in the front matter is missing or is not named `README.<tag>.md` or `CHANGELOG.<tag>.md` with a canonical tag other than `en`. It does not check the Markdown body.

## Tooling

From this repository, with a new-api checkout as a sibling directory (`../new-api`, required by the module replace directive):

```bash
cd tools/pluginindex
go run . generate ../..   # (re)build index.json
go run . check ../..      # CI gate: recompile all, diff index
```
