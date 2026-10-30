---
changelogVersion: 1
plugin: "xai"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 可以在网关上直接调用 xAI 官方路径：`POST /v1/videos/generations`、`/v1/videos/edits` 和 `/v1/videos/extensions`，用 `GET /v1/videos/{request_id}` 查询。网关同时在 `/xai` 下提供每个路径，原来的 `/xai/v1/videos/...` 地址继续可用。
- 新增 `21:9` 和 `5:2` 两种宽高比。

### Changed

- 原生请求按原样发给 xAI，响应保留 xAI 的格式、状态码和错误码。插件不再限制 `reference_images`、`keyframes` 和 `reference_audios` 的数量，由 xAI 检查。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 查询仍在生成中的视频时返回 202，与 xAI 一致。
- xAI 仍然收费的失败视频，查询结果会保留 xAI 的 `usage.cost_in_usd_ticks`。
- 带 `storage_options` 的请求会被拒绝。以前这个字段会被直接丢弃，不作提示。
- 网关判定任务失败后（例如超时），原生查询会显示任务失败，即使 xAI 最后一次返回的状态仍是等待中。

### Migration

- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 xAI 路径。
