---
locale: zh-CN
---
# Kling

使用快手可灵（含可灵 4.0），根据文字描述或图片生成视频。

## 渠道设置

- 渠道类型选 Kling (50)。
- 密钥填可灵开发者平台的 AccessKey 和 SecretKey，格式为 `AccessKey|SecretKey`。也可以直接填可灵签发的单个 API Key。
- API 地址留空时使用 `https://api.klingai.com`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。
- Kling 渠道的密钥以 `sk-` 开头时，插件会把它当作另一台 New API 网关的令牌，请求发往渠道的 API 地址。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 根据提示词生成视频。图生视频时传 `input_reference`，或在 multipart 请求里上传 `input_reference` 文件。 | 下列全部模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频，支持流式、同步等待和后台三种调用方式。第二张图作为尾帧。 | 下列全部模型 |
| `POST /v1/videos/text2video`、`POST /v1/videos/image2video` | 用可灵原生请求体生成视频，必须带 `model_name`。 | `kling-v2-5-turbo`、`kling-v2-6`、`kling-v3` |
| `GET /v1/videos/text2video/{task_id}`、`GET /v1/videos/image2video/{task_id}` | 查询任务。 | `kling-v2-5-turbo`、`kling-v2-6`、`kling-v3` |
| `POST /text-to-video/kling-4.0`、`POST /image-to-video/kling-4.0`、`POST /omni-video/kling-4.0` | 用可灵原生请求体生成可灵 4.0 视频。 | `kling-4.0` |
| `GET /tasks?task_ids={task_id}` | 查询可灵 4.0 任务。 | `kling-4.0` |

支持的模型有 `kling-v2-5-turbo`、`kling-v2-6`、`kling-v3` 和 `kling-4.0`。

客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
可灵原生路径也可以加 `/kling` 前缀访问。

## 计费

“可灵资源包单位单价”按可灵资源包消耗的单位数计费。

- 任务开始时，网关先按预估量预扣。`kling-v2-5-turbo`、`kling-v2-6` 和 `kling-v3` 的预估依据模型、`mode`、时长、原生音频（`sound: "on"`）和参考音色（`voice_list`），未填时长时按 5 秒计。
- 可灵 4.0 的预估依据时长、分辨率和是否有输入视频，未填时长时按 30 秒预扣。
- 视频生成后，按可灵返回的实际扣减单位数结算。可灵 4.0 任务没有返回扣减记录时，按可灵返回的视频时长结算。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "kling-v2-6", "prompt": "A lantern floating over a quiet lake at dusk", "seconds": 5, "mode": "pro", "size": "1280x720"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 时长必须是整数秒。可灵 4.0 支持 3 到 30 秒，其他模型的时长由可灵校验。
- 请求未填 `mode` 时按 `std` 处理。
- `size` 写成“宽x高”的形式。文生视频时换算成最接近的画面比例：16:9、9:16 或 1:1，可灵 4.0 还支持 21:9。可灵 4.0 同时按 `size` 选择分辨率：短边小于 900 像素时为 720p，否则为 1080p。
- 上传的 `input_reference` 文件不超过 10 MiB。
- 带 `callback_url` 或 `external_task_id` 的请求会被拒绝，请改为轮询任务。
