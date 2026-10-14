---
changelogVersion: 1
plugin: "vidu"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 支持将另一台 New API 作为上游，调用 Vidu 视频生成。两台服务器都需启用这个版本，并把插件绑定到填写了上游地址和 API 密钥的 New API 渠道。
- 可以直接按 Vidu 官方格式调用 `/vidu/ent/v2/text2video`、`/vidu/ent/v2/img2video`、`/vidu/ent/v2/start-end2video` 或 `/vidu/ent/v2/reference2video`，结果用 `/vidu/ent/v2/tasks/{task_id}/creations` 查询。

### Fixed

- Vidu 实际生成的模型、时长和分辨率现在与计费一致。此前 `metadata` 里的 `model`、`duration` 或 `resolution` 会在计费之后覆盖这些值，生成的视频可能换了模型，或比扣费的更长、更清晰。

### Migration

- **安装要求：**1.1.0 需要安装在 new-api v1.0.0-rc.39 或更高版本上，旧版本会拒绝该插件。
- **请求写法：**视频时长请用顶层的 `seconds` 或 `duration` 指定。`metadata` 里的 `duration` 和 `model` 现在会被忽略；`metadata` 里的 `resolution` 只在请求没有 `size` 或 `resolution` 时生效。
