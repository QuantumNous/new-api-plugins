---
changelogVersion: 1
plugin: "param-override"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Apply parameter override rules to all requests or to chosen groups, with the same syntax as a channel's parameter override.
- Set, delete or rewrite body fields and request headers, with conditions on the model, group, channel, path or request format.
- Reject requests with a `return_error` rule. The client gets HTTP 400 with the rule's message.
- A channel's own parameter override runs after these rules and wins on the same field or header.
