---
changelogVersion: 1
plugin: "hailuo"
version: "1.4.0"
locale: "zh-CN"
---
# Changelog

## [1.4.0]

### Added

- 可以在网关上直接调用 MiniMax 官方路径：海螺 02、2.3 和 01 系列用 `POST /v1/video_generation` 和 `GET /v1/query/video_generation?task_id=`，`MiniMax-H3` 用 `POST /v2/video_generation` 和 `GET /v2/query/video_generation/{task_id}`。1.3.0 的 `/hailuo/...` 地址继续可用。
- `GET /v1/files/retrieve?file_id=` 会返回生成好的视频的下载链接，与 MiniMax 一致。文件 ID 就是任务 ID。

### Changed

- 原生请求按原样发给 MiniMax：`MiniMax-H3` 请求不再补默认值，`aigc_watermark` 也会保留。错误按 MiniMax 的 `base_resp` 格式返回。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 未指定时长的 `MiniMax-H3` 请求先按 H3 最长的 15 秒预扣，再按 MiniMax 报告的时长结算。可用时长为 4 到 15 秒。
- 带 `callback_url` 的请求会被拒绝，因为 MiniMax 回调时用的任务 ID 不是网关发出的 ID。请改为查询任务。
- 网关判定任务失败后（例如超时），原生查询会显示任务失败，即使 MiniMax 最后一次返回的状态仍是进行中。

### Migration

- **安装要求：**1.4.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 MiniMax 路径。
