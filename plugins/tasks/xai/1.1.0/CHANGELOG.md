---
changelogVersion: 1
plugin: "xai"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Changed

- Edits and extensions on `grok-imagine-video` are now charged at the 480p or 720p price of the video they produce, as xAI prices them. An edited or extended video keeps the input video's resolution, up to 720p, so the 720p price is held when the task starts and the task settles at 480p when the finished video is no larger than 480p.
- To find that resolution, the plugin reads the beginning of the finished video from `vidgen.x.ai`.

### Removed

- The "Edit or extend video" row of the `grok-imagine-video` pricing table. The column is "Output video resolution" again, with 480p and 720p rows.

### Migration

- **Price configuration:** Edits and extensions on `grok-imagine-video` no longer use the "Edit or extend video" price. If that price differed from your 720p price, review the 480p and 720p prices, because edits and extensions are now charged at them.
- **Installation:** Settling edits and extensions at 480p needs a new-api version that lets plugins read upstream data (`utils.fetch`). On older versions they are charged at the 720p price.
