---
changelogVersion: 1
plugin: "request-redact"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Replace email addresses, phone numbers, ID card numbers, bank card numbers and API keys in request text before it is sent upstream. IPv4 addresses can be turned on as well.
- Add your own regular expressions, each with its own replacement text.
- Redact the text of chat, Responses, Claude and Gemini requests, the prompts of task and Midjourney submissions, and text files sent with the request.
- Before a request is sent, reject it if it still holds such information, for example in a PDF or DOCX file that cannot be changed. This check is on by default and can be turned off.
- Optionally mask the same information in model replies.
