---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.3.0"
locale: "zh-CN"
---
# Changelog

## [1.3.0]

### Added

- 支持通过 OpenAI Responses 和 OpenAI 视频接口调用 `gemini-omni-1.2-flash`，视频时长可设为 3 到 15 秒的整数。Vertex AI 对 Omni 1.1 和首尾帧插值的时长上限为 10 秒。
- 支持通过 OpenAI Responses 向 Gemini Omni 提交多张图片。两张图片分别用作首帧和尾帧，更多图片用作主体参考。图片按提交顺序传给 Vertex AI。

### Fixed

- Gemini Omni 任务失败时，如果 Vertex AI 的错误列表提供了错误码和说明，失败原因会显示这些内容。
- Responses 请求会在提交前校验每张输入图片，第一张之后的无效图片也会被拒绝。

### Migration

- 启用 `gemini-omni-1.2-flash` 前，请分别配置输入、文本输出（含思考）和视频输出的每百万 token 单价。
