---
locale: zh-CN
---
# SunoAPI

通过你自己部署的 [SunoAPI](https://github.com/Suno-API/Suno-API) 服务生成歌曲和歌词。

## 渠道设置

- 渠道类型选 SunoAPI (36)。
- 密钥填你的 SunoAPI 服务要求的访问密钥，插件会以 Bearer 令牌的形式发送。
- API 地址填 SunoAPI 服务的地址。这个渠道类型没有默认地址。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装 SunoAPI 插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /suno/submit/music` | 根据描述或你写好的歌词生成歌曲。 | `suno_music` |
| `POST /suno/submit/lyrics` | 根据提示词生成歌词。 | `suno_lyrics` |
| `POST /v1/responses` | 通过 Responses API 生成歌曲或歌词，支持流式、同步等待和后台三种调用方式。输入的文字会作为歌曲描述或歌词提示词。 | `suno_music`、`suno_lyrics` |

客户端轮询 `GET /suno/fetch/{task_id}` 查询单个任务，查询多个任务时发送 `POST /suno/fetch`，请求体为 `{"ids": [...]}`。
Responses 的结果里带有歌词，以及网关提供的每首歌的音频链接。

## 计费

- “生成歌曲或歌词单价”按生成的条数计费。
- “生成歌曲或歌词”可以给歌曲和歌词分别定价。

提交时，歌曲请求按 2 条预扣，歌词请求按 1 条预扣。任务完成后，按 SunoAPI 实际返回的条数结算。

## 示例

```bash
curl https://your-gateway.example.com/suno/submit/music \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"gpt_description_prompt": "A cheerful acoustic song about a cat cafe on a rainy afternoon"}'
```

返回的 `data` 就是任务 ID。轮询该任务，直到状态变为 `SUCCESS`：

```bash
curl https://your-gateway.example.com/suno/fetch/{task_id} \
  -H "Authorization: Bearer sk-..."
```

## 限制

- 歌曲请求不带 `mv` 时使用 `chirp-v3-0`。不带 `make_instrumental` 时生成带人声的歌曲。
- 歌词请求必须带 `prompt`。
- 在 `/v1/responses` 上，`metadata` 里的字段会作为请求字段传给 SunoAPI。
