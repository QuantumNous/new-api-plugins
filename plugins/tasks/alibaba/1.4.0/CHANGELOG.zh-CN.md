---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.0"
locale: "zh-CN"
---
# Changelog

## [1.4.0]

### Added

- 新增千问图像系列，支持文生图和图片编辑，涵盖 Qwen-Image、Plus、Max、2.0/Pro、3.0/Pro 和 Edit 系列。可通过原生图片接口和 OpenAI Responses 调用。
- 新增 Z-Image-Turbo 文生图，可选择开启提示词改写。
- 支持通过 OpenAI 兼容接口 `/v1/images/generations` 和 `/v1/images/edits` 生成、编辑图片。默认生成一张，等待图片生成完成后返回结果。
- 编辑图片时可直接上传参考图，也可传入图片链接或 Base64 数据。OpenAI 图片接口可返回链接或 Base64 结果，单张上传图片最大为 10 MiB。
- 新增 `wan2.5-i2i-preview` 图片编辑，支持一至三张参考图，可通过原生图片接口、OpenAI Responses 和 OpenAI 图片接口调用。
- 新增 `wanx2.1-imageedit` 单图编辑，支持蒙版编辑、风格化、扩图、超分辨率、上色、去水印等十种操作。通过 OpenAI 编辑接口调用时，默认按文字指令编辑，传入蒙版后默认改为蒙版编辑。
- 图片生成可选择直接改写或由智能体辅助改写提示词，也可开启模型思考，具体取决于模型支持的选项。
- 支持将另一台 New API 作为上游，调用图片和视频生成，也支持同步返回包含文字和图片的完整结果。

### Changed

- Qwen-Image-3.0 支持为 1K、2K 输出图片分别定价，并单独计算输入图片费用。
- Z-Image 支持按是否开启提示词改写设置不同价格。使用按次计费时，开启改写会使图片费用翻倍。
- `wan2.6-i2v-flash` 支持为有声和无声视频分别定价，默认生成音频。
- 同步原生图片接口的结果现在只在创建图片时返回，之后无法再次查询获取。

### Migration

- 为要启用的图片模型配置价格。Qwen-Image-3.0 的计费表达式用 `u("image_count")` 计算输出张数，用 `u("output_image_type")` 区分 1K、2K 单价，用 `u("input_image_count")` 计算输入张数。
- Z-Image 的计费表达式用 `u("prompt_extend")` 区分是否开启提示词改写。使用按次计费时，请在开启改写前核对基础单价，确认翻倍后的收费符合预期。
- 为 `wan2.6-i2v-flash` 设置有声和无声两种价格时，在计费表达式中使用 `u("audio")` 区分。
- 本版依赖[更新后的插件 API](../../../../docs/plugin-api/v1.md)，请先将 new-api 升级到支持该 API 的版本，再安装插件。
- 连接上游 New API 时，两台服务器均需启用本插件，并在类型 60 渠道中填写上游地址和 API 密钥。
- 请在创建图片时保存同步原生接口返回的结果。切换到 OpenAI 图片接口后，需要多张图片时请设置 `n`。
