---
locale: zh-CN
---
# Hailuo Video

使用 MiniMax 海螺根据文字描述或图片生成视频，`MiniMax-H3` 还可以参考图片、视频和音频生成视频。

## 渠道设置

- 渠道类型选 MiniMax (35)。
- 密钥填 MiniMax 开放平台创建的 API Key。
- API 地址留空时使用 `https://api.minimax.chat`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 根据提示词生成视频。图生视频时传 `input_reference`，或在 multipart 请求里上传 `input_reference` 文件。 | 下列全部模型 |
| `POST /v1/responses` | 通过 Responses API 生成视频，支持流式、同步等待和后台三种调用方式。第二张图作为尾帧。 | 下列全部模型 |
| `POST /v1/video_generation` | 用 MiniMax 原生请求体生成视频。 | 海螺 2.3、02 和 01 系列 |
| `GET /v1/query/video_generation?task_id={task_id}` | 查询任务。 | 海螺 2.3、02 和 01 系列 |
| `GET /v1/files/retrieve?file_id={task_id}` | 返回已完成视频的下载链接，文件 ID 就是任务 ID。 | 海螺 2.3、02 和 01 系列 |
| `POST /v2/video_generation` | 用 MiniMax 原生的 H3 请求体生成视频。 | `MiniMax-H3` |
| `GET /v2/query/video_generation/{task_id}` | 查询 H3 任务。 | `MiniMax-H3` |

支持的模型有 `MiniMax-H3`、`MiniMax-Hailuo-2.3`、`MiniMax-Hailuo-2.3-Fast`、`MiniMax-Hailuo-02`，以及 01 系列的 `T2V-01-Director`、`T2V-01`、`I2V-01-Director`、`I2V-01-live`、`I2V-01` 和 `S2V-01`。

客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
MiniMax 原生路径也可以加 `/hailuo` 前缀访问。

## 计费

每个模型系列有各自的计价字段。

- “视频生成单价”按视频秒数计费，所有模型都有。海螺 2.3、02 和 01 系列的秒数取自请求，未填写时按 6 秒计。`MiniMax-H3` 未填写时按 5 秒计，走 MiniMax 原生路径时按 15 秒预扣，完成后按 MiniMax 返回的秒数结算。
- “输出视频分辨率”可以给每档分辨率分别设置价格：`MiniMax-H3` 有 `768P` 和 `2K`，海螺 2.3 和 2.3 Fast 有 `768P` 和 `1080P`，海螺 02 有 `512P`、`768P` 和 `1080P`。分辨率取自请求，未填写时按 `768P` 计（MiniMax 原生 H3 路径按 `2K` 预扣），完成后按成片的分辨率结算。01 系列不按分辨率计价。
- “输入图片单价”按输入图片张数计费，仅 `MiniMax-H3` 有。网关先按请求里的图片数计，完成后按 MiniMax 返回的张数结算。
- “输入视频单价”按输入视频秒数计费，仅 `MiniMax-H3` 有。请求带参考视频时先按 15 秒预扣，完成后按 MiniMax 返回的秒数结算。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "MiniMax-Hailuo-2.3", "prompt": "A fox trotting through fresh snow in a pine forest", "seconds": 6, "size": "1920x1080"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 海螺 2.3、2.3 Fast 和 02 生成 6 秒或 10 秒的视频。10 秒视频为 `768P`，海螺 02 图生视频时也可以是 `512P`。01 系列生成 6 秒的视频。
- `MiniMax-H3` 生成 4 到 15 秒的视频，最多接受 9 张图片、3 段参考视频和 3 段参考音频。一个请求要么用首尾帧图片，要么用参考素材。
- `MiniMax-Hailuo-2.3-Fast` 只做图生视频。
- 上传的 `input_reference` 文件不超过 20 MiB。
- 带 `callback_url` 的请求会被拒绝，请改为轮询任务。
