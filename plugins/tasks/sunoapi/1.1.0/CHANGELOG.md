---
changelogVersion: 1
plugin: "sunoapi"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Added

- Use another New API server as the upstream for Suno song and lyrics generation.

### Migration

- To use this connection method, upgrade new-api to a build supporting the [plugin upstream API](../../../../docs/plugin-api/v1.md).
- To connect to an upstream New API server, enable the SunoAPI plugin on both servers and configure a type-60 channel with the upstream address and API key.
