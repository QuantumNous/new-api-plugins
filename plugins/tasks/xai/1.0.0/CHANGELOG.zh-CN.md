---
changelogVersion: 1
plugin: "xai"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 使用 `grok-imagine-video`、`grok-imagine-video-1.5` 或 `grok-imagine-video-1.5-lite` 通过 xAI Grok Imagine 生成视频：支持文生视频、图生视频和参考图生视频，可附加关键帧和参考音频。
- 通过 `POST /xai/v1/videos/edits` 和 `POST /xai/v1/videos/extensions`，用 `grok-imagine-video` 编辑或延长视频。
- 可调用 `/xai/v1/videos` 下的 xAI 原生路由，也可使用 OpenAI 兼容的 `/v1/videos` 和 `/v1/responses` 接口，支持流式和后台响应。
- 每个视频按输出秒数（按模型和分辨率区分）、输入图片数计费；`grok-imagine-video` 的编辑和延长还按输入视频秒数计费。
- xAI 生成后又因内容审核拦截的视频，以及 xAI 仍然收费的失败请求，会照常计费而不退款，因为 xAI 对它们收费。
- 可以用 xAI 渠道及其现有 API Key 提供这些模型，也可以通过已安装此插件的另一台 New API 网关调用。

### Migration

- 为 Grok Imagine 模型启用任务用量表达式计费：为每种分辨率设置每秒输出单价，设置每张输入图片单价，并为 `grok-imagine-video` 设置每秒输入视频单价。xAI 不返回编辑和延长的分辨率，这两类请求按“与输入视频相同”分辨率计费；被拦截的视频可以单独定价。
