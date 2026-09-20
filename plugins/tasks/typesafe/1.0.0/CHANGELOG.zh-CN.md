---
changelogVersion: 1
plugin: "typesafe"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 通过 `POST /typesafe/v1/systemone` 调用 TypeSafe Jev，支持 `jev-1.13.0`、`jev-latest` 和 `jev-preview`。一次请求可组合 Choice、Score、Noul 问题，同步返回原始答案、概率和 token 用量。
- 输入单价按每 100 万（1M）token 设置，每次调用按 TypeSafe 返回的实际输入 token 数结算，输出 token 免费。
- 支持直连 TypeSafe，也可通过已安装此插件的另一台 New API 网关调用。

### Migration

- 为 Jev 模型启用任务用量表达式计费。可视化编辑器中的输入 token 单价以每 100 万（1M）token 为单位；选择美元时，填写 `0.042` 即对应 TypeSafe 当前公布的价格，也可设置自己的销售价格。对应表达式为 `tier("base", u("input_tokens") * 0.042 / 1000000)`。
- 新建任务插件渠道并绑定此插件，填入 TypeSafe API Key，默认上游地址为 `https://api.typesafe.ai`。使用 SDK 发起判断请求时，将 Base URL 设置为 `https://<gateway>/typesafe`，并使用网关令牌。
