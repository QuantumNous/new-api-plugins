---
locale: zh-CN
---
# Alibaba Bailian

通过阿里云百炼调用万相生成视频，调用万相、千问图像和 Z-Image 生成图片。

## 渠道设置

- 渠道类型选阿里百炼 (17)。
- 密钥填百炼控制台创建的 DashScope API Key。
- API 地址留空时使用 `https://dashscope.aliyuncs.com`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 生成视频。图片以 URL 形式放在 `image`、`images` 或 `input_reference` 里。 | 万相视频模型 |
| `POST /v1/images/generations` | 生成图片，图片生成完毕后一次返回。 | 图片模型 |
| `POST /v1/images/edits` | 编辑图片。输入图可以是 URL、data URL 或 multipart 上传的文件。 | 图片模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频或图片，支持流式、同步等待和后台三种调用方式。 | 下列全部模型 |
| `POST /api/v1/services/aigc/video-generation/video-synthesis`、`POST /api/v1/services/aigc/image2video/video-synthesis` | 用 DashScope 原生请求体生成视频。 | 万相视频模型 |
| `POST /api/v1/services/aigc/multimodal-generation/generation` | 用 DashScope 原生请求体生成图片，生成完毕后直接返回。 | `wan2.7-image-pro`、`wan2.7-image`、`wan2.6-image`、`wan2.6-t2i`、千问图像和 Z-Image 模型 |
| `POST /api/v1/services/aigc/image-generation/generation`、`POST /api/v1/services/aigc/text2image/image-synthesis`、`POST /api/v1/services/aigc/image2image/image-synthesis` | 用 DashScope 原生请求体以任务方式生成或编辑图片。 | 万相图片模型 |
| `GET /api/v1/tasks/{task_id}` | 查询 DashScope 任务。 | 万相视频和图片模型 |

万相视频模型有 `wan3.0-video`、`wan3.0-video-prime`、`wan2.7-t2v`、`wan2.7-i2v`、`wan2.6-t2v`、`wan2.6-t2v-us`、`wan2.6-i2v`、`wan2.6-i2v-flash`、`wan2.6-i2v-us`、`wan2.5-t2v-preview`、`wan2.5-i2v-preview`、`wan2.2-t2v-plus`、`wan2.2-i2v-flash`、`wan2.2-i2v-plus`、`wan2.2-kf2v-flash`、`wan2.2-s2v`、`wanx2.1-t2v-plus`、`wanx2.1-t2v-turbo`、`wanx2.1-i2v-plus` 和 `wanx2.1-i2v-turbo`。
万相图片模型有 `wan2.7-image-pro`、`wan2.7-image`、`wan2.6-image`、`wan2.6-t2i`、`wan2.5-t2i-preview`、`wan2.2-t2i-flash`、`wan2.2-t2i-plus`、`wanx2.1-t2i-turbo`、`wanx2.1-t2i-plus`、`wanx2.0-t2i-turbo`、`wan2.5-i2i-preview` 和 `wanx2.1-imageedit`。
千问图像模型有 `qwen-image`、`qwen-image-plus`、`qwen-image-max`、`qwen-image-2.0`、`qwen-image-2.0-pro`、`qwen-image-3.0`、`qwen-image-3.0-pro`、`qwen-image-edit`、`qwen-image-edit-plus` 和 `qwen-image-edit-max`。Z-Image 模型是 `z-image-turbo`。
带日期的快照版本（例如 `wan2.7-t2v-2026-04-25`、`qwen-image-2.0-pro-2026-03-03`）与对应的基础模型用法相同。

视频客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频或图片链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
DashScope 原生路径也可以加 `/ali` 前缀访问。

## 计费

万相视频模型的计价字段如下。

- “视频生成单价”按视频秒数计费。秒数取自请求，未填写时按 5 秒计。`wan2.2-s2v` 的时长跟随音频，先按 20 秒预扣。万相 3.0 在时长为 `-1` 或带参考视频时按 30 秒预扣，计费秒数为输入视频与输出视频的秒数之和。视频生成后按 DashScope 返回的秒数结算。
- “输出视频分辨率”可以给模型支持的每档分辨率分别设置价格。分辨率取自 `resolution` 或 `size`，都没有时按模型默认分辨率计，完成后按 DashScope 返回的分辨率结算。
- “是否生成音频”可以给 `wan2.6-i2v-flash` 的有声、无声视频分别设置价格。默认按有声计，请求把 `audio` 设为 `false` 时按无声计，完成后按 DashScope 的返回结算。

图片模型的计价字段如下。

- “图片生成单价”按图片张数计费。网关先按请求的 `n` 预扣（图文混排输出按 `max_images`），完成后按 DashScope 实际返回的张数结算。
- “是否开启提示词改写”可以给开启了 `prompt_extend` 的 `z-image-turbo` 请求单独定价。
- “输出图片档位”可以给 `qwen-image-3.0` 和 `qwen-image-3.0-pro` 的 1K、2K 输出分别设置价格。网关按 `size` 估算档位，没填 `size` 时按 2K 预扣，完成后按 DashScope 的返回结算。
- “输入图片单价”按输入图片张数计费，适用于 `qwen-image-3.0` 和 `qwen-image-3.0-pro`。网关先按请求里的图片数计，完成后按 DashScope 返回的张数结算。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "wan2.6-t2v", "prompt": "A tram climbing a hillside street in the rain", "seconds": 5, "resolution": "1080P", "ratio": "16:9"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 每个万相视频模型支持的时长和分辨率各不相同，超出范围的请求在发往 DashScope 前就会被拒绝。
- 时长设为 `-1` 时由万相 3.0 自己决定视频长度，其他模型会拒绝这个值。
- `/v1/videos` 的图片必须是 URL，上传文件会被拒绝。
- OpenAI Images API 在全部图片生成后一次返回，不支持 `stream`。上传的单张图片不超过 10 MiB。
- 千问图像和 Z-Image 模型走同步的 `multimodal-generation` 路径。`text2image` 和 `image2image` 上较早的万相图片模型以任务方式运行。
