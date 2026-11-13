---
changelogVersion: 1
plugin: "doubao"
version: "1.5.0"
locale: "zh-CN"
---
# Changelog

## [1.5.0]

### Added

- 新增 Seedream 5.5（`doubao-seedream-5-5-261028`）的图片生成和编辑，支持 1K、1.5K、2K 尺寸，最多 10 张参考图。
- Seedream 5.5 可以从文本或多张参考图生成透明背景图片。传 `background: "auto"` 由模型决定，传 `background: "transparent"` 时也不受 Seedream 5.0 pro 和 flash 只能有一张参考图的限制。
- Seedream 5.5 的图层拆分支持 `layer_decomposition_options.max_images`，可设 2 到 100 张（含底图）。不设置时方舟最多返回 30 张。

### Changed

- Seedream 5.5 图层拆分在图片返回前，按请求的 `max_images` 预扣，未设置时按 30 张预扣。Seedream 5.0 pro 和 flash 仍按 17 张预扣。
- 插件不认识的方舟接入点 ID 现在也接受 `background: "auto"` 和 `layer_decomposition_options`，图层拆分默认按 30 张预扣，原来是 17 张。

### Migration

- 给用户启用 `doubao-seedream-5-5-261028` 前，先为它设置价格。它的价格表字段和 Seedream 5.0 pro 相同：两档图片生成单价、输入图片单价和图层拆分价格。
