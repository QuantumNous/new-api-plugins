---
changelogVersion: 1
plugin: "vidu"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Add Vidu's current models: `viduq4-preview`, `viduq3-pro`, `viduq3-turbo`, `viduq3-pro-fast`, `viduq3`, `viduq3-mix`, `viduq3-ad`, `viduq3-drama`, `viduq2-pro`, `viduq2-turbo`, `viduq2-pro-fast` and `viduq1-classic`, with the 2K and 4K resolutions where Vidu offers them.
- Call Vidu's own paths on the gateway: `POST /ent/v2/text2video`, `/ent/v2/img2video`, `/ent/v2/start-end2video` and `/ent/v2/reference2video`, queried at `GET /ent/v2/tasks/{task_id}/creations`. The gateway also serves each path under `/vidu`, so existing `/vidu/ent/v2/...` addresses keep working.
- Native requests can authenticate with `Authorization: Token <key>`, as Vidu takes it, as well as `Bearer`.

### Changed

- Native requests reach Vidu as written: the plugin no longer adds `movement_amplitude: "auto"`, `images: null` or `prompt: null`, and no longer drops `bgm: false`. Replies keep Vidu's format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- When a request leaves out the duration or resolution, the reservation uses the model's defaults. For a model the plugin does not list, it uses the longest duration and highest resolution Vidu offers.
- Requests with `callback_url` are refused, because Vidu would call back with task IDs the gateway does not hand out. Query the task instead.
- Once the gateway has failed a task, for example after a timeout, a native query shows it as failed even if Vidu's last reply still showed it running.

### Migration

- **Price configuration:** Set prices for the new models before enabling them, and add 2K and 4K prices for the models that offer those resolutions.
- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Vidu paths.
