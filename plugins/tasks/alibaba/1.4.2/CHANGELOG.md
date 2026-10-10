---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.2"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.4.2]

### Fixed

- `wan3.0-video` and `wan3.0-video-prime` now accept smart duration (`-1`) in `metadata.parameters.duration` on `/v1/video/generations`. Before, the gateway rejected the request. As on the other routes, 30 seconds are held when the task starts and the final charge follows the durations Alibaba Cloud reports.

### Migration

- **Installation:** Install 1.4.2 on a new-api version that supports model-chosen video duration (`duration-auto@1`); older versions reject the plugin.
