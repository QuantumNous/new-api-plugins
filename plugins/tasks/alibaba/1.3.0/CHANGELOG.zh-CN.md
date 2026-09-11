---
changelogVersion: 1
plugin: "alibaba"
version: "1.3.0"
locale: "zh-CN"
---
# Changelog

## [1.3.0]

### Added

- 阿里百炼从视频任务扩展为图片与视频生成，新增十个图片模型。其中 `wan2.7-image-pro`、`wan2.7-image` 支持文生图和参考图编辑，最多接收九张参考图。
- 新增 `wan2.6-t2i` 文生图，以及 `wan2.6-image` 图片编辑和图文混排。普通编辑最多接收四张参考图，图文混排最多接收一张。
- 新增 `wan2.5-t2i-preview`、`wan2.2-t2i-flash`、`wan2.2-t2i-plus`、`wanx2.1-t2i-turbo`、`wanx2.1-t2i-plus`、`wanx2.0-t2i-turbo` 的异步文生图支持。
- 新增原生异步提交入口：Wan 2.6/2.7 图片模型使用 `/ali/api/v1/services/aigc/image-generation/generation`，旧版图片模型使用 `/ali/api/v1/services/aigc/text2image/image-synthesis`。请求沿用厂商的 `input`、`parameters` 对象，通过 `/ali/api/v1/tasks/:task_id` 查询结果。
- 新增原生同步图片入口 `/ali/api/v1/services/aigc/multimodal-generation/generation`，在提交响应中以 JSON 返回已完成的图片结果。
- OpenAI Responses 扩展到全部十个图片模型，支持 `stream`、`sync`、`background` 三种请求模式、文本和参考图输入及图片生成参数。Responses 图片请求默认异步调用上游，也可显式选择同步调用。
- 同步普通出图支持一至四张，默认一张。Wan 2.7 通过 `enable_sequential` 支持最多十二张组图；异步组图默认十二张，同步请求未指定 `n` 时仍默认一张。
- Wan 2.6 通过 `enable_interleave` 支持图文混排，`max_images` 可设为一至五张，默认五张。原生同步请求以聚合 JSON 按顺序返回完整的文字和图片结果。
- 普通 `wan2.6-image` 和 Wan 2.7 图片请求支持 `1K`、`2K` 尺寸预设；`wan2.7-image-pro` 在无参考图且未开启组图的文生图场景下额外支持 `4K`。显式像素尺寸接受 `width*height` 或 `widthxheight`，并按模型检查总像素和宽高比限制。
- 透传负向提示词、提示词扩展、水印、种子、组图、图文混排、思考模式、边界框和调色板等图片参数。合并兼容字段和原生参数时保留显式 `seed: 0` 及值为 `false` 的选项。
- 每张成功图片分别提供 `image-1`、`image-2` 等任务产物和独立内容下载。Responses 按原始顺序展示文字和图片链接。

### Changed

- 同步和异步图片任务均按实际成功图片数报告用量。优先使用上游 `usage.image_count`，缺失时统计返回图片；部分成功响应只计入成功图片，不计入失败结果项。
- 提交前校验图片数量、图文混排上限、布尔选项、参考图数量、提示词结构和尺寸限制，也覆盖 `metadata.parameters` 中传入的值。拒绝通过 metadata 改写已选模型、对旧版图片模型发起同步调用，以及在同步 Wan 2.6 图文混排之外启用上游流式输出。
- 拒绝未完成或无图片的普通同步结果。异步图片任务虽标记成功但没有图片输出时按失败处理；Wan 2.6 图文混排仅返回文本时仍可成功，并报告零图片用量。
- 拒绝将非数字、负数、小数、超限数值，以及已有图片但计数为零等非法计数作为有效图片用量返回。

### Migration

- **价格配置：**需要按实际成功图片数计费时，为启用的图片模型配置使用 `u("image_count")` 的按张任务用量定价，必要时从按次计费切换。升级不会自动修改已保存的价格或表达式，现有视频价格和表达式无需仅因此版本而调整。
- **安装要求：**本版本需要配套支持的 new-api。安装 1.3.0 前应先升级不兼容的宿主，并按[插件 API 兼容要求](../../../../docs/plugin-api/v1.md)确认兼容性。
- **调用模式：**Responses 图片请求默认异步调用上游，需要同步调用时设置 `metadata.upstream_mode: "sync"`；原生同步请求使用新增的 multimodal-generation 路由。旧版图片模型仅支持异步调用。依赖特定出图数量时请显式设置 `n`：同步调用和 Wan 2.7 普通异步调用默认一张，Wan 2.6 和旧版模型的普通异步调用默认四张，Wan 2.7 异步组图默认十二张。
- **图文混排：**同步 Wan 2.6 图文混排需在 Responses 请求中设置 `metadata.parameters.enable_interleave: true`，或在原生同步入口设置 `parameters.enable_interleave: true`，并按需指定 `max_images`。插件自动启用上游流式输出。原生请求不得设置顶层 `stream: true`；原生响应为聚合 JSON，上游事件帧不会直接转发给客户端。
