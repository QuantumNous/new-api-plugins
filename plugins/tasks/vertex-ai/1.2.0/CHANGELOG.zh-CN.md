---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 可以在网关上直接调用 Vertex AI 官方路径：用 `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning` 发起 Veo 视频生成，在同一路径上用 `:fetchPredictOperation` 查询结果。`/v1beta1` 下的同名路径同样可用；网关也在 `/vertex-ai` 下提供每个路径，原来的 `/vertex-ai/v1/...` 地址继续可用。
- 新增 Gemini Omni Flash 视频生成（`gemini-omni-1.1-flash-preview` 和 `gemini-omni-flash-preview`），可以通过 OpenAI Responses 和 OpenAI 视频接口用文字或首帧图片生成。视频时长 3 到 10 秒，清晰度可选 360p、720p、1080p 或 4k。任务按 Vertex AI 报告的 token 计费，分为输入、文本输出（含思考）和视频输出。New API 上游渠道不支持 Gemini Omni。

### Changed

- 原生请求按原样发给 Vertex AI，响应保留它的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 网关判定任务失败后（例如超时），原生查询会显示该操作失败，即使 Vertex AI 最后一次返回的状态仍是运行中。

### Fixed

- 生成好的 Veo 视频可以重新下载。Vertex AI 直接在响应里返回的视频不再保存在任务中，下载时会重新向 Vertex AI 获取；写入 Cloud Storage 的视频保留 `gs://` 地址。通过另一台 New API 网关调用时，从那台网关的任务下载视频。

### Migration

- **价格配置：**启用 `gemini-omni-1.1-flash-preview` 和 `gemini-omni-flash-preview` 前，请设置输入、文本输出和视频输出的 token 单价。Vertex AI 的标价为每百万 token 1.50、9 和 17.50 美元。提交时按请求的时长和清晰度预扣视频输出费用。
- **渠道：**Gemini Omni 需要服务账号 JSON 密钥，渠道 API 地址填 `https://aiplatform.googleapis.com`。
- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 Vertex AI 路径。
