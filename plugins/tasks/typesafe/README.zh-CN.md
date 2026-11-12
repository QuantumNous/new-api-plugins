---
locale: zh-CN
---
# TypeSafe

用 TypeSafe System One 和 Jev 模型，就一段文字、数据或图片回答带类型的问题。答案在同一个响应里返回。

## 渠道设置

- 渠道类型选任务插件 (61)，并绑定 TypeSafe 插件。
- 密钥填 TypeSafe 的 API Key，详见 [TypeSafe API 文档](https://docs.typesafe.ai/api)。
- API 地址留空时使用 `https://api.typesafe.ai`。地址末尾带 `/v1` 也可以。
- 如果上游是另一台 New API 网关，渠道类型选 New API (60)，填写该网关的地址和它的令牌。上游网关也需要安装 TypeSafe 插件。

## 模型与接口

| 接口 | 用途 | 模型 |
| --- | --- | --- |
| `POST /v1/systemone` | 发送 System One 请求，返回 TypeSafe 的答案。 | `jev-1.14.0`、`jev-1.13.0`、`jev-latest`、`jev-preview` |
| `POST /typesafe/v1/systemone` | 同上，使用插件自己的路径前缀。 | 同上 |

TypeSafe 官方 SDK 把网关根地址设为 base URL 即可使用。
问题类型有 `choice`、`score`、`noul` 和 `box`。
图片放在 `assets` 里，可以在 JSON 中用 base64，也可以由 SDK 以文件形式上传。插件会为这类请求加上 TypeSafe 的图片请求头。
需要 TypeSafe 的长上下文时，在请求头里加 `TypeSafe-Danger-Zone-Long-Context: true`。
TypeSafe 返回的错误按 TypeSafe 的格式交给客户端。

## 计费

每个模型只有一个价格字段：“输入 token 单价”。
请求到达时，网关先按 65,792 个输入 token 预扣额度，带长上下文请求头时按 262,400 个。
TypeSafe 返回结果后，按它报告的输入 token 数结算。图片 token 计入输入 token。

## 示例

```bash
curl https://your-gateway.example.com/v1/systemone \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{
    "model": "jev-1.14.0",
    "state": {"ticket": "The package arrived crushed and the mug inside is broken."},
    "questions": {
      "request": {"type": "choice", "criteria": {"refund": "Wants the money back", "replacement": "Wants a new mug", "other": null}},
      "severity": {"type": "score", "criteria": ["Minor", "Moderate", "Severe"]}
    }
  }'
```

## 限制

- 请求的顶层字段限于 `model`、`state`、`questions` 和 `assets`。带其他顶层字段的请求会被拒绝。
- `choice` 问题有 1 到 255 个选项，`score` 问题有 1 到 10 个等级。
