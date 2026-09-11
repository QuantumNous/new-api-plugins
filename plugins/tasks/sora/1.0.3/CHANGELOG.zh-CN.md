---
changelogVersion: 1
plugin: "sora"
version: "1.0.3"
locale: "zh-CN"
---
# Changelog

## [1.0.3]

### Fixed

- 当 `task.data` 为对象时，OpenAI Video 查询呈现保留最新持久化上游响应中的字段，包括输出地址和提供商元数据。宿主仍会替换标准任务字段；没有对象快照的任务继续使用原有呈现逻辑。
