---
translations:
  zh-CN: README.zh-CN.md
---
# Request filter

Rejects requests whose text contains a keyword or matches a regular expression that you list. It can also stop model replies that match.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Request filter**. In its **Settings**, add rules or a keyword list.
3. Under **Where it runs**, choose **All requests**, **Chosen groups** or **Channels that add it**. For the last, add the instance in each channel's settings under **Request processing plugins**. Under **Models**, you can limit it to some models.
4. Save. **Try it** in the editor runs the plugin on a sample request with the current settings, and **Dry run** on the **Request processing** page runs the saved instances.

## Settings

- **Rules**: a request that matches any rule is rejected. Up to 500 rules. Each rule has:
  - **Keyword or expression**: the text to look for.
  - **Match type**: **Keyword** (the default) finds the text anywhere. **Regular expression** uses JavaScript syntax. An expression that does not compile is refused when you save.
  - **Case sensitive**: off by default.
  - **Rule ID**: recorded in the logs when the rule matches. Leave it empty to number the rules in order: `rule-1`, `rule-2` and so on.
- **Keyword list**: one keyword per line. A request that contains any of them is rejected. Use it for long word lists.
- **Keywords are case sensitive**: off by default.
- **Rejection message**: the message the client sees. Leave it empty for the gateway's default message.
- **Check replies**: off by default. Turn it on to stop model replies that match a rule or the keyword list.

## What is checked

- All text of Chat Completions, Responses, Claude Messages and Gemini requests: system prompts, every message, tool calls and their results, and tool declarations.
- The prompt of completion and image generation requests, the input of embedding requests, and the query and documents of rerank requests.
- The prompts, negative prompts and lyrics of task and Midjourney submissions, such as video or music generation.
- Text, PDF and DOCX files sent inside the request.
- The final body sent upstream is checked again, so text a channel's parameter override adds is checked too.

## What clients see

- A matching request is rejected with HTTP 400 and error code `request_rejected`. The message is your rejection message, or the default "Request rejected by policy", followed by `(code: request_filtered)`. The request is not sent to another channel.
- With **Check replies** on, a matching reply is stopped. A non-stream reply becomes HTTP 400 with error code `response_rejected`. A stream ends with an error event.
- The matched text is never logged. Administrators see the ID of the rule that matched, and `keywords` for the keyword list.
- If the plugin fails on a request, the request fails with HTTP 500 and error code `request_plugin_failed`.

## Limits

- Images, and files the request names only by URL or file ID, are not checked.
- A streamed reply is held back a little, so that a match split across chunks is still found. With a regular expression rule, 256 characters are held back.
- Reply checks cover the text of Chat Completions, Responses, Claude Messages and Gemini replies. Task results, images and audio are not checked.
- Requires new-api v1.0.0-rc.50 or later.
