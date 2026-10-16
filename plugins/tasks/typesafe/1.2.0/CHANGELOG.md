---
changelogVersion: 1
plugin: "typesafe"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Call Jev 1.14 as `jev-1.14.0`. `jev-preview` already points to it, and its input price is the same as Jev 1.13.
- Put images in the state, in question instructions, or in criteria. Upload them with the official TypeSafe SDKs or embed them as base64 in JSON. The plugin adds TypeSafe's image header to these requests, so clients can leave it out. The gateway bills image tokens as input tokens.
- Ask Box questions to find where something is in an image. They need `jev-1.14.0` or `jev-preview`.
- Send the `TypeSafe-Danger-Zone-Long-Context: true` header to use TypeSafe's long context of about 262k tokens. The gateway reserves quota for 262,400 input tokens and settles on the usage TypeSafe reports.
- Point the SDKs at the gateway root. The plugin now also serves `POST /v1/systemone`.

### Changed

- Score questions can have a single level. Earlier versions required at least two.
- Noul criteria can contain fields other than `true` and `false`. Earlier versions rejected these requests.
- A request with a top-level field that TypeSafe does not define now gets an error, as it does from TypeSafe. Earlier versions dropped the field and sent the rest.

### Fixed

- The gateway bills requests over 65,536 input tokens for every token TypeSafe reports. Earlier versions billed them as 65,536.
- Errors from TypeSafe reach clients in TypeSafe's format, such as `{"detail": {"error_type": "max_tokens_exceeded"}}`. Earlier versions put the error text in `detail` as a string.

### Migration

- Set a price for `jev-1.14.0` before offering it, using task usage expression billing like the other Jev models. At TypeSafe's current price of 0.042 USD per 1 million input tokens, the expression is `tier("base", u("input_tokens") * 0.042 / 1000000)`.
- Upgrade New API to a release after v1.0.0-rc.44 to upload images with the official SDKs, to use long context, or to send requests whose state, questions, or options use names such as `n`, `count`, or `duration`. Older gateways accept images only as base64 in JSON.
- If you reach TypeSafe through another New API gateway, upgrade that gateway and its TypeSafe plugin as well. A TypeSafe plugin older than 1.2.0 on that gateway drops images and rejects `jev-1.14.0` and Box questions.
