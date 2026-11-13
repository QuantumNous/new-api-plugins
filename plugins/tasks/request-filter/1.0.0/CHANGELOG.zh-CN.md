---
changelogVersion: 1
plugin: "request-filter"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 拒绝文本含有关键词或命中正则表达式的请求，最多 500 条规则。
- 很长的词表可以放进关键词列表，每行一个关键词。
- 检查对话、Responses、Claude 和 Gemini 请求中的全部文本，任务提交和 Midjourney 提交的提示词，以及随请求发送的文本、PDF 和 DOCX 文件。
- 可以设置请求被拒绝时客户端看到的提示。
- 可选中止命中规则或关键词列表的模型回复。
