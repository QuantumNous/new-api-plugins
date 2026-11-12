---
locale: zh-CN
---
# Sora

使用 OpenAI 的 Sora 2 模型，根据文字描述或参考图生成视频，也可以对已生成的视频做 remix。

## 渠道设置

- 渠道类型选 Sora (55)，选 OpenAI (1) 也可以。
- 密钥填 OpenAI 平台创建的 API Key。
- API 地址留空时使用 `https://api.openai.com`。只有通过其他地址访问 OpenAI 时才需要填写，填写时去掉末尾的 `/v1`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装 Sora 插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 根据提示词生成视频。图生视频时，在 JSON 里传 `input_reference`，或在 multipart 请求里上传 `input_reference` 文件。 | `sora-2`、`sora-2-pro` |
| `POST /v1/videos/{video_id}/remix` | 用新的提示词对已完成的视频做 remix。 | 与原视频相同 |
| `POST /v1/responses` | 通过 Responses API 生成视频，支持流式、同步等待和后台三种调用方式。 | `sora-2`、`sora-2-pro` |

客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。

## 计费

- “视频生成单价”按视频秒数计费。秒数取自请求中的 `seconds`（或 `duration`），未填写时按 4 秒计。
- “输出视频尺寸”可以给 `720x1280`、`1280x720`、`1792x1024` 和 `1024x1792` 分别定价。尺寸取自请求中的 `size`，未填写时按 `720x1280` 计。

视频生成完成后，网关按 Sora 返回结果中的秒数和尺寸结算。remix 的秒数和尺寸都以该结果为准。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "sora-2", "prompt": "A paper boat drifting down a rainy city street", "seconds": "8", "size": "1280x720"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- multipart 请求只能上传一个 `input_reference` 文件。
- 在 `/v1/responses` 上，插件转发提示词、`seconds`、`size`、`metadata` 和一张图片。请求带多张图片时使用第一张。
