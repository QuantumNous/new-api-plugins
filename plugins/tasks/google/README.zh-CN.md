---
locale: zh-CN
---
# Google Veo (Gemini API)

通过 Gemini API 调用 Google Veo，根据文字描述或首帧图生成视频。

## 渠道设置

- 渠道类型选 Gemini (24)。
- 密钥填 Google AI Studio 创建的 Gemini API Key。
- API 地址留空时使用 `https://generativelanguage.googleapis.com`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 根据提示词生成视频。图生视频时，在 multipart 请求里上传 `input_reference` 文件，或在 JSON 的 `images` 里传图片。 | 下列全部模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频，支持流式、同步等待和后台三种调用方式。 | 下列全部模型 |
| `POST /v1beta/models/{model}:predictLongRunning` | 用 Gemini API 原生请求体生成视频。 | 下列全部模型 |
| `GET /v1beta/models/{model}/operations/{task_id}` | 查询 `predictLongRunning` 返回的操作。 | 下列全部模型 |

支持的模型有 `veo-3.0-generate-001`、`veo-3.0-fast-generate-001`、`veo-3.1-generate-preview` 和 `veo-3.1-fast-generate-preview`。

客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
Gemini API 路径也可以加 `/google` 前缀访问。网关令牌可以放在 `Authorization: Bearer sk-...`、`x-goog-api-key` 或 `?key=` 里。

## 计费

- “视频生成单价”按视频秒数计费。秒数取自请求，未填写时按 8 秒计。
- “输出视频分辨率”可以给 `720p`、`1080p` 和 `4k` 分别设置价格。分辨率取自 `metadata.resolution`，没有时按 `size` 换算，都没有时按 `720p` 计。

费用按请求内容计算。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "veo-3.1-fast-generate-preview", "prompt": "Waves rolling onto a black sand beach at sunset", "seconds": 8, "size": "1920x1080"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- Veo 支持 4、6、8 秒，分辨率支持 `720p`、`1080p` 和 `4k`。每次请求生成一个视频。
- 图生视频使用一张图，格式为 data URL 或 base64。上传的文件不超过 20 MiB。
- `size` 高大于宽时生成 9:16 的视频，其他尺寸生成 16:9。
- Gemini API 请求包含一个 instance，字段为 `prompt` 和可选的 `image`。
