---
changelogVersion: 1
plugin: "hailuo"
version: "1.3.0"
locale: "zh-CN"
---
# Changelog

## [1.3.0]

### Added

- 支持将另一台 New API 作为上游，调用 MiniMax 海螺视频生成，包括 `MiniMax-H3`。两台服务器都需启用这个版本，并把插件绑定到填写了上游地址和 API 密钥的 New API 渠道。
- 可以直接按 MiniMax 官方格式调用 `/hailuo/v1/video_generation`（海螺 02、2.3 和 01 系列）或 `/hailuo/v2/video_generation`（`MiniMax-H3`）。查询任务用 `/hailuo/v1/query/video_generation/{task_id}` 或 `/hailuo/v2/query/video_generation/{task_id}`；v1 查询的任务 ID 放在路径里，不再用 `task_id` 查询参数。

### Migration

- **安装要求：**1.3.0 需要安装在 new-api v1.0.0-rc.43 或更高版本上。旧版本会拒绝该插件，或无法下载生成好的视频。
