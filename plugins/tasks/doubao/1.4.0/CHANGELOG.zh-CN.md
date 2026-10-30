---
changelogVersion: 1
plugin: "doubao"
version: "1.4.0"
locale: "zh-CN"
---
# Changelog

## [1.4.0]

### Added

- 可以在网关上直接调用火山方舟官方路径：视频任务用 `POST /api/v3/contents/generations/tasks` 和 `GET /api/v3/contents/generations/tasks/{task_id}`，Seedream 图片用 `POST /api/v3/images/generations`。网关同时在 `/doubao` 下提供每个路径，原来的 `/doubao/api/v3/...` 地址继续可用。
- 原生 Seedream 请求设置 `stream: true` 时，图片会随方舟的返回实时流式发给客户端。按实际返回的图片收费，生成任何图片之前就失败的流不收费。
- 原生 Seedream 请求可以使用 `response_format: "b64_json"`。OpenAI Responses 请求仍返回图片链接。

### Changed

- 原生请求按原样发给方舟：`content` 数组原样发送，不再合并进提示词；`sequential_image_generation` 和 `stream` 交给方舟检查。响应保留方舟的格式、状态码和错误码，例如 `InvalidParameter`、`AuthenticationError`。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 查询视频任务时返回方舟最后一次的响应。网关判定任务失败后（例如超时），查询会显示任务失败，即使方舟最后一次返回的状态仍是运行中。
- 带 `callback_url` 的请求会被拒绝，因为方舟回调时用的任务 ID 不是网关发出的 ID。请改为查询任务。

### Migration

- **安装要求：**1.4.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的方舟路径。
