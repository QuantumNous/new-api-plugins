---
translations:
  zh-CN: README.zh-CN.md
---
# Google Veo (Gemini API)

Generates videos from text or a first-frame image with Google Veo on the Gemini API.

## Channel setup

- Choose the Gemini (24) channel type.
- The key is a Gemini API key from Google AI Studio.
- Leave the API address empty to use `https://generativelanguage.googleapis.com`.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video from a prompt. For image to video, upload the image as the `input_reference` file in a multipart request, or send it in `images` in JSON. | All models below |
| `POST /v1/responses` | Creates a video through the Responses API, as a stream, a blocking call or a background call. | All models below |
| `POST /v1beta/models/{model}:predictLongRunning` | Creates a video with the Gemini API's own request body. | All models below |
| `GET /v1beta/models/{model}/operations/{task_id}` | Reads the operation that `predictLongRunning` returned. | All models below |

The models are `veo-3.0-generate-001`, `veo-3.0-fast-generate-001`, `veo-3.1-generate-preview` and `veo-3.1-fast-generate-preview`.

Clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains a gateway link to the video. Background responses are read with `GET /v1/responses/{response_id}`.
The Gemini API paths also work under `/google`. They take the gateway token as `Authorization: Bearer sk-...`, `x-goog-api-key` or `?key=`.

## Billing

- "Video generation unit price" is charged per second of video. The gateway reads the duration from the request and uses 8 seconds when it is missing.
- "Output video resolution" gives `720p`, `1080p` and `4k` their own prices. The gateway reads `metadata.resolution`, or derives it from `size`, and uses `720p` otherwise.

The charge is based on the request.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "veo-3.1-fast-generate-preview", "prompt": "Waves rolling onto a black sand beach at sunset", "seconds": 8, "size": "1920x1080"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Veo takes 4, 6 or 8 seconds and the resolutions `720p`, `1080p` and `4k`. Each request returns one video.
- Image to video takes one image, as a data URL or base64. An uploaded file is at most 20 MiB.
- A `size` taller than it is wide gives a 9:16 video. Other sizes give 16:9.
- A Gemini API request holds one instance with `prompt` and an optional `image`.
