---
changelogVersion: 1
plugin: "hailuo"
version: "1.4.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.4.0]

### Added

- Call MiniMax's own paths on the gateway: `POST /v1/video_generation` and `GET /v1/query/video_generation?task_id=` for Hailuo 02, 2.3 and the 01 series, and `POST /v2/video_generation` and `GET /v2/query/video_generation/{task_id}` for `MiniMax-H3`. The `/hailuo/...` addresses of 1.3.0 keep working.
- `GET /v1/files/retrieve?file_id=` returns the download link of a finished video, as MiniMax does. The file ID is the task ID.

### Changed

- Native requests reach MiniMax as written: `MiniMax-H3` requests are no longer filled with defaults, and `aigc_watermark` is kept. Errors come back in MiniMax's `base_resp` format. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- A `MiniMax-H3` request without a duration reserves 15 seconds, the longest H3 renders, and settles on the duration MiniMax reports. Durations from 4 to 15 seconds are accepted.
- Requests with `callback_url` are refused, because MiniMax would call back with task IDs the gateway does not hand out. Query the task instead.
- Once the gateway has failed a task, for example after a timeout, a native query shows it as failed even if MiniMax's last reply still showed it running.

### Migration

- **Installation:** Install 1.4.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its MiniMax paths.
