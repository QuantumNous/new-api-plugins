---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.3.1"
locale: "zh-CN"
---
# Changelog

## [1.3.1]

### Fixed

- 视频请求改为在渠道的部署地区运行，不再固定使用 `global`。
- 部署地区填 JSON 时，用法和渠道里的其他模型一样：每个模型名对应自己的地区，其余模型用 `default`。按客户端请求的模型名查找，在模型映射之前。

### Migration

- 升级前请检查 Vertex AI 渠道的部署地区。视频请求现在会发到这个地区，所用的视频模型必须在该地区可用。要让某个模型继续在 `global` 运行，在 JSON 里把它对应到 `global`。
- **安装要求：**部署地区在 v1.0.0-rc.49 之后发布的 new-api 版本上才生效。在 v1.0.0-rc.49 及更早版本上，请求仍在 `global` 运行。
