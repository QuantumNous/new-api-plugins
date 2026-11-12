---
locale: zh-CN
---
# Incho

使用音潮 Incho 的音乐模型，根据提示词或歌词生成带人声的歌曲。

## 渠道设置

- 渠道类型选任务插件 (61)，并绑定 Incho 插件。
- 密钥填[音潮开放平台](https://platform.yinchaoyongxian.com)创建的 API Key。
- API 地址留空时使用 `https://open.yinchaoyongxian.com`。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /incho/submit/music` | 用音潮的请求字段 `prompt`、`lyric`、`n` 和 `model` 生成歌曲。 | `incho_music` |
| `GET /incho/fetch/{task_id}` | 查询任务。 | `incho_music` |
| `POST /v1/responses` | 通过 Responses API 生成歌曲，支持流式、同步等待和后台三种调用方式。`lyric`、`n` 和 `model` 放在 `metadata` 里。 | `incho_music` |

`incho_music` 默认使用音潮 v4.0。要用 v3.5，把 `model` 设为 `v3.5`。

客户端轮询 `GET /incho/fetch/{task_id}`。任务完成后，结果里列出每首歌的标题、歌词和音频链接。
Responses 的结果里带有歌词和网关提供的音频链接。后台调用的结果用 `GET /v1/responses/{response_id}` 查询。

## 计费

“生成歌曲单价”按歌曲数量计费。任务开始时按 `n` 首预扣，未填写 `n` 时按 2 首计。任务结束后，按成功生成音频的歌曲数结算。

## 示例

```bash
curl https://your-gateway.example.com/incho/submit/music \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"prompt": "An upbeat city pop song about a summer night drive", "n": 1}'
```

返回的 `data` 就是任务 ID，用它轮询，直到歌曲生成完毕：

```bash
curl https://your-gateway.example.com/incho/fetch/{task_id} \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 请求至少要有 `prompt` 或 `lyric` 之一。`prompt` 不超过 1000 个字符，`lyric` 不超过 3000 个字符。
- `n` 只能是 1 或 2。
- 插件使用音潮的提示词生成歌曲模式，参考生成、续写和单独生成歌词都不可用。
- Incho 不能用于 New API (60) 渠道。
