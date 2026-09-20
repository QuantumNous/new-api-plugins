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

- Support Suno song and lyrics generation through an upstream New API gateway on a type-60 channel, with the same plugin enabled on both gateways.

### Migration

- **Pricing:** Usage quantities and billing conditions are unchanged; no price or billing-expression reconfiguration is required for this release.
- **Gateway connection:** Upgrade to a new-api build supporting the [current plugin contract](../../../../docs/plugin-api/v1.md) before installing 1.1.0. For an upstream New API connection, enable this plugin on both gateways, select a type-60 channel, and configure the upstream gateway URL and token.
