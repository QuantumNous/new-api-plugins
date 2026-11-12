---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.3.1"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.3.1]

### Fixed

- Video requests now run in the channel's deployment region instead of always in `global`.
- A deployment region written as JSON works as it does for the channel's other models: each model name maps to its own region, and `default` covers the rest. The lookup uses the model name the client sent, before model mapping.

### Migration

- Before upgrading, check the deployment region of your Vertex AI channels. Video requests now go to that region, so the video models you use must be available there. To keep a model in `global`, map it to `global` in the JSON form.
- **Installation:** The deployment region takes effect on new-api releases after v1.0.0-rc.49. On v1.0.0-rc.49 and earlier, requests still run in `global`.
