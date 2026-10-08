---
changelogVersion: 1
plugin: "xai"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Generate videos with xAI Grok Imagine using `grok-imagine-video`, `grok-imagine-video-1.5`, or `grok-imagine-video-1.5-lite`: from a text prompt, from an image, or from reference images, with optional keyframes and reference audio.
- Edit a video or extend it with `grok-imagine-video` through `POST /xai/v1/videos/edits` and `POST /xai/v1/videos/extensions`.
- Call the native xAI routes under `/xai/v1/videos`, or use the OpenAI-compatible `/v1/videos` and `/v1/responses` endpoints, including streaming and background responses.
- Each video is billed by its output seconds at its model and resolution, by its input images, and on `grok-imagine-video` by the seconds of input video in an edit or extension.
- Videos that xAI generates but then blocks for moderation, and failed requests that xAI still bills, are charged instead of refunded, because xAI charges for them.
- Serve the models from an xAI channel with its existing API key, or through another New API gateway with this plugin installed.

### Migration

- Configure the Grok Imagine models with task usage expression billing: a price per output second for each resolution, a price per input image, and for `grok-imagine-video` a price per second of input video. Edits and extensions bill under the "Same as input video" resolution because xAI does not report it, and blocked videos can be given their own price.
