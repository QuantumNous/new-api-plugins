---
translations:
  zh-CN: README.zh-CN.md
---
# Content moderation

Checks what users send with an OpenAI-compatible moderation API, such as OpenAI's `/v1/moderations`, and rejects requests the API flags. It can also check images, documents and the model's replies.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Content moderation**. In its **Settings**, enter the moderation API address and its API key.
3. Under **Where it runs**, choose **All requests**, **Chosen groups** or **Channels that add it**. For the last, add the instance in each channel's settings under **Request processing plugins**. Under **Models**, you can limit it to some models.
4. Save. **Try it** in the editor runs the plugin on a sample request with the current settings, and **Dry run** on the **Request processing** page runs the saved instances. Both send real requests to the moderation API.

## Settings

- **Moderation API URL**: an endpoint that accepts OpenAI moderation requests. The default is `https://api.openai.com/v1/moderations`.
- **API key**: sent as a Bearer token. Leave it empty when the API needs no key.
- **Moderation model**: the default is `omni-moderation-latest`. Leave it empty to use the API's default model.
- **Text to check**: **Last user message** (the default) or **All user messages**.
- **Check images**: on by default. Images in the checked messages go to the moderation API too.
- **Blocked categories**: leave empty to reject every flagged request. Pick categories, such as **Harassment** or **Violence**, to reject only requests flagged for them.
- **Moderate replies**: off by default. Turn it on to check the model's replies as well.

## What is checked

- The user messages of Chat Completions, Responses, Claude Messages and Gemini requests. System prompts, assistant messages and tool results are not checked.
- The prompt of completion and image generation requests, and the input of embedding requests.
- The prompts of task and Midjourney submissions, such as video or music generation. Negative prompts are not checked.
- Images that the checked messages carry, or link to with an `http` or `https` URL: up to three per request. An image sent inside the request is skipped when it is over 20 MiB.
- Text, PDF and DOCX files that the checked messages carry: up to four per request, the first 10,000 characters of each.
- Requests to the gateway's own `/v1/moderations` are never checked. You can point the plugin at this gateway with one of its tokens.

## What clients see

- A flagged request is rejected with HTTP 400, error code `request_rejected` and the message "The request was flagged by content moderation (code: moderation.flagged)". It is not sent to any channel.
- Text too long to check in four moderation calls is rejected with code `moderation.too_long`.
- If the moderation API cannot be reached, or answers with an error, the request fails with HTTP 500 and error code `request_plugin_failed`. The request is not sent.
- With **Moderate replies** on, a flagged reply is stopped. A non-stream reply becomes HTTP 400 with error code `response_rejected`. A stream ends with an error event. A reply the moderation API fails to check is stopped too, with error code `request_plugin_failed`.
- The log of a stopped request shows the categories that blocked it.

## Limits

- Each check makes up to four calls to the moderation API and waits up to 30 seconds for them in all. Text goes first; images use the calls left, one image per call.
- With **Moderate replies** on, a streamed reply reaches the client in bursts of about 1,000 characters, each released after the moderation API passes it.
- Reply checks cover the text of Chat Completions, Responses, Claude Messages and Gemini replies. Task results, images and audio are not checked.
- Requires new-api v1.0.0-rc.50 or later.
