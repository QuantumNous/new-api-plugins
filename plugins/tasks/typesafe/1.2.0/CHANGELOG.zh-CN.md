---
changelogVersion: 1
plugin: "typesafe"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Added

- 支持 Jev 1.14，模型 ID 为 `jev-1.14.0`。`jev-preview` 已经指向它，输入价格和 Jev 1.13 相同。
- 请求可以带图片。图片可以放在 state、问题的 instructions 或 criteria 里，用官方 TypeSafe SDK 上传，或以 base64 写进 JSON。插件会给这类请求加上 TypeSafe 的图片请求头，客户端可以省略它。图片按输入 token 计费。
- 支持 Box 问题，用来找出目标在图片中的位置，需要使用 `jev-1.14.0` 或 `jev-preview`。
- 请求带上 `TypeSafe-Danger-Zone-Long-Context: true` 请求头，就能使用 TypeSafe 约 262k token 的长上下文。网关先按 262,400 个输入 token 预扣额度，再按 TypeSafe 返回的用量结算。
- 插件现在同时提供 `POST /v1/systemone`，SDK 可以直接把网关根地址当作 base URL。

### Changed

- Score 问题可以只有一个等级，之前的版本要求至少两个。
- Noul 的 criteria 可以带 `true` 和 `false` 以外的字段，之前的版本会拒绝这类请求。
- 请求顶层带了 TypeSafe 未定义的字段时，现在会报错，和直接请求 TypeSafe 一样。之前的版本会删掉这个字段，把其余内容照常发出。

### Fixed

- 输入超过 65,536 个 token 的请求，现在按 TypeSafe 报告的全部 token 计费。之前的版本只按 65,536 计费。
- TypeSafe 返回的错误按 TypeSafe 的原始格式交给客户端，例如 `{"detail": {"error_type": "max_tokens_exceeded"}}`。之前的版本会把错误内容当作字符串放进 `detail`。

### Migration

- 提供 `jev-1.14.0` 前先为它定价，和其他 Jev 模型一样使用任务用量表达式计费。按 TypeSafe 当前每 100 万输入 token 0.042 美元的价格，表达式为 `tier("base", u("input_tokens") * 0.042 / 1000000)`。
- 要用官方 SDK 上传图片、使用长上下文，或者发送 state、问题、选项里带 `n`、`count`、`duration` 这类字段名的请求，需要把 New API 升级到 v1.0.0-rc.44 之后的版本。更早的网关只接受以 base64 写进 JSON 的图片。
- 如果通过另一台 New API 网关调用 TypeSafe，那台网关和它的 TypeSafe 插件也要一起升级。那台网关上 1.2.0 之前的 TypeSafe 插件会丢掉图片，也会拒绝 `jev-1.14.0` 和 Box 问题。
