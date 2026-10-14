---
changelogVersion: 1
plugin: "hailuo"
version: "1.3.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.3.0]

### Added

- Use another New API server as the upstream for MiniMax Hailuo video generation, including `MiniMax-H3`. Enable this version on both servers and bind the plugin to a New API channel that holds the upstream address and API key.
- Send requests in MiniMax's own format to `/hailuo/v1/video_generation` (Hailuo 02, 2.3 and the 01 series) or `/hailuo/v2/video_generation` (`MiniMax-H3`). Query a task at `/hailuo/v1/query/video_generation/{task_id}` or `/hailuo/v2/query/video_generation/{task_id}`; the v1 query takes the task ID in the path instead of a `task_id` query parameter.

### Migration

- **Installation:** Install 1.3.0 on new-api v1.0.0-rc.43 or later. Older versions reject the plugin or cannot download finished videos.
