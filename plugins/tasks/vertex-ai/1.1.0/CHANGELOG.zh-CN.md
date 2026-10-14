---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 支持将另一台 New API 作为上游，调用 Vertex AI 上的 Google Veo 视频生成。两台服务器都需启用这个版本，并把插件绑定到填写了上游地址和 API 密钥的 New API 渠道。上游服务器会在自己的 Google Cloud 项目里执行所有请求。
- 可以直接按 Vertex AI 格式调用 `/vertex-ai/v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning`，传入提示词和可选的首帧 `image`，再在同一路径上用 `:fetchPredictOperation` 查询进度。返回的 operation 名称以任务 ID 结尾，项目和区域沿用请求路径里的值。

### Migration

- **安装要求：**1.1.0 需要安装在 new-api v1.0.0-rc.39 或更高版本上，旧版本会拒绝该插件。
