---
translations:
  zh-CN: README.zh-CN.md
---
# Vidu

Generates videos with Shengshu's Vidu models from text, an image, first and last frames, or reference images.

## Channel setup

- Choose the Vidu (52) channel type.
- The key is a Vidu API key from the Vidu platform.
- Leave the API address empty to use `https://api.vidu.cn`. Change it only if your Vidu account uses a different API address, and leave `/ent/v2` off the end.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs the Vidu plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /ent/v2/text2video` | Text to video | `viduq3-pro`, `viduq3-turbo`, `viduq2`, `viduq1`, `vidu1.5` |
| `POST /ent/v2/img2video` | Image to video | `viduq4-preview`, `viduq3-pro`, `viduq3-turbo`, `viduq3-pro-fast`, `viduq2-pro`, `viduq2-turbo`, `viduq1-classic`, `viduq1`, `vidu2.0`, `vidu1.5` |
| `POST /ent/v2/start-end2video` | Video from a first and a last frame | `viduq3-pro`, `viduq3-turbo`, `viduq3-pro-fast`, `viduq2-pro`, `viduq2-turbo`, `viduq2-pro-fast`, `viduq1`, `viduq1-classic`, `vidu2.0`, `vidu1.5` |
| `POST /ent/v2/reference2video` | Video from reference images | `viduq4-preview`, `viduq3`, `viduq3-turbo`, `viduq3-mix`, `viduq3-ad`, `viduq3-drama`, `viduq2-pro`, `viduq2`, `viduq1`, `vidu2.0`, `vidu1.5` |
| `POST /v1/videos` | OpenAI video API | All models above |
| `POST /v1/responses` | Responses API, as a stream, a blocking call or a background call | All models above |

The Vidu paths take Vidu's own request body and return Vidu's replies. They also work under the `/vidu` prefix, such as `/vidu/ent/v2/text2video`. Clients authenticate with `Authorization: Token sk-...`, as Vidu does, or with `Bearer`.

On `/v1/videos` and `/v1/responses`, the number of images picks the mode. No image means text to video, one means image to video, two mean first and last frames, and three or more mean reference images. Other Vidu fields go in `metadata`.

Clients poll a Vidu path task at `GET /ent/v2/tasks/{task_id}/creations`. The video link is in `creations`.
Clients poll a `/v1/videos` task at `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`. A finished Responses result contains a gateway link to the video.

## Billing

- "Vidu credit unit price" is charged per Vidu credit. Vidu reports the credits a task used when it finishes, and the gateway adds them at settlement.
- "Video generation unit price" is charged per second of video. The gateway reads `duration` (`seconds` on the OpenAI video API) from the request.
- "Output video resolution" gives `360p`, `540p`, `720p`, `1080p`, `2K` and `4K` their own prices. The gateway reads `resolution` from the request. On `/v1/videos` and `/v1/responses` it also accepts `size`, such as `1920x1080`, and turns it into a resolution. On `/v1/responses`, `resolution` goes in `metadata`.

A request without a duration or resolution is charged at the model's default for that endpoint. On the Vidu paths, a model the plugin does not list for the endpoint, or a value it cannot read, is charged at the longest duration and highest resolution Vidu offers for that model.

## Example

```bash
curl https://your-gateway.example.com/ent/v2/text2video \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "viduq3-pro", "prompt": "A red fox running through fresh snow", "duration": 5, "resolution": "1080p"}'
```

Poll the task with the `task_id` from the reply:

```bash
curl https://your-gateway.example.com/ent/v2/tasks/{task_id}/creations \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Requests with `callback_url` are rejected. Poll the task instead.
- On `/v1/videos`, an uploaded `input_reference` file is at most 15 MiB.
