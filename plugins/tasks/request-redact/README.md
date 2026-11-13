---
translations:
  zh-CN: README.zh-CN.md
---
# Request redaction

Replaces personal information and secret keys in request text before the gateway sends the request upstream, so the model never sees them. It can also mask the same information in model replies.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Request redaction** and fill in its **Settings**.
3. Under **Where it runs**, choose **All requests**, **Chosen groups** or **Channels that add it**. For the last, add the instance in each channel's settings under **Request processing plugins**. Under **Models**, you can limit it to some models.
4. Save. **Try it** in the editor shows a sample request before and after redaction, and **Dry run** on the **Request processing** page runs the saved instances.

## Settings

- **Built-in types**: the kinds of information to replace. All but IPv4 addresses are on by default.
  - **Email addresses**.
  - **Phone numbers**: mainland China mobile numbers, and international numbers that start with `+`.
  - **ID card numbers**: mainland China resident ID numbers with a valid check digit.
  - **Bank card numbers**: card numbers that pass the card check digit.
  - **API keys and tokens**: keys such as `sk-...`, AWS, Google, GitHub and Slack keys, and the token after `Bearer`.
  - **IPv4 addresses**: off by default, because version numbers and sample addresses look the same. Addresses that start with `127.` or `0.` are never replaced.
- **Custom rules**: up to 50 regular expressions, applied after the built-in types. Each has its own **Replacement text** and **Rule ID**. An expression that does not compile is refused when you save.
- **Default replacement text**: replaces each match, `[REDACTED]` by default.
- **Check before sending**: **Reject the request** (the default) or **Off**. See below.
- **Mask replies**: off by default. Turn it on to mask the same types and rules in model replies.

## What is redacted

- The text of messages, system prompts and instructions in Chat Completions, Responses, Claude Messages and Gemini requests.
- The prompt of completion and image generation requests, and the input of embedding requests.
- The prompts, negative prompts and lyrics of task and Midjourney submissions, such as video or music generation.
- Text files sent inside the request. PDF and DOCX files are read but cannot be changed.

## What clients see

- The model receives the redacted text, so its answer may contain the replacement text.
- With **Check before sending** on, the gateway rejects a request it cannot fully redact: a PDF or DOCX file that holds such information, or a final body that still holds it, for example after a channel's parameter override added it. The client gets HTTP 400 with error code `request_rejected` and the message "Request rejected by policy (code: sensitive_content)".
- With **Check before sending** off, those requests are sent as they are.
- With **Mask replies** on, each match in a reply is replaced by `***`. A streamed reply is held back by up to 128 characters, so that matches split across chunks are still found.
- The usage log shows which types were replaced and how many times. The original text is never logged.
- If the plugin fails on a request, the request fails with HTTP 500 and error code `request_plugin_failed`.

## Limits

- Tool call arguments, tool declarations and images are not redacted.
- Reply masking covers the text of Chat Completions, Responses, Claude Messages and Gemini replies. Task results, images and audio are not masked.
- Requires new-api v1.0.0-rc.50 or later.
