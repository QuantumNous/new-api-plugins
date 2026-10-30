---
changelogVersion: 1
plugin: "doubao"
version: "1.4.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.4.0]

### Added

- Call Ark's own paths on the gateway: `POST /api/v3/contents/generations/tasks` and `GET /api/v3/contents/generations/tasks/{task_id}` for video tasks, and `POST /api/v3/images/generations` for Seedream images. The gateway also serves each path under `/doubao`, so existing `/doubao/api/v3/...` addresses keep working.
- Native Seedream requests with `stream: true` stream the images to the client as Ark sends them. They are charged by the images delivered, and a stream that fails before any image is not charged.
- Native Seedream requests can ask for `response_format: "b64_json"`. OpenAI Responses requests still return image links.

### Changed

- Native requests reach Ark as written. The `content` array is sent as it is instead of being merged into the prompt, and Ark checks `sequential_image_generation` and `stream` itself. Replies keep Ark's format, status codes and error codes, such as `InvalidParameter` and `AuthenticationError`. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- A video task query returns Ark's last reply for the task. Once the gateway has failed a task, for example after a timeout, the query shows it as failed even if Ark's last reply still showed it running.
- Requests with `callback_url` are refused, because Ark would call back with task IDs the gateway does not hand out. Query the task instead.

### Migration

- **Installation:** Install 1.4.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Ark paths.
