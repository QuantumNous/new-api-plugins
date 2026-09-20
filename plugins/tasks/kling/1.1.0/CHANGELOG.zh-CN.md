---
changelogVersion: 1
plugin: "kling"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 可灵视频生成支持通过类型 60 渠道连接上游 New API 网关，两端均需启用同一插件。 渠道明确指定 New API 上游时，使用网关凭据和带插件前缀的提交及查询路径，也支持不带 sk- 前缀的令牌。

### Migration

- **价格：** 用量及计费条件保持不变，本次升级无需重新配置价格或计费表达式。
- **网关连接：** 安装 1.1.0 前，升级至支持[当前插件契约](../../../../docs/plugin-api/v1.md)的 new-api 构建。连接上游 New API 时，两端均需启用本插件，选择类型 60 渠道，并配置上游网关地址和令牌。
