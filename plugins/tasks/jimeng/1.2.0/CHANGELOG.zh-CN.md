---
changelogVersion: 1
plugin: "jimeng"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 可以在网关上直接调用即梦的 Bearer 接口：用 `POST /api/common/v3/submit_task_v2` 提交任务，用 `POST /api/common/v3/get_result_v2` 查询结果，请求体与 `CVSync2AsyncSubmitTask`、`CVSync2AsyncGetResult` 两个 Action 相同。网关同时在 `/jimeng` 下提供每个路径。

### Changed

- 原生请求按原样发给即梦：未填 `seed` 时沿用即梦的随机默认值，不再固定为 0；输入图片也不再被改写。响应保留即梦的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 插件无法计费的 `req_key` 会被拒绝，不再按 S2.0 Pro 收费。
- `frames` 只接受 121（5 秒）或 241（10 秒）。
- 带 `req_json` 的查询会被拒绝，因为网关返回的是它最后一次轮询到的结果。
- 即梦报告为 `not_found` 或 `expired` 的任务显示为失败。网关判定任务失败后（例如超时），查询也会显示失败，即使即梦最后一次返回的状态仍是进行中。

### Removed

- `POST /jimeng/?Action=...` 地址。

### Migration

- **请求方式：**原生调用请从 `/jimeng/?Action=CVSync2AsyncSubmitTask` 和 `/jimeng/?Action=CVSync2AsyncGetResult` 改为 `/api/common/v3/submit_task_v2` 和 `/api/common/v3/get_result_v2`，也可以用 `/jimeng` 下的同名路径。请求体不变。
- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的即梦路径。
