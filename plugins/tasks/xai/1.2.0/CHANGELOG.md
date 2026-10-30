---
changelogVersion: 1
plugin: "xai"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Call xAI's own paths on the gateway: `POST /v1/videos/generations`, `/v1/videos/edits` and `/v1/videos/extensions`, queried at `GET /v1/videos/{request_id}`. The gateway also serves each path under `/xai`, so existing `/xai/v1/videos/...` addresses keep working.
- Add the `21:9` and `5:2` aspect ratios.

### Changed

- Native requests reach xAI as written, and replies keep xAI's format, status codes and error codes. The plugin no longer limits how many `reference_images`, `keyframes` or `reference_audios` a request carries; xAI checks them. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- A query for a video still in progress returns 202, as xAI does.
- A failed video that xAI still charged keeps xAI's `usage.cost_in_usd_ticks` in the query reply.
- Requests with `storage_options` are refused. Before, the field was dropped without notice.
- Once the gateway has failed a task, for example after a timeout, a native query shows it as failed even if xAI's last reply still showed it pending.

### Migration

- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its xAI paths.
