---
changelogVersion: 1
plugin: "auto-reasoning"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Set the reasoning effort of each request from its size, with rules on the number of messages and tools, the length of the text and whether it has images.
- Give requests that no rule matches a default effort, and keep every effort between a lowest and a highest level.
- Keep the effort a client already set, or replace it with the plugin's effort.
- By default, requests with tools or with at least 4,000 characters of text get high effort.
