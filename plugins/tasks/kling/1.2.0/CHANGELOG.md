---
changelogVersion: 1
plugin: "kling"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Add `kling-v2-5-turbo`, `kling-v2-6` and `kling-v3`. They are charged in Kling resource units per second by mode, with the higher rates Kling sets for native audio (`sound: "on"`) and reference voices (`voice_list`).
- Add Kling 4.0 (`kling-4.0`) at Kling's own paths: `POST /text-to-video/kling-4.0`, `POST /image-to-video/kling-4.0` and `POST /omni-video/kling-4.0`, queried at `GET /tasks?task_ids=`. It is also available through OpenAI Responses and the OpenAI video API.
- Kling 4.0 costs 1.0 units per second at 720p and 1.5 at 1080p. With an input video it costs 0.6 or 0.9 units per second of twice the output length, or of input plus output when the input is longer. A request without a duration reserves 30 seconds, and the task settles on the units Kling deducts, or else on the durations Kling reports.
- Call Kling's own paths for the other models on the gateway: `POST /v1/videos/text2video` and `POST /v1/videos/image2video`, queried at `GET /v1/videos/text2video/{task_id}` and `GET /v1/videos/image2video/{task_id}`. The gateway also serves each path under `/kling`, so existing `/kling/v1/videos/...` addresses keep working.

### Changed

- Native requests reach Kling as written and need `model_name`. The plugin no longer adds `aspect_ratio: "1:1"`, `cfg_scale: 0.5` or `model`. Replies keep Kling's format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- Durations must be whole seconds: 3 to 30 for Kling 4.0, and at most 3600 for the other models. A duration outside the range is rejected before anything is charged.
- On OpenAI Responses and the OpenAI video API, a text-to-video `size` maps to the nearest ratio Kling offers: 21:9, 16:9, 9:16 or 1:1 for Kling 4.0, and 16:9, 9:16 or 1:1 for the other models. Without a size, Kling uses its default, 16:9. Image-to-video sends no ratio.
- For Kling 4.0, `size` also sets the resolution: 720p when the shorter side is under 900 pixels, otherwise 1080p. A `size` not written as WIDTHxHEIGHT, such as `1280x720`, is rejected.
- Requests with `callback_url` or `external_task_id` are refused, because Kling would call back with task IDs the gateway does not hand out. Query the task instead.
- Once the gateway has failed a task, for example after a timeout, a native query shows it as failed even if Kling's last reply still showed it running.

### Removed

- `kling-v1`, `kling-v1-6` and `kling-v2-master`, which Kling shut down on 2026-09-15.

### Migration

- **Price configuration:** Set a unit price for `kling-v2-5-turbo`, `kling-v2-6`, `kling-v3` and `kling-4.0` before enabling them. Tasks are charged by the Kling units they use, at that price.
- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Kling paths.
