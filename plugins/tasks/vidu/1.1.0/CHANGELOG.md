---
changelogVersion: 1
plugin: "vidu"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Added

- Use another New API server as the upstream for Vidu video generation. Enable this version on both servers and bind the plugin to a New API channel that holds the upstream address and API key.
- Send requests in Vidu's own format to `/vidu/ent/v2/text2video`, `/vidu/ent/v2/img2video`, `/vidu/ent/v2/start-end2video` or `/vidu/ent/v2/reference2video`, and read the result at `/vidu/ent/v2/tasks/{task_id}/creations`.

### Fixed

- Vidu now generates the model, length and resolution the request is billed for. Before, a `model`, `duration` or `resolution` inside `metadata` replaced them after billing, so the generated video could be a different model, or longer or sharper, than the one charged.

### Migration

- **Installation:** Install 1.1.0 on new-api v1.0.0-rc.39 or later; older versions reject the plugin.
- **Requests:** Set the video length with the top-level `seconds` or `duration` field. A `duration` or `model` inside `metadata` is now ignored, and a `resolution` inside `metadata` applies only when the request sets no `size` or `resolution`.
