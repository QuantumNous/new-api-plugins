---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.2"
locale: "zh-CN"
---
# Changelog

## [1.4.2]

### Fixed

- `wan3.0-video` 和 `wan3.0-video-prime` 在 `/v1/video/generations` 上现在可以通过 `metadata.parameters.duration` 传 `-1` 使用智能时长，此前网关会拒绝这个请求。和其他接口一样，任务开始时预扣 30 秒，最终按阿里云返回的时长扣费。

### Migration

- **安装要求：**1.4.2 需要安装在支持模型自选视频时长（`duration-auto@1`）的 new-api 版本上，旧版本会拒绝该插件。
