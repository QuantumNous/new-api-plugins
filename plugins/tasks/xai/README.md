---
translations:
  zh-CN: README.zh-CN.md
---
# xAI

Generates, edits and extends videos with xAI's Grok Imagine models.

## Channel setup

- Choose the xAI (48) channel type.
- The key is an xAI API key from the xAI console.
- Leave the API address empty to use `https://api.x.ai`. Enter another address only when you reach xAI through a different host, and leave `/v1` off the end.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs the xAI plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos/generations` | Text to video, image to video and reference to video, with xAI's own request body | `grok-imagine-video`, `grok-imagine-video-1.5`, `grok-imagine-video-1.5-lite` |
| `POST /v1/videos/edits` | Edits a video with a prompt | `grok-imagine-video` |
| `POST /v1/videos/extensions` | Extends a video | `grok-imagine-video` |
| `POST /v1/videos` | OpenAI video API | All three models |
| `POST /v1/responses` | Responses API, as a stream, a blocking call or a background call | All three models |

The xAI paths return xAI's replies. They also work under the `/xai` prefix, such as `/xai/v1/videos/generations`.

Clients poll a task from an xAI path at `GET /v1/videos/{request_id}`. The reply keeps xAI's format, answers 202 while the video renders and carries the video URL when it is done.
For a task from `POST /v1/videos`, the same path returns the OpenAI video object, and clients download the file from `GET /v1/videos/{task_id}/content`. A finished Responses result contains a gateway link to the video.

## Billing

The price fields depend on the model. `grok-imagine-video` has five, and the 1.5 models have all but the input video price.

- "Video generation unit price" is charged per second of output video. A generation is charged for the requested `duration`, 8 seconds by default, and settled on the length xAI reports. For an edit the gateway reserves 9 seconds and settles on the output length. An extension is charged for the added length, 6 seconds by default.
- "Output video resolution" gives each resolution its own price. A generation uses the requested `resolution`, `480p` by default. For edits and extensions the gateway reserves `720p` and settles on the resolution of the finished file.
- "Input image unit price" is charged per image in the request. The gateway counts `image`, `reference_images`, `last_frame` and the images in `keyframes`.
- "Input video unit price" is charged per second of input video, on `grok-imagine-video` only. The gateway reserves 9 seconds for an edit and 15 for an extension, then settles on the input length worked out from xAI's result.
- "Blocked by content moderation" is true when xAI blocked the video and still billed it. The gateway charges these tasks, and this field lets them have their own price.

## Example

```bash
curl https://your-gateway.example.com/v1/videos/generations \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "grok-imagine-video-1.5", "prompt": "A hot air balloon rising over a misty valley at dawn", "duration": 8, "resolution": "720p", "aspect_ratio": "16:9"}'
```

Poll the task with the `request_id` from the reply:

```bash
curl https://your-gateway.example.com/v1/videos/{request_id} \
  -H "Authorization: Bearer sk-..."
```

## Limits

- `grok-imagine-video` renders `480p` and `720p`. The 1.5 models also render `1080p`. Reference to video renders `480p` or `720p`.
- A generation lasts 1 to 15 seconds. An extension adds 2 to 10 seconds.
- Images and videos are passed as URLs or base64 data URLs. Requests with `file_id` or `storage_options` are rejected.
