---
locale: zh-CN
---
# Jimeng

使用火山引擎即梦，根据文字描述、单张图片或首尾帧生成视频。

## 渠道设置

- 渠道类型选 Jimeng (51)。
- 密钥填火山引擎控制台创建的访问密钥，格式为 `AccessKeyID|SecretAccessKey`。
- API 地址留空时使用 `https://visual.volcengineapi.com`。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装本插件。
- Jimeng 渠道的密钥以 `sk-` 开头时，插件会把它当作另一台 New API 网关的令牌，请求发往渠道的 API 地址。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/videos` | 生成视频。图生视频时用 `input_reference` 传一张图，首尾帧时在 `images` 里传两张图。multipart 请求可以上传一个 `input_reference` 文件。 | `jimeng_vgfm_t2v_l20` |
| `POST /v1/responses` | 通过 Responses API 根据文字生成视频，支持流式、同步等待和后台三种调用方式。 | `jimeng_vgfm_t2v_l20` |
| `POST /api/common/v3/submit_task_v2` | 用即梦 `CVSync2AsyncSubmitTask` 的请求体提交任务。 | 请求体里的 `req_key` |
| `POST /api/common/v3/get_result_v2` | 用即梦 `CVSync2AsyncGetResult` 的请求体查询任务。 | 请求体里的 `req_key` |

`jimeng_vgfm_t2v_l20` 是即梦 S2.0 Pro，请求带图片时按图生视频处理。
要用即梦 3.0，把 `metadata.req_key` 设为 `jimeng_v30`（720p）、`jimeng_v30_1080p` 或 `jimeng_v30_pro`，也可以在渠道的模型映射里把模型映射成这几个名字。插件按图片数量选用文生视频、首帧或首尾帧对应的 `req_key`。
即梦原生路径把 `req_key` 当作模型名，客户端会用到的每个 `req_key` 都要加进渠道的模型列表。

客户端轮询 `GET /v1/videos/{task_id}` 查看任务状态，完成后从 `GET /v1/videos/{task_id}/content` 下载视频。
Responses 的结果里带有网关提供的视频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。
即梦原生路径也可以加 `/jimeng` 前缀访问。

## 计费

- “视频生成单价”按视频秒数计费。S2.0 Pro 的视频按 5 秒计。即梦 3.0 的视频在请求 241 帧或 10 秒时按 10 秒计，其他情况按 5 秒计。
- “产品档位”可以给 S2.0 Pro、3.0 720p、3.0 1080p 和 3.0 Pro 分别设置价格。档位由发给即梦的 `req_key` 决定。

费用按请求内容计算。不属于这些档位的 `req_key` 在发往即梦前就会被拒绝。

## 示例

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "jimeng_vgfm_t2v_l20", "prompt": "A red kite rising over a wheat field", "seconds": 5, "size": "1280x720"}'
```

用返回的 `id` 轮询任务，完成后下载视频：

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## 限制

- S2.0 Pro 生成 5 秒的视频。即梦 3.0 生成 5 秒或 10 秒的视频，`metadata.frames` 设为 241 是 10 秒，设为 121 是 5 秒。
- 画面比例支持 `16:9`、`9:16`、`1:1`、`4:3`、`3:4` 和 `21:9`，`size` 会换算成最接近的比例。
- `/v1/responses` 只接受文字输入。
- 上传的 `input_reference` 文件不超过 4.7 MiB。
- 在即梦原生路径上查询结果时不能带 `req_json`，因为网关返回的是它最近一次轮询到的任务结果。
