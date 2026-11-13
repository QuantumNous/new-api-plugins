---
changelogVersion: 1
plugin: "param-override"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 对所有请求或指定分组应用参数覆盖规则，语法与渠道的参数覆盖相同。
- 可以设置、删除或改写请求体字段和请求头，并按模型、分组、渠道、路径或请求格式设置条件。
- 可以用 `return_error` 规则拒绝请求，客户端收到 HTTP 400 和规则里的提示。
- 渠道自己的参数覆盖在这些规则之后运行，同一字段或请求头以渠道的设置为准。
