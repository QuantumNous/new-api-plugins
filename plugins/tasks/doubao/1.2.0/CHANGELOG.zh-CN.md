---
changelogVersion: 1
plugin: "doubao"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 支持通过 OpenAI Images API 的 `/v1/images/generations` 和 `/v1/images/edits` 接口调用 Seedream 图片生成与编辑。
- 编辑图片时可通过 multipart 请求直接上传参考图，支持同时上传多张。每张上传图片上限为 30 MiB，同时受网关下载大小限制；JSON 请求可使用图片 URL 或 Base64 data URL 作为参考图。
- 通过 `response_format: "url"` 或 `"b64_json"` 选择返回图片链接或 Base64 数据。图片生成完成后一次性返回结果，不支持流式输出和遮罩编辑。
- 在支持的模型上，可通过 OpenAI Images API 使用 Seedream 组图生成和图层拆分。`n` 仅接受 `1`；生成多张图片请使用 `sequential_image_generation`。
- OpenAI Images 调用会报告各尺寸档位的生成图片数量、参考图数量及是否开启图层拆分，供按用量定价使用。

### Migration

- 安装前，如果当前 new-api 版本尚不支持 [OpenAI Images 插件协议](../../../../docs/plugin-api/v1.md#host-protocols)，请先升级网关。
