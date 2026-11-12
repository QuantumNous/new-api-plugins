---
locale: zh-CN
---
# xAI

使用 xAI 的 Grok Imagine 模型生成、编辑和延长视频。

## 渠道设置

- 渠道类型选 xAI (48)。
- 密钥填 xAI 控制台创建的 API Key。
- API 地址留空时使用 `https://api.x.ai`。只有通过其他地址访问 xAI 时才需要填写，填写时去掉末尾的 `/v1`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装 xAI 插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos/generations` | 文生视频、图生视频和参考图生视频，使用 xAI 原生请求体 | `grok-imagine-video`、`grok-imagine-video-1.5`、`grok-imagine-video-1.5-lite` |
| `POST /v1/videos/edits` | 按提示词编辑视频 | `grok-imagine-video` |
| `POST /v1/videos/extensions` | 延长视频 | `grok-imagine-video` |
| `POST /v1/videos` | OpenAI 视频接口 | 以上三个模型 |
| `POST /v1/responses` | Responses API，支持流式、同步等待和后台三种调用方式 | 以上三个模型 |

xAI 原生路径按 xAI 的格式返回。这些路径加上 `/xai` 前缀也能调用，例如 `/xai/v1/videos/generations`。

原生路径的任务用 `GET /v1/videos/{request_id}` 轮询。返回内容保持 xAI 的格式，视频生成中返回 202，完成后带有视频地址。
`POST /v1/videos` 创建的任务在同一路径返回 OpenAI 视频对象，视频从 `GET /v1/videos/{task_id}/content` 下载。Responses 的结果里带有网关提供的视频链接。

## 计费

价格字段因模型而异。`grok-imagine-video` 有五个，1.5 系列模型有除输入视频单价以外的四个。

- “视频生成单价”按输出视频秒数计费。生成视频按请求的 `duration` 计，默认 8 秒，完成后按 xAI 报告的时长结算。编辑视频先按 9 秒预扣，完成后按输出时长结算。延长视频按新增的时长计，默认 6 秒。
- “输出视频分辨率”可以给每种分辨率分别定价。生成视频使用请求的 `resolution`，默认 `480p`。编辑和延长先按 `720p` 预扣，完成后按成品文件的分辨率结算。
- “输入图片单价”按请求里的图片张数计费，统计 `image`、`reference_images`、`last_frame` 以及 `keyframes` 中的图片。
- “输入视频单价”按输入视频秒数计费，只用于 `grok-imagine-video`。编辑先按 9 秒、延长先按 15 秒预扣，完成后按 xAI 结果推算出的输入时长结算。
- “被内容审核拦截”在 xAI 拦截了视频但仍然收费时为 true。网关对这类任务照常扣费，可以通过这个字段为它们单独定价。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos/generations \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "grok-imagine-video-1.5", "prompt": "A hot air balloon rising over a misty valley at dawn", "duration": 8, "resolution": "720p", "aspect_ratio": "16:9"}'
```

用返回的 `request_id` 轮询任务：

```bash
curl https://your-gateway.example.com/v1/videos/{request_id} \
  -H "Authorization: Bearer sk-..."
```

## 限制

- `grok-imagine-video` 支持 `480p` 和 `720p`，1.5 系列模型还支持 `1080p`。参考图生视频只支持 `480p` 和 `720p`。
- 生成视频时长为 1 到 15 秒，延长视频每次增加 2 到 10 秒。
- 图片和视频用 URL 或 base64 data URL 传入。带 `file_id` 或 `storage_options` 的请求会被拒绝。
