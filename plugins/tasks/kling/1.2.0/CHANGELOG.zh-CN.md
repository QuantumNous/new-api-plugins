---
changelogVersion: 1
plugin: "kling"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 新增 `kling-v2-5-turbo`、`kling-v2-6` 和 `kling-v3`。按模式以可灵资源包单位按秒计费，开启原生音频（`sound: "on"`）或使用参考音色（`voice_list`）时按可灵规定的更高单价计费。
- 新增可灵 4.0（`kling-4.0`），可以直接调用可灵官方路径 `POST /text-to-video/kling-4.0`、`POST /image-to-video/kling-4.0` 和 `POST /omni-video/kling-4.0`，用 `GET /tasks?task_ids=` 查询。也可以通过 OpenAI Responses 和 OpenAI 视频接口调用。
- 可灵 4.0 的消耗为 720p 每秒 1.0 单位、1080p 每秒 1.5 单位。带输入视频时为每秒 0.6 或 0.9 单位，时长按输出的两倍计算；输入视频比输出长时，按输入加输出计算。未指定时长的请求先按 30 秒预扣。结算按可灵的扣费：资源包扣的单位数，或现金支付的原价（list_price）。没有扣费数据时按可灵报告的时长结算。
- 其他模型也可以在网关上直接调用可灵官方路径：`POST /v1/videos/text2video` 和 `POST /v1/videos/image2video`，用 `GET /v1/videos/text2video/{task_id}` 和 `GET /v1/videos/image2video/{task_id}` 查询。网关同时在 `/kling` 下提供每个路径，原来的 `/kling/v1/videos/...` 地址继续可用。

### Changed

- 原生请求按原样发给可灵，必须填写 `model_name`。插件不再自动加上 `aspect_ratio: "1:1"`、`cfg_scale: 0.5` 或 `model`。响应保留可灵的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 时长必须是整数秒：可灵 4.0 为 3 到 30 秒，其他模型最长 3600 秒。超出范围的时长会在扣费前直接被拒绝。
- 通过 OpenAI Responses 和 OpenAI 视频接口文生视频时，`size` 会映射到可灵支持的最接近的比例：可灵 4.0 为 21:9、16:9、9:16 或 1:1，其他模型为 16:9、9:16 或 1:1。不传尺寸时，可灵使用默认的 16:9。图生视频不发送比例。
- 可灵 4.0 还会按 `size` 决定清晰度：短边小于 900 像素为 720p，否则为 1080p。不是“宽x高”格式（例如 `1280x720`）的 `size` 会被拒绝。
- 带 `callback_url` 或 `external_task_id` 的请求会被拒绝，因为可灵回调时用的任务 ID 不是网关发出的 ID。请改为查询任务。
- 网关判定任务失败后（例如超时），原生查询会显示任务失败，即使可灵最后一次返回的状态仍是进行中。

### Removed

- `kling-v1`、`kling-v1-6` 和 `kling-v2-master`，可灵已于 2026-09-15 下线这些模型。

### Migration

- **价格配置：**启用 `kling-v2-5-turbo`、`kling-v2-6`、`kling-v3` 和 `kling-4.0` 前，请为它们设置单位单价。任务按消耗的可灵单位数乘以该单价收费。
- **安装要求：**1.2.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的可灵路径。
