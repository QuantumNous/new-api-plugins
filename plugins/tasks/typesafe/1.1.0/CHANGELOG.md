---
changelogVersion: 1
plugin: "typesafe"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Fixed

- Requests now reach TypeSafe with `state`, `questions`, and every nested object in the order you sent them. Earlier versions sorted object keys alphabetically, which changed what Jev read and could make scores differ from calling TypeSafe directly.

### Migration

- Upgrade New API to a version later than v1.0.0-rc.41 before installing this version. Earlier gateways refuse to load it.
- If you call TypeSafe through another New API gateway, upgrade that gateway and its TypeSafe plugin as well, or the order is still lost there.
