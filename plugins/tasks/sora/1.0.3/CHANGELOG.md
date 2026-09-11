---
changelogVersion: 1
plugin: "sora"
version: "1.0.3"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.3]

### Fixed

- Preserve fields from the latest persisted upstream response, including output URLs and provider metadata, when rendering OpenAI Video retrieval from an object-valued `task.data`. The host still replaces canonical task fields; tasks without an object snapshot retain the legacy renderer.
