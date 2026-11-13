---
changelogVersion: 1
plugin: "request-redact"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 在请求发往上游前，替换请求文本中的邮箱地址、电话号码、身份证号、银行卡号和 API 密钥，也可以开启 IPv4 地址。
- 可以添加自己的正则表达式，每条使用自己的替换文本。
- 对对话、Responses、Claude 和 Gemini 请求的文本，任务提交和 Midjourney 提交的提示词，以及随请求发送的文本文件进行脱敏。
- 请求发送前如果仍含此类信息（例如在无法修改的 PDF 或 DOCX 文件里），就拒绝该请求。此检查默认开启，可以关闭。
- 可选在模型回复中遮盖同类信息。
