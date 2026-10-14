---
changelogVersion: 1
plugin: "google"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Added

- Use another New API server as the upstream for Google Veo video generation on the Gemini API. Enable this version on both servers and bind the plugin to a New API channel that holds the upstream address and API key.
- Send requests in the Gemini API format to `/google/v1beta/models/{model}:predictLongRunning` with a prompt and an optional first-frame `image`. The returned operation is named `models/{model}/operations/{task_id}`; poll it at `/google/v1beta/models/{model}/operations/{task_id}`.

### Migration

- **Installation:** Install 1.1.0 on new-api v1.0.0-rc.39 or later; older versions reject the plugin.
