---
changelogVersion: 1
plugin: "typesafe"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Call TypeSafe Jev through `POST /typesafe/v1/systemone` with `jev-1.13.0`, `jev-latest`, or `jev-preview`. Send Choice, Score, and Noul questions together and receive the original answers, probabilities, and token usage in a synchronous response.
- Set the input price per 1 million (1M) tokens. Each evaluation is billed for the actual input tokens reported by TypeSafe; output tokens are free.
- Connect directly to TypeSafe or through another New API gateway with this plugin installed.

### Migration

- Configure the Jev models with task usage expression billing. In the visual editor, the input token price is per 1 million (1M) tokens. With USD selected, enter `0.042` to match TypeSafe's current published price, or choose your own selling price. The equivalent expression is `tier("base", u("input_tokens") * 0.042 / 1000000)`.
- Bind the plugin to a Task Plugin channel with your TypeSafe API key. The default upstream address is `https://api.typesafe.ai`. For SDK evaluation calls, set the base URL to `https://<gateway>/typesafe` and use a gateway API token.
