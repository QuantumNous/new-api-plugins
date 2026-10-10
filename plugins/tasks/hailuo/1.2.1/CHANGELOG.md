---
changelogVersion: 1
plugin: "hailuo"
version: "1.2.1"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.1]

### Fixed

- Finished videos can be downloaded again on every model except `MiniMax-H3`. The plugin now asks MiniMax for the file's download link and downloads from it; before, it used an address MiniMax does not serve, so the download failed.

### Migration

- **Installation:** The fix needs a new-api version that lets plugins read upstream data (`utils.fetch`). On older versions the download still fails.
