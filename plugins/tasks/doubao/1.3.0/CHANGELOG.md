---
changelogVersion: 1
plugin: "doubao"
version: "1.3.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.3.0]

### Added

- Let Seedance 1.5 pro, the Seedance 2.0 models, and Seedance 2.5 choose the video length by sending a duration of `-1`. On `/v1/videos`, send it as `seconds` or `duration`; on `/v1/video/generations`, send it as `metadata.duration`.

### Changed

- When a request sends `-1` or no duration, the amount held at the start now covers the model's longest video: 12 seconds for Seedance 1.0 and 1.5 pro, 15 seconds for the Seedance 2.0 models, and 30 seconds for Seedance 2.5 and for endpoint IDs the plugin does not recognize. Before, 15 seconds were held for every model. The final charge still follows the tokens Ark reports.

### Fixed

- On `/v1/videos`, a JSON request that sets `duration` instead of `seconds` now gets the requested length. Before, the value never reached Ark, which used its default length.

### Migration

- **Installation:** Install 1.3.0 on a new-api version that supports model-chosen video duration (`duration-auto@1`); older versions reject the plugin.
