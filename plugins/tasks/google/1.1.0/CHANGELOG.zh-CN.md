---
changelogVersion: 1
plugin: "google"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Added

- 支持将另一台 New API 作为上游，调用 Gemini API 上的 Google Veo 视频生成。两台服务器都需启用这个版本，并把插件绑定到填写了上游地址和 API 密钥的 New API 渠道。
- 可以直接按 Gemini API 格式调用 `/google/v1beta/models/{model}:predictLongRunning`，传入提示词和可选的首帧 `image`。返回的 operation 名为 `models/{model}/operations/{task_id}`，用 `/google/v1beta/models/{model}/operations/{task_id}` 查询进度。

### Migration

- **安装要求：**1.1.0 需要安装在 new-api v1.0.0-rc.39 或更高版本上，旧版本会拒绝该插件。
