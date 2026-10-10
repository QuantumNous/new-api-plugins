---
changelogVersion: 1
plugin: "doubao"
version: "1.3.0"
locale: "zh-CN"
---
# Changelog

## [1.3.0]

### Added

- Seedance 1.5 pro、Seedance 2.0 系列和 Seedance 2.5 支持由模型自行决定视频时长，时长传 `-1` 即可。`/v1/videos` 可放在 `seconds` 或 `duration` 中；`/v1/video/generations` 请放在 `metadata.duration` 中。

### Changed

- 请求传 `-1` 或不传时长时，任务开始时的预扣改为覆盖该模型最长的视频：Seedance 1.0 和 1.5 pro 为 12 秒，Seedance 2.0 系列为 15 秒，Seedance 2.5 以及插件无法识别的接入点 ID 为 30 秒。此前所有模型都预扣 15 秒。最终扣费仍以方舟返回的 token 用量为准。

### Fixed

- `/v1/videos` 的 JSON 请求用 `duration` 而不是 `seconds` 指定时长时，现在会按请求的时长生成。此前这个值传不到方舟，方舟会按默认时长生成。

### Migration

- **安装要求：**1.3.0 需要安装在支持模型自选视频时长（`duration-auto@1`）的 new-api 版本上，旧版本会拒绝该插件。
