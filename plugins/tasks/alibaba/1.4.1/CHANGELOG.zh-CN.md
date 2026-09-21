---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.1"
locale: "zh-CN"
---
# Changelog

## [1.4.1]

### Added

- OpenAI 兼容的图片生成和编辑接口在响应中增加 `metadata`，携带上游结果，调用方可在获取图片的同时读取 `request_id`、`usage` 等厂商字段。
