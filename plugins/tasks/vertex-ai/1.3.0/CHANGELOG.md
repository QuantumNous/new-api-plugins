---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.3.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.3.0]

### Added

- Generate videos with `gemini-omni-1.2-flash` through OpenAI Responses and the OpenAI video API, with integer durations from 3 to 15 seconds. Vertex AI limits Omni 1.1 and first/last-frame interpolation to 10 seconds.
- Send multiple images to Gemini Omni through OpenAI Responses. Two images provide the first and last frames; when more than two are supplied, they act as subject references. Images reach Vertex AI in their submitted order.

### Fixed

- Failed Gemini Omni tasks now show the error code and message from Vertex AI's error list when available, instead of only the task status.
- Responses requests validate every input image before submission, so an invalid image after the first is also rejected.

### Migration

- Before enabling `gemini-omni-1.2-flash`, configure its input, text output (including thinking), and video output token prices, each per million tokens.
