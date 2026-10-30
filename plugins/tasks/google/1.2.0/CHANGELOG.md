---
changelogVersion: 1
plugin: "google"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Call the Gemini API's own paths on the gateway: `POST /v1beta/models/{model}:predictLongRunning` to start a Veo video and `GET /v1beta/models/{model}/operations/{operation}` to query it. The gateway sends these requests to the plugin instead of its Gemini relay, and also serves each path under `/google`.
- Native requests can authenticate with the `x-goog-api-key` header or the `?key=` query parameter, as the Gemini API takes them, as well as `Authorization: Bearer`.

### Changed

- Native requests reach the Gemini API as written, and replies keep its format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- Once the gateway has failed a task, for example after a timeout, a native query shows the operation as failed even if Google's last reply still showed it running.

### Fixed

- Finished Veo videos download again. The plugin now reads them from `generatedSamples`, where the Gemini API lists them, instead of `generatedVideos`.

### Migration

- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Gemini API paths.
