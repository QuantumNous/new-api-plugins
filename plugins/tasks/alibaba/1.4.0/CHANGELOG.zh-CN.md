---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.0"
locale: "zh-CN"
---
# Changelog

## [1.4.0]

### Added

- 原生图片生成和 OpenAI Responses 新增 Qwen-Image、Plus、Max、2.0/2.0-Pro、3.0/3.0-Pro、Edit、Edit-Plus、Edit-Max、已声明的日期快照，以及 Z-Image-Turbo。这些模型使用同步上游调用，输出数量和参考图数量遵循所选模型的限制。
- 为支持的图片模型新增 OpenAI 兼容接口 `/v1/images/generations` 和 `/v1/images/edits`。生成接收 JSON，编辑接收 JSON 或 multipart 图片上传。请求默认生成一张图片并返回完整图片列表，上游为异步任务时也会等待结果。
- OpenAI 图片接口支持图片 URL、Base64 数据 URL、multipart 参考图上传，以及 `url` 或 `b64_json` 输出。每张上传的输入图片限制为 10 MiB，不支持的流式请求会被拒绝。
- 在原生 `/ali/api/v1/services/aigc/image2image/image-synthesis` 路由、Responses 和 OpenAI 图片接口新增 `wan2.5-i2i-preview` 与 `wanx2.1-imageedit`，前者接收一至三张参考图，后者接收一张底图。
- 支持 `wanx2.1-imageedit` 的十种编辑操作，包括指令编辑、蒙版编辑、风格化、扩图、超分辨率、上色和去水印。OpenAI 编辑默认使用指令编辑，提供蒙版时默认使用蒙版编辑，并转发操作专属参数。
- 图片请求支持 `prompt_extend_mode`（`direct` 或 `agent`）和 `enable_thinking`，并支持通过 OpenAI 图片接口传入厂商的 `input` 与 `parameters` 对象。
- 支持通过类型 60 渠道连接上游 New API 网关，两端均需启用本插件。图片和视频请求使用插件原生路由，同步图文混排返回完整 JSON。

### Changed

- Qwen-Image-3.0 报告输出图片数量、1K/2K 输出档位及输入图片数量，并采用有效的上游用量。Z-Image 将提示词改写作为计费条件；按次计费时，开启改写会应用两倍系数。
- `wan2.6-i2v-flash` 报告是否生成音频，省略参数时默认为开启，并在上游提供完成音频标记时采用该值。
- 同步原生图片结果仅在创建响应中返回，之后无法通过结果查询或产物接口获取。

### Migration

- **价格配置：** 为新启用的图片模型配置价格。Qwen-Image-3.0 表达式应按 `u("output_image_type")` 为 `u("image_count")` 定价，并计入 `u("input_image_count")`；Z-Image 表达式应处理 `u("prompt_extend")`。已有表达式不会自动改写。Z-Image 按次计费在开启提示词改写时将图片费用乘以二，启用前请核对基础价格。
- **视频价格：** 如需为 `wan2.6-i2v-flash` 的有声和无声输出设置不同价格，请在表达式中按 `u("audio")` 分支；不使用该字段则继续采用同一价格。请核对已保存的分辨率条件与各启用模型支持的档位，已有价格不会自动修改。
- **安装与网关互联：** 安装 1.4.0 前，升级至支持[当前插件契约](../../../../docs/plugin-api/v1.md)的 new-api 构建。使用类型 60 连接时，两端均需启用本插件，并配置上游网关地址和令牌。
- **图片调用：** 请保存同步原生接口创建响应中的结果。依赖特定图片数量时显式指定 `n`：OpenAI 图片请求默认一张，现有原生异步模型仍采用各自默认值。仅在兼容的编辑操作中使用 `function` 和蒙版。
