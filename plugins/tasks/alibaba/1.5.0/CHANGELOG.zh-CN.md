---
changelogVersion: 1
plugin: "alibaba"
version: "1.5.0"
locale: "zh-CN"
---
# Changelog

## [1.5.0]

### Added

- 可以在网关上直接调用 DashScope 官方路径：图片和视频生成用 `/api/v1/services/aigc/...`，查询任务用 `GET /api/v1/tasks/{task_id}`。原来的 `/ali/...` 路径继续可用。
- 请求带上 `X-DashScope-Sse: enable` 时，`wan2.6-image` 的图文交错输出会随 DashScope 的返回实时流式发给客户端。
- 输入图片可以使用 DashScope 临时上传的文件（`oss://`）。网关会把客户端的 `X-DashScope-OssResourceResolve` 请求头转发给 DashScope。
- `wan3.0-video` 和 `wan3.0-video-prime` 支持 `21:9` 比例。

### Changed

- 原生请求按原样发到客户端调用的 DashScope 接口，响应保留 DashScope 的格式、状态码和错误码。只有模型缺失或未知、请求无法计费、或用到网关不支持的功能时，网关才自己拒绝请求。
- 查询网关中不存在的任务时，返回 200 和 `task_status: "UNKNOWN"`，与 DashScope 一致。
- 网关判定任务失败后（例如超时），原生查询会显示任务失败，即使 DashScope 最后一次返回的状态仍是排队中或运行中。
- 图文交错流在生成任何图片之前失败时不收费；中途断开的流按开始时预扣的金额收费。

### Migration

- **安装要求：**1.5.0 需要安装在 v1.0.0-rc.47 之后发布的 new-api 版本上。旧版本会拒绝该插件，或无法提供它的 DashScope 路径。
