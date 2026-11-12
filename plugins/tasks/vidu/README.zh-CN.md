---
locale: zh-CN
---
# Vidu

使用生数科技的 Vidu 模型生成视频，支持文生视频、图生视频、首尾帧生视频和参考图生视频。

## 渠道设置

- 渠道类型选 Vidu (52)。
- 密钥填 Vidu 开放平台的 API Key。
- API 地址留空时使用 `https://api.vidu.cn`。只有 Vidu 账号使用其他 API 地址时才需要修改，填写时去掉末尾的 `/ent/v2`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装 Vidu 插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /ent/v2/text2video` | 文生视频 | `viduq3-pro`、`viduq3-turbo`、`viduq2`、`viduq1`、`vidu1.5` |
| `POST /ent/v2/img2video` | 图生视频 | `viduq4-preview`、`viduq3-pro`、`viduq3-turbo`、`viduq3-pro-fast`、`viduq2-pro`、`viduq2-turbo`、`viduq1-classic`、`viduq1`、`vidu2.0`、`vidu1.5` |
| `POST /ent/v2/start-end2video` | 首尾帧生视频 | `viduq3-pro`、`viduq3-turbo`、`viduq3-pro-fast`、`viduq2-pro`、`viduq2-turbo`、`viduq2-pro-fast`、`viduq1`、`viduq1-classic`、`vidu2.0`、`vidu1.5` |
| `POST /ent/v2/reference2video` | 参考图生视频 | `viduq4-preview`、`viduq3`、`viduq3-turbo`、`viduq3-mix`、`viduq3-ad`、`viduq3-drama`、`viduq2-pro`、`viduq2`、`viduq1`、`vidu2.0`、`vidu1.5` |
| `POST /v1/videos` | OpenAI 视频接口 | 以上全部模型 |
| `POST /v1/responses` | Responses API，支持流式、同步等待和后台三种调用方式 | 以上全部模型 |

Vidu 原生路径接收 Vidu 自己的请求体，按 Vidu 的格式返回。这些路径加上 `/vidu` 前缀也能调用，例如 `/vidu/ent/v2/text2video`。客户端可以像调用 Vidu 一样用 `Authorization: Token sk-...` 认证，也可以用 `Bearer`。

在 `/v1/videos` 和 `/v1/responses` 上，图片数量决定生成方式：没有图片是文生视频，一张是图生视频，两张是首尾帧，三张及以上是参考图生视频。其他 Vidu 字段放在 `metadata` 里。

原生路径的任务用 `GET /ent/v2/tasks/{task_id}/creations` 轮询，视频链接在 `creations` 里。
`/v1/videos` 的任务用 `GET /v1/videos/{task_id}` 轮询，完成后从 `GET /v1/videos/{task_id}/content` 下载。Responses 的结果里带有网关提供的视频链接。

## 计费

- “Vidu 积分单价”按 Vidu 积分计费。任务完成时 Vidu 会报告消耗的积分，网关在结算时计入。
- “视频生成单价”按视频秒数计费。秒数取自请求中的 `duration`（OpenAI 视频接口用 `seconds`）。
- “输出视频分辨率”可以给 `360p`、`540p`、`720p`、`1080p`、`2K` 和 `4K` 分别定价。分辨率取自请求中的 `resolution`。在 `/v1/videos` 和 `/v1/responses` 上也可以传 `size`（例如 `1920x1080`），网关会换算成分辨率。在 `/v1/responses` 上，`resolution` 要放在 `metadata` 里。

请求未填写时长或分辨率时，按该模型在对应接口上的默认值计费。在 Vidu 原生路径上，如果插件没有为该接口列出这个模型，或者无法识别请求里的值，就按 Vidu 为该模型提供的最长时长和最高分辨率计费。

## 示例

```bash
curl https://your-gateway.example.com/ent/v2/text2video \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "viduq3-pro", "prompt": "A red fox running through fresh snow", "duration": 5, "resolution": "1080p"}'
```

用返回的 `task_id` 轮询任务：

```bash
curl https://your-gateway.example.com/ent/v2/tasks/{task_id}/creations \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 带 `callback_url` 的请求会被拒绝，请改为轮询任务。
- 在 `/v1/videos` 上，上传的 `input_reference` 文件最大 15 MiB。
