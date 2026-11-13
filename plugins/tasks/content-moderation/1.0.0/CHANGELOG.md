---
changelogVersion: 1
plugin: "content-moderation"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Check user messages with an OpenAI-compatible moderation API and reject the requests it flags. By default it uses OpenAI's moderation API with the `omni-moderation-latest` model.
- Check the last user message or all of them, together with their images and their text, PDF and DOCX files.
- Check the prompts of task and Midjourney submissions, such as video or music generation.
- Reject only requests flagged for the categories you choose, or every flagged request.
- Optionally check the model's replies and stop the ones the moderation API flags.
