---
locale: zh-CN
---
# 全局参数覆盖

在请求发往上游前，对所有请求或指定分组的请求应用参数覆盖规则。规则语法与渠道的参数覆盖相同，一套规则就能覆盖多个渠道。只修改单个渠道时，请使用该渠道自己的参数覆盖。

## 添加方法

1. 在控制台打开“请求处理”，点击“添加”。
2. 选择“全局参数覆盖”，在“设置”里以 JSON 填写规则。
3. 在“用在哪里”选择“所有请求”或“指定分组”。“模型”可以限定只对部分模型生效。
4. 保存。编辑时可以用“试一试”按当前设置在示例请求上运行规则；“请求处理”页面的“试运行”会运行已保存的实例。

## 设置

- **参数覆盖**：规则内容，写法与渠道的参数覆盖设置相同，语法见[参数覆盖说明](https://docs.newapi.ai/en/docs/guide/console/channel-management#parameter-override-settings-documentation)。默认为空，不做任何修改。

规则可以设置、删除或改写请求体字段和请求头。条件可以判断模型、上游模型、分组、渠道、请求路径和请求格式。

## 示例

下面的规则对名称以 `gpt-` 开头的模型把 `temperature` 设为 0.7，删除 `metadata`，并添加 `X-Request-Source` 请求头：

```json
{
  "operations": [
    {
      "path": "temperature",
      "mode": "set",
      "value": 0.7,
      "conditions": [{"path": "model", "mode": "prefix", "value": "gpt-"}]
    },
    {"path": "metadata", "mode": "delete"},
    {"path": "X-Request-Source", "mode": "set_header", "value": "new-api"}
  ]
}
```

## 对客户端的影响

- `return_error` 规则会拒绝请求，无论规则里写的是什么状态码，都返回 HTTP 400 和错误码 `request_rejected`。客户端看到的是规则里的提示。请求不会再发往其他渠道。
- 规则试图修改计费字段（如 `model`、`stream`、`n`）或受保护的请求头（如 `Authorization`、`Content-Type`）时，请求失败，返回 HTTP 500 和错误码 `request_plugin_failed`。
- 使用日志会列出规则改动的字段和请求头。被渠道自己的参数覆盖再次改写的改动会标记为已被覆盖。

## 限制

- 渠道自己的参数覆盖和请求头覆盖在本插件之后运行。同一字段或请求头以渠道的设置为准。
- 规则作用于网关为渠道构建好的请求体，每次尝试渠道时都会运行。
- 任务提交、Midjourney 提交、音频请求，以及网关以表单形式发往上游的请求（例如图片编辑），插件不会运行。
- 网关或渠道开启请求体透传时，插件不会运行。
- 复制或透传客户端请求头的规则在这里不起作用，因为插件读不到客户端的请求头。
- 需要 new-api v1.0.0-rc.50 或更高版本。
