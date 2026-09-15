---
changelogVersion: 1
plugin: "hailuo"
version: "1.2.0"
locale: "zh-CN"
---
# Changelog

## [1.2.0]

### Changed

- 只接受各模型真正能出片的输出分辨率：`MiniMax-H3` 为 `768P`、`2K`，`MiniMax-Hailuo-2.3` 和 `MiniMax-Hailuo-2.3-Fast` 为 `768P`、`1080P`，`MiniMax-Hailuo-02` 为 `512P`、`768P`、`1080P`。向这些模型请求其他分辨率，例如给 `MiniMax-Hailuo-02` 传 `720P` 或 `2K`，现在会直接失败，不再悄悄按 `768P` 出片并计费。传 `1280x720` 这类像素尺寸仍会折算为最接近的受支持分辨率；01 系列模型（`T2V-01`、`T2V-01-Director`、`I2V-01`、`I2V-01-Director`、`I2V-01-live`、`S2V-01`）仍接受任意分辨率。

### Migration

- **价格配置：**已保存的价格继续生效，计费金额不变。每个模型现在只提供它真正会收取的价格项：`MiniMax-H3` 可配置视频秒数、输出分辨率、输入图片和输入视频秒数；`MiniMax-Hailuo-2.3`、`MiniMax-Hailuo-2.3-Fast`、`MiniMax-Hailuo-02` 可配置视频秒数和各自支持的分辨率；01 系列模型只按视频秒数计价。如果此前为 `MiniMax-H3` 以外的模型配置过输入图片或输入视频价格，下次保存该模型价格时，请把这部分金额折进按秒或按分辨率的价格，并删除该模型出不了的分辨率价格。
- **安装要求：**1.2.0 需要安装在支持按模型配置计费参数的 new-api 版本上，旧版网关会拒绝该插件。
