---
changelogVersion: 1
plugin: "xai"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Changed

- `grok-imagine-video` 的编辑和延长改按成片的 480p 或 720p 价格收费，与 xAI 的定价一致。编辑或延长后的视频沿用输入视频的分辨率，最高 720p，所以任务开始时先按 720p 预扣，成片不超过 480p 时按 480p 结算。
- 为了得知成片的分辨率，插件会从 `vidgen.x.ai` 读取成片开头的一小段。

### Removed

- 去掉 `grok-imagine-video` 价格表中的“编辑或延长视频”一行。该列恢复为“输出视频分辨率”，分 480p 和 720p 两行。

### Migration

- **价格配置：**`grok-imagine-video` 的编辑和延长不再使用“编辑或延长视频”的价格。如果这个价格和 720p 价格不同，请检查 480p 和 720p 的价格，编辑和延长现在按这两档收费。
- **安装要求：**编辑和延长要按 480p 结算，需要支持插件读取上游数据（`utils.fetch`）的 new-api 版本；旧版本上一律按 720p 收费。
