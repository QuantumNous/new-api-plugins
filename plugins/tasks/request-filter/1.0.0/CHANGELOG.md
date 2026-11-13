---
changelogVersion: 1
plugin: "request-filter"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Reject requests whose text contains a keyword or matches a regular expression, with up to 500 rules.
- Keep long word lists in a keyword list, one keyword per line.
- Check all text of chat, Responses, Claude and Gemini requests, the prompts of task and Midjourney submissions, and text, PDF and DOCX files sent with the request.
- Set the message clients see when a request is rejected.
- Optionally stop model replies that match a rule or the keyword list.
