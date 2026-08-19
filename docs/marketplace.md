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
- `allowedHosts` and `auth` are display hints for the install confirmation page, derived from compiled meta like every other index field.

## Trust model

The index is a **derived cache, never a trust anchor**:

- Display fields (`name`, `channelTypes`, `models`) are generated from the compiled plugin meta by `tools/pluginindex`; hand-edited indexes will fail the official repository's CI.
- Installation fetches the plugin source, verifies `sha256`, and then submits it through the standard upload pipeline — compile, `ValidateV1Meta`, pre-flight conflict check. Whatever the index claimed, the compiled meta decides.
- The index carries the three things a single plugin file cannot: the catalog (raw hosting has no directory listing), the cross-version table, and the integrity hash (a hash cannot live inside the file it hashes).

## Repository layout convention

```
plugins/<kind>/<key>/<version>/plugin.js
index.json
```

Directory names are organization; the compiled meta is the declaration. `tools/pluginindex` enforces that directory `<key>`/`<version>` agree with `meta.key`/`meta.version`, and that the kind directory matches the plugin kind (task plugins never declare `meta.kind`; the default is exactly the `tasks/` directory). Published version directories are immutable — ship changes as a new version.

## Tooling

From this repository, with a new-api checkout as a sibling directory (`../new-api`, required by the module replace directive):

```bash
cd tools/pluginindex
go run . generate ../..   # (re)build index.json
go run . check ../..      # CI gate: recompile all, diff index
```
