---
locale: zh-CN
---
# Google Veo & Gemini Omni (Vertex AI)

在 Google Cloud Vertex AI 上调用 Google Veo 和 Gemini Omni Flash，根据文字或图片生成视频。

## 渠道设置

- 渠道类型选 Vertex AI (41)，Vertex AI 密钥格式选 JSON。
- 密钥填服务账号的 JSON 密钥，在 Google Cloud 控制台为已启用 Vertex AI 的项目创建。请求在密钥里的 `project_id` 项目下运行。
- API 地址留空时使用 `https://aiplatform.googleapis.com`。只有通过其他地址访问 Vertex AI 时才需要填写，填写时去掉末尾的 `/v1`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 根据提示词生成视频。图生视频时，在 multipart 请求里上传 `input_reference` 图片文件，或在 JSON 的 `images` 里传图片。 | 下列全部模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频，支持流式、同步等待和后台三种调用方式。 | 下列全部模型 |
| `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning` | 使用 Vertex AI 原生请求体生成视频。 | Veo 模型 |
| `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:fetchPredictOperation` | 查询 `predictLongRunning` 返回的操作。 | Veo 模型 |

Veo 模型有 `veo-3.0-generate-001`、`veo-3.0-fast-generate-001`、`veo-3.1-generate-preview` 和 `veo-3.1-fast-generate-preview`。
Gemini Omni 模型有 `gemini-omni-1.2-flash`、`gemini-omni-1.1-flash-preview` 和 `gemini-omni-flash-preview`。

客户端轮询 `GET /v1/videos/{task_id}`，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果是一条提示视频已生成的文字消息。客户端需要视频文件时，请改用 `/v1/videos`。
Vertex AI 原生路径也可以把 `/v1/` 换成 `/v1beta1/`，或加上 `/vertex-ai` 前缀。这些路径用 `Authorization: Bearer sk-...` 传网关令牌。无论路径里写的是哪个项目和区域，请求都在渠道密钥所属的项目下运行。

## 计费

Veo 模型有三个价格字段，取值都来自请求。

- “视频生成单价”按视频秒数计费。请求未填写时长时按 8 秒计。
- “输出视频分辨率”可以给 `720p`、`1080p` 和 `4k` 分别定价。网关读取 `resolution`，没有时根据 `size` 推算，仍无法确定时按 `720p` 计。
- “是否生成音频”可以让有声和无声视频使用不同价格。除非请求把 `generateAudio` 设为 `false`，否则按生成音频计。

Gemini Omni 模型按 token 计费，有三个价格字段。

- “输入 Token 单价”和“文本输出 Token 单价”按视频完成时 Vertex AI 报告的数量计费。文本输出包含思考内容。
- “视频输出 Token 单价”在提交时根据时长和分辨率预估，完成后按 Vertex AI 报告的视频和音频输出 token 结算。预估时每秒按以下数量计：360p 为 1,931，720p 为 5,792，1080p 为 8,688，4k 为 17,376。

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

- 密钥格式为 API Key 的渠道无法使用本插件，请求会失败。
- 请求都在 `global` 区域运行，渠道的部署区域对本插件无效。
- Veo 支持 4、6、8 秒，分辨率 `720p`、`1080p` 和 `4k`，可以带一张输入图片。每次请求生成一个视频。
- Gemini Omni 支持 3 到 15 秒的整数时长，分辨率 `360p`、`720p`、`1080p` 和 `4k`。在 `/v1/responses` 上可以带多张图片：两张时作为首帧和尾帧，更多时作为主体参考图。
- 输入图片须为 JPEG 或 PNG，用 base64 或 data URL 传入。上传的文件最大 20 MiB。
- `size` 的高大于宽时生成 9:16 视频，其他尺寸生成 16:9 视频。
- Vertex AI 原生请求只能包含一个 instance，其中只有 `prompt` 和可选的 `image`。
- New API (60) 渠道无法使用 Gemini Omni。
