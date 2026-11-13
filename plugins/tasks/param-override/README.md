---
translations:
  zh-CN: README.zh-CN.md
---
# Global parameter override

Applies parameter override rules to every request, or to the requests of chosen groups, before they are sent upstream. The rules use the same syntax as a channel's parameter override, so one rule set can cover many channels. To change a single channel, use that channel's own parameter override.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Global parameter override**. In its **Settings**, enter the rules as JSON.
3. Under **Where it runs**, choose **All requests** or **Chosen groups**. Under **Models**, you can limit it to some models.
4. Save. **Try it** in the editor runs the rules on a sample request with the current settings, and **Dry run** on the **Request processing** page runs the saved instances.

## Settings

- **Parameter override**: the rules, written as a channel's parameter override setting. The syntax is described in the [parameter override guide](https://docs.newapi.ai/en/docs/guide/console/channel-management#parameter-override-settings-documentation). Empty by default, which changes nothing.

Rules can set, delete or rewrite body fields and request headers. Conditions can test the model, the upstream model, the group, the channel, the request path and the request format.

## Example

These rules set `temperature` to 0.7 for models whose name starts with `gpt-`, remove `metadata`, and add an `X-Request-Source` header:

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

## What clients see

- A `return_error` rule rejects the request with HTTP 400 and error code `request_rejected`, whatever status the rule names. The client sees the rule's message. The request is not sent to another channel.
- A rule that tries to change a billing field, such as `model`, `stream` or `n`, or a protected header, such as `Authorization` or `Content-Type`, fails the request with HTTP 500 and error code `request_plugin_failed`.
- The usage log lists the fields and headers the rules changed. A change that the channel's own parameter override replaced is marked as overridden.

## Limits

- The channel's own parameter override and header override run after this plugin. On the same field or header, the channel's setting wins.
- It runs on the body after the gateway builds it for the channel, on each channel attempt.
- It does not run on task and Midjourney submissions, audio requests, or requests the gateway sends upstream as a form, such as image edits.
- It does not run when the gateway or the channel passes the request body through unchanged.
- Rules that copy or pass on client request headers have no effect here, because the plugin cannot read them.
- Requires new-api v1.0.0-rc.50 or later.
