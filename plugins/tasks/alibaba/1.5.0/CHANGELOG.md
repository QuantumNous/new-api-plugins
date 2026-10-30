---
changelogVersion: 1
plugin: "alibaba"
version: "1.5.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.5.0]

### Added

- Call DashScope's own paths on the gateway: `/api/v1/services/aigc/...` for image and video generation and `GET /api/v1/tasks/{task_id}` for task queries. The `/ali/...` paths keep working.
- With the `X-DashScope-Sse: enable` header, `wan2.6-image` interleaved text-and-image output streams to the client as DashScope sends it.
- Image inputs can be DashScope temporary uploads (`oss://`). The gateway passes the client's `X-DashScope-OssResourceResolve` header on to DashScope.
- `wan3.0-video` and `wan3.0-video-prime` accept the `21:9` ratio.

### Changed

- Native requests reach DashScope as written, at the endpoint the client called, and replies keep DashScope's format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- Querying a task the gateway does not have returns 200 with `task_status: "UNKNOWN"`, as DashScope does.
- Once the gateway has failed a task, for example after a timeout, a native query shows it as failed even if DashScope's last reply still showed it pending or running.
- An interleaved stream that fails before producing any image is not charged. A stream cut off midway is charged the amount reserved when it started.

### Migration

- **Installation:** Install 1.5.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its DashScope paths.
