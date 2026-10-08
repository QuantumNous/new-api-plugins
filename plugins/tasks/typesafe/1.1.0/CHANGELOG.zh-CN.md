---
changelogVersion: 1
plugin: "typesafe"
version: "1.1.0"
locale: "zh-CN"
---
# Changelog

## [1.1.0]

### Fixed

- 请求转发给 TypeSafe 时，`state`、`questions` 及其中所有嵌套对象的字段顺序与你发送的一致。之前的版本会把对象字段按字母排序，改变了 Jev 读到的内容，可能导致打分与直接调用 TypeSafe 不同。

### Migration

- 安装此版本前，先把 New API 升级到 v1.0.0-rc.41 之后的版本，更早的网关无法加载此版本。
- 如果通过另一台 New API 网关调用 TypeSafe，也要升级那台网关及其 TypeSafe 插件，否则字段顺序仍会在那里被打乱。
