---
changelogVersion: 1
plugin: "google"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 可以在网关上直接调用 Gemini API 官方路径：用 `POST /v1beta/models/{model}:predictLongRunning` 发起 Veo 视频生成，用 `GET /v1beta/models/{model}/operations/{operation}` 查询结果。网关会把这些请求交给插件处理，而不是交给 Gemini 转发，同时也在 `/google` 下提供每个路径。
- 原生请求除了 `Authorization: Bearer`，也可以像 Gemini API 一样用 `x-goog-api-key` 请求头或 `?key=` 查询参数鉴权。

### Changed

- 原生请求按原样发给 Gemini API，响应保留它的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 网关判定任务失败后（例如超时），原生查询会显示该操作失败，即使 Google 最后一次返回的状态仍是运行中。

### Fixed

- 生成好的 Veo 视频可以重新下载。插件改为从 Gemini API 实际列出视频的 `generatedSamples` 读取，不再读 `generatedVideos`。

### Migration

- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 Gemini API 路径。
