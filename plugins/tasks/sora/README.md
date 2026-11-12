---
translations:
  zh-CN: README.zh-CN.md
---
# Sora

Generates videos with OpenAI's Sora 2 models from a text prompt or a reference image, and remixes finished videos.

## Channel setup

- Choose the Sora (55) channel type. An OpenAI (1) channel works as well.
- The key is an OpenAI API key from the OpenAI platform.
- Leave the API address empty to use `https://api.openai.com`. Enter another address only when you reach OpenAI through a different host, and leave `/v1` off the end.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs the Sora plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video from a prompt. For image to video, send `input_reference` in JSON or upload it as a file in a multipart request. | `sora-2`, `sora-2-pro` |
| `POST /v1/videos/{video_id}/remix` | Remixes a finished video with a new prompt. | The model of the original video |
| `POST /v1/responses` | Creates a video through the Responses API, as a stream, a blocking call or a background call. | `sora-2`, `sora-2-pro` |

Clients poll `GET /v1/videos/{task_id}` for the status and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains a gateway link to the video. Clients read a background response with `GET /v1/responses/{response_id}`.

## Billing

- "Video generation unit price" is charged per second of video. The gateway reads `seconds` (or `duration`) from the request and uses 4 seconds when it is missing.
- "Output video dimensions" gives `720x1280`, `1280x720`, `1792x1024` and `1024x1792` their own prices. The gateway reads `size` from the request and uses `720x1280` when it is missing.

When the video is ready, the gateway settles the charge on the seconds and size in Sora's result. A remix takes both values from that result.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "sora-2", "prompt": "A paper boat drifting down a rainy city street", "seconds": "8", "size": "1280x720"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- A multipart request takes one `input_reference` file.
- On `/v1/responses`, the plugin forwards the prompt, `seconds`, `size`, `metadata` and one image. When the request has several images, the first one is used.
