---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.1"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.4.1]

### Added

- Include the upstream result in `metadata` on OpenAI-compatible image generation and editing responses, so callers can read provider fields such as `request_id` and `usage` alongside the images.
