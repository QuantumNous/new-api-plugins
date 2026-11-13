---
changelogVersion: 1
plugin: "content-moderation"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 用 OpenAI 兼容的审核接口检查用户消息，拒绝被标记的请求。默认使用 OpenAI 的审核接口和 `omni-moderation-latest` 模型。
- 可以只检查最后一条用户消息，也可以检查全部用户消息，连同其中的图片以及文本、PDF 和 DOCX 文件。
- 检查视频、音乐生成等任务提交和 Midjourney 提交的提示词。
- 可以只拒绝因所选类别被标记的请求，也可以拒绝所有被标记的请求。
- 可选审核模型回复，中止被审核接口标记的回复。
