---
changelogVersion: 1
plugin: "doubao"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 新增 Seedream 5.0 Pro、5.0 Lite、5.0、4.5 和 4.0，支持文生图和参考图编辑，可通过 `/doubao/api/v3/images/generations` 或 OpenAI Responses 调用。
- 可按模型支持的选项设置图片尺寸、提示词优化、水印、随机种子和输出格式。
- 支持在一次请求中生成一组图片，适用于支持连续组图的模型。参考图与生成图片合计最多 15 张；开启组图后，默认最多生成 15 张。
- Seedream 5.0 Pro 可将一张图片拆成一张底图和最多 16 个图层，支持 1K、1.5K、2K 或自动尺寸。通过原生接口调用时，符合输入和输出格式要求的请求还可生成透明背景。
- OpenAI Responses 支持同步、流式和后台生成图片，生成后可再次获取结果。原生图片接口在生成完成后返回图片链接，不支持流式或 Base64 输出。
- 支持将另一台 New API 作为上游，调用 Seedream 图片生成和 Seedance 视频生成。

### Changed

- Seedream 支持为 1.5K 及以下（不超过 261 万像素）和更大尺寸的图片分别定价，并单独计算参考图费用。图层拆分可设置独立价格，每个生成的图层按自身尺寸计费。
- Seedance 现在会拒绝所选模型不支持的分辨率。
- Seedance 1.5 Pro 支持为有声和无声视频分别定价，默认生成音频。
- 参考视频计费现在仅适用于 Seedance 2.x 模型。

### Migration

- 为要启用的 Seedream 模型配置价格。按图片尺寸或参考图收费时，请使用计费表达式：`u("images_up_to_1_5k")` 和 `u("images_above_1_5k")` 分别计算两档尺寸的输出张数，`u("input_images")` 计算参考图张数。按次计费只计算输出张数。
- Seedream 5.0 Pro 从第二张参考图开始收费时，使用 `max(u("input_images") - 1, 0)`；需要为图层拆分单独定价时，使用 `u("layer_decomposition")` 区分。
- 已有 Seedance 1.0、1.5 计费表达式中使用了 `video_input` 的，请移除该条件。为 1.5 Pro 设置有声和无声两种价格时，使用 `u("generate_audio")` 区分。
- 本版依赖[更新后的插件 API](../../../../docs/plugin-api/v1.md)，请先将 new-api 升级到支持该 API 的版本，再安装插件。
- 连接上游 New API 时，两台服务器均需启用本插件，并在类型 60 渠道中填写上游地址和 API 密钥。
- 请在创建图片时保存原生接口返回的结果；需要之后再次获取图片时，改用 OpenAI Responses。
