---
translations:
  zh-CN: README.zh-CN.md
---
# Kling

Generates videos from text or images with Kuaishou Kling, including Kling 4.0.

## Channel setup

- Choose the Kling (50) channel type.
- The key is the access key and secret key from the Kling developer platform, written as `AccessKey|SecretKey`. A single API key issued by Kling works as well.
- Leave the API address empty to use `https://api.klingai.com`.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.
- On a Kling channel, a key that starts with `sk-` is read as a token for another New API gateway at the channel's API address.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video from a prompt. For image to video, send `input_reference`, or upload it as a file in a multipart request. | All models below |
| `POST /v1/responses` | Creates a video through the Responses API, as a stream, a blocking call or a background call. A second image sets the last frame. | All models below |
| `POST /v1/videos/text2video`, `POST /v1/videos/image2video` | Creates a video with Kling's own body, which needs `model_name`. | `kling-v2-5-turbo`, `kling-v2-6`, `kling-v3` |
| `GET /v1/videos/text2video/{task_id}`, `GET /v1/videos/image2video/{task_id}` | Reads a task. | `kling-v2-5-turbo`, `kling-v2-6`, `kling-v3` |
| `POST /text-to-video/kling-4.0`, `POST /image-to-video/kling-4.0`, `POST /omni-video/kling-4.0` | Creates a Kling 4.0 video with Kling's own body. | `kling-4.0` |
| `GET /tasks?task_ids={task_id}` | Reads a Kling 4.0 task. | `kling-4.0` |

The models are `kling-v2-5-turbo`, `kling-v2-6`, `kling-v3` and `kling-4.0`.

Clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains a gateway link to the video. Background responses are read with `GET /v1/responses/{response_id}`.
Kling's own paths also work under `/kling`.

## Billing

"Kling credit unit price" is charged per Kling resource unit.

- When the task starts, the gateway holds an estimate. For `kling-v2-5-turbo`, `kling-v2-6` and `kling-v3` it uses the model, `mode`, duration, native audio (`sound: "on"`) and reference voices (`voice_list`). A missing duration counts as 5 seconds.
- For Kling 4.0 the estimate uses the duration, the resolution and any input video. A request without a duration holds 30 seconds.
- When the video is ready, the charge is settled on the units Kling reports deducting. A Kling 4.0 task without a reported deduction is settled on the video durations Kling reports.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "kling-v2-6", "prompt": "A lantern floating over a quiet lake at dusk", "seconds": 5, "mode": "pro", "size": "1280x720"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Durations are whole seconds. Kling 4.0 takes 3 to 30 seconds. Kling checks the durations of the other models.
- `mode` is `std` when the request leaves it out.
- `size` is written as WIDTHxHEIGHT. For text to video it maps to the nearest aspect ratio: 16:9, 9:16 or 1:1, and also 21:9 on Kling 4.0. On Kling 4.0 it also sets the resolution: 720p when the shorter side is under 900 pixels, and 1080p otherwise.
- An uploaded `input_reference` file is at most 10 MiB.
- Requests with `callback_url` or `external_task_id` are refused. Poll the task instead.
