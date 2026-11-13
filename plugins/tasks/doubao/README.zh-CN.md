---
locale: zh-CN
---
# Doubao

通过火山方舟调用豆包 Seedance 生成视频，调用豆包 Seedream 生成图片。

## 渠道设置

- 渠道类型选 Doubao (54)，选火山方舟 (45) 也可以。
- 密钥填火山方舟控制台创建的 API Key。
- API 地址留空时使用 `https://ark.cn-beijing.volces.com`。
- 可以在渠道的模型映射里把模型映射到方舟的推理接入点 ID，价格和参数仍按客户端请求的模型处理。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 生成视频。参考图以 URL 形式放在 `images` 里，其他参考素材放在 `metadata.content` 里。`resolution`、`ratio`、`generate_audio` 等方舟字段放在 `metadata` 里。 | Seedance 模型 |
| `POST /v1/images/generations` | 生成图片，图片生成完毕后一次返回。 | Seedream 模型 |
| `POST /v1/images/edits` | 编辑图片。参考图可以是 URL、data URL 或 multipart 上传的文件。 | Seedream 模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频或图片，支持流式、同步等待和后台三种调用方式。 | 下列全部模型 |
| `POST /api/v3/contents/generations/tasks` | 用方舟原生请求体生成视频。 | Seedance 模型 |
| `GET /api/v3/contents/generations/tasks/{task_id}` | 查询视频任务。 | Seedance 模型 |
| `POST /api/v3/images/generations` | 用方舟原生请求体生成图片。带 `stream: true` 时按方舟的推送顺序流式返回图片。 | Seedream 模型 |

Seedance 模型有 `doubao-seedance-1-0-pro-250528`、`doubao-seedance-1-0-lite-t2v`、`doubao-seedance-1-0-lite-i2v`、`doubao-seedance-1-5-pro-251215`、`doubao-seedance-2-0-260128`、`doubao-seedance-2-0-fast-260128`、`doubao-seedance-2-0-mini-260615` 和 `doubao-seedance-2-5-260628`。
Seedream 模型有 `doubao-seedream-5-5-261028`、`doubao-seedream-5-0-pro-260628`、`doubao-seedream-5-0-flash-260915`、`doubao-seedream-5-0-lite-260128`、`doubao-seedream-5-0-260128`、`doubao-seedream-4-5-251128` 和 `doubao-seedream-4-0-250828`。

视频客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频或图片链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
方舟原生路径也可以加 `/doubao` 前缀访问。

## 计费

Seedance 模型的计价字段如下。

- “计费 Token 单价”按 Token 数计费。任务开始时，网关按时长和分辨率估算 Token 数。时长为 `-1` 或未填写时按该模型最长的视频预扣：Seedance 1.0 和 1.5 pro 为 12 秒，Seedance 2.0 为 15 秒，Seedance 2.5 为 30 秒。视频生成后按方舟返回的 Token 数结算。
- “输出视频分辨率”可以给模型支持的每档分辨率分别设置价格。分辨率取自 `metadata.resolution` 或 `size`，都没有时按 `1080p` 预扣（Seedance 2.0 fast 和 mini 按 `720p`），完成后按方舟返回的分辨率结算。
- “参考视频输入”可以给有、无参考视频的请求分别设置价格，适用于 Seedance 2.0 和 2.5。
- “是否生成音频”可以给有声、无声视频分别设置价格，适用于 Seedance 1.5 pro。默认按有声计，请求把 `generate_audio` 设为 `false` 时按无声计。

Seedream 模型的计价字段如下。

- “图片生成单价（1.5K 及以下）”和“图片生成单价（1.5K 以上）”按输出图片张数计费，以 261 万像素为界。Seedream 5.0 lite、5.0 和 4.5 的输出都在这条线以上，只有一个“图片生成单价”。网关先按请求的尺寸和张数估算，完成后按每张图的实际尺寸结算。
- “输入图片单价”按参考图张数计费。网关先按请求里的图片数计，完成后按方舟返回的张数结算。
- “是否开启图层拆分”可以给图层拆分单独定价，适用于 Seedream 5.5、5.0 pro 和 flash。拆出的每一层都按一张输出图片计。

组图生成和图层拆分在图片返回前，按请求允许的最多张数预扣。图层拆分在 Seedream 5.0 pro 和 flash 上按 17 张预扣；在 Seedream 5.5 上按请求的 `layer_decomposition_options.max_images` 预扣，未设置时按 30 张。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "doubao-seedance-2-0-260128", "prompt": "A hot air balloon drifting over green hills at sunrise", "seconds": 5, "metadata": {"resolution": "1080p", "ratio": "16:9"}}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- Seedance 1.0、1.5 pro 和 2.5 支持 `480p`、`720p` 和 `1080p`，Seedance 2.0 另外支持 `4k`，Seedance 2.0 fast 和 mini 支持 `480p` 和 `720p`。请求其他分辨率会被拒绝。
- 时长设为 `-1` 时，Seedance 1.5 pro、2.0 和 2.5 会自己决定视频长度。
- `/v1/videos` 的参考素材必须是 URL，上传文件会被拒绝。
- 在 OpenAI Images API 上，`n` 只能为 1，需要多张图时用 `sequential_image_generation`。这个接口不支持流式返回和 `mask`。
- Seedream 5.0 pro 和 flash 生成透明背景（`background: "transparent"`）时，必须恰好有一张参考图。Seedream 5.5 没有这个限制，还可以传 `background: "auto"`。透明背景图片不能使用 `output_format: "jpeg"`。
- Seedream 5.5、5.0 pro 和 flash 最多接受 10 张参考图，其他 Seedream 模型最多 14 张。上传的单张图片不超过 30 MiB。
- 带 `callback_url` 的请求会被拒绝，请改为轮询任务。
