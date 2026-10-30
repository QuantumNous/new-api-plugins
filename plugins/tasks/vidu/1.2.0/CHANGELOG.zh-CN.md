---
changelogVersion: 1
plugin: "vidu"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 新增 Vidu 当前的模型：`viduq4-preview`、`viduq3-pro`、`viduq3-turbo`、`viduq3-pro-fast`、`viduq3`、`viduq3-mix`、`viduq3-ad`、`viduq3-drama`、`viduq2-pro`、`viduq2-turbo`、`viduq2-pro-fast` 和 `viduq1-classic`，并在 Vidu 提供的模型上支持 2K 和 4K 分辨率。
- 可以在网关上直接调用 Vidu 官方路径：`POST /ent/v2/text2video`、`/ent/v2/img2video`、`/ent/v2/start-end2video` 和 `/ent/v2/reference2video`，用 `GET /ent/v2/tasks/{task_id}/creations` 查询。网关同时在 `/vidu` 下提供每个路径，原来的 `/vidu/ent/v2/...` 地址继续可用。
- 原生请求除了 `Bearer`，也可以像 Vidu 一样用 `Authorization: Token <key>` 鉴权。

### Changed

- 原生请求按原样发给 Vidu：插件不再自动加上 `movement_amplitude: "auto"`、`images: null` 或 `prompt: null`，也不再删除 `bgm: false`。响应保留 Vidu 的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 请求未填时长或分辨率时，按模型的默认值预扣；插件未列出的模型按 Vidu 最长的时长和最高的分辨率预扣。
- 带 `callback_url` 的请求会被拒绝，因为 Vidu 回调时用的任务 ID 不是网关发出的 ID。请改为查询任务。
- 网关判定任务失败后（例如超时），原生查询会显示任务失败，即使 Vidu 最后一次返回的状态仍是进行中。

### Migration

- **价格配置：**启用新模型前请先设置价格，并为支持 2K、4K 的模型补上这两档分辨率的价格。
- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 Vidu 路径。
