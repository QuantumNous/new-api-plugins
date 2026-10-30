---
changelogVersion: 1
plugin: "jimeng"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Call Jimeng's Bearer API on the gateway: `POST /api/common/v3/submit_task_v2` to submit a task and `POST /api/common/v3/get_result_v2` to query it, with the bodies of the `CVSync2AsyncSubmitTask` and `CVSync2AsyncGetResult` actions. The gateway also serves each path under `/jimeng`.

### Changed

- Native requests reach Jimeng as written: an omitted `seed` stays Jimeng's random default instead of 0, and image inputs are no longer rewritten. Replies keep Jimeng's format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- A `req_key` the plugin cannot price is rejected instead of being charged as S2.0 Pro.
- `frames` accepts only 121 (5 seconds) or 241 (10 seconds).
- Queries with `req_json` are refused, because the gateway answers from its last poll of the task.
- Tasks Jimeng reports as `not_found` or `expired` read as failed. Once the gateway has failed a task, for example after a timeout, a query shows it as failed even if Jimeng's last reply still showed it running.

### Removed

- The `POST /jimeng/?Action=...` address.

### Migration

- **Requests:** Move native calls from `/jimeng/?Action=CVSync2AsyncSubmitTask` and `/jimeng/?Action=CVSync2AsyncGetResult` to `/api/common/v3/submit_task_v2` and `/api/common/v3/get_result_v2`, or the same paths under `/jimeng`. The request bodies stay the same.
- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Jimeng paths.
