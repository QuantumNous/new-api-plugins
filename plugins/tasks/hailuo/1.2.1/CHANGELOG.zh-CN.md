---
changelogVersion: 1
plugin: "hailuo"
version: "1.2.1"
locale: "zh-CN"
---
# Changelog

## [1.2.1]

### Fixed

- 除 `MiniMax-H3` 外的模型，成片视频又可以下载了。插件现在先向 MiniMax 查询文件的下载链接，再从这个链接下载；此前用的地址 MiniMax 并不提供，所以下载会失败。

### Migration

- **安装要求：**这个修复需要支持插件读取上游数据（`utils.fetch`）的 new-api 版本，旧版本上下载仍会失败。
