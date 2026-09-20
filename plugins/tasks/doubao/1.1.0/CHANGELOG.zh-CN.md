---
changelogVersion: 1
plugin: "doubao"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 通过 `/doubao/api/v3/images/generations` 和 OpenAI Responses 新增 Seedream 5.0 Pro、5.0 Lite、5.0、4.5、4.0 图片生成，支持文本提示和参考图编辑。
- 支持各模型的分辨率预设与像素尺寸、参考图数量限制、提示词优化、水印、种子及输出格式。保留显式零值和 false，并在提交前拒绝不支持的参数组合。
- 兼容模型支持连续组图，`max_images` 范围为 1 至 15，启用组图时默认 15。参考图与生成图片共享 15 张上限。
- Seedream 5.0 Pro 支持从一张输入图片拆分图层，输出一张底图和最多 16 层。图层尺寸支持 `1K`、`1.5K`、`2K` 或 `auto`；原生路由在输入和格式组合符合要求时支持透明背景。
- 原生调用返回完整的图片 URL 结果，Responses 在 stream、sync 和 background 模式下返回图片链接。Responses 保留图片产物供后续获取；原生图片创建只返回一次结果，不支持转发上游事件流或 Base64 图片输出。
- 图片和视频请求支持通过类型 60 渠道连接上游 New API 网关，两端均需启用本插件。

### Changed

- 按输出像素不超过 261 万及超过该边界分别报告图片数量，同时报告参考图数量和图层拆分条件。成功生成的各图层按自身尺寸档位分别计量。
- Seedance 请求会拒绝所选模型不支持的分辨率档位。Seedance 1.5 Pro 报告音频生成条件，默认开启；Seedance 2.x 模型报告参考视频输入条件。

### Migration

- **图片价格：** 为新启用的 Seedream 模型配置价格。表达式应为 `u("images_up_to_1_5k")`、`u("images_above_1_5k")` 和 `u("input_images")` 定价，必要时在 `u("layer_decomposition")` 为 true 时采用独立价格。Seedream 5.0 Pro 从第二张参考图开始计费时，可使用 `max(u("input_images") - 1, 0)`。按次计费仅按输出图片数量收费，不区分这些条件。已有价格不会自动改写。
- **视频价格：** 升级前检查已有 Seedance 表达式。仅 Seedance 2.x 报告 `video_input`，请从 1.0 和 1.5 模型表达式中移除该字段。1.5 Pro 可使用 `u("generate_audio")` 区分有声与无声价格，并将分辨率条件限定为各模型支持的档位。已有表达式不会自动迁移。
- **安装与网关互联：** 安装 1.1.0 前，升级至支持[当前插件契约](../../../../docs/plugin-api/v1.md)的 new-api 构建。使用类型 60 连接时，两端均需启用本插件，并配置上游网关地址和令牌。
- **图片调用：** 请保存原生图片创建响应中的结果；需要后续获取产物时使用 Responses。原生路由仅支持 URL 输出和非流式请求，原生透明背景字符串 `background` 与 Responses 后台执行布尔标记含义不同。
