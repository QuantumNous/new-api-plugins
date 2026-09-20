---
changelogVersion: 1
plugin: "sora"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 支持将另一台 New API 作为上游，调用 Sora 视频生成。

### Migration

- 使用这种连接方式前，请将 new-api 升级到支持[插件上游连接](../../../../docs/plugin-api/v1.md)的版本。
- 连接上游 New API 时，两台服务器均需启用 Sora 插件，并在类型 60 渠道中填写上游地址和 API 密钥。
