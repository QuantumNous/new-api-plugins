---
translations:
  zh-CN: README.zh-CN.md
---
# Alibaba Bailian

Generates videos with Wan, and images with Wan, Qwen-Image and Z-Image, on Alibaba Cloud Bailian.

## Channel setup

- Choose the Ali (17) channel type.
- The key is a DashScope API key from the Bailian console.
- Leave the API address empty to use `https://dashscope.aliyuncs.com`.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video. Images are URLs in `image`, `images` or `input_reference`. | Wan video models |
| `POST /v1/images/generations` | Creates images and returns them when they are ready. | Image models |
| `POST /v1/images/edits` | Edits images. Input images are URLs, data URLs or multipart uploads. | Image models |
| `POST /v1/responses` | Creates a video or images through the Responses API, as a stream, a blocking call or a background call. | All models below |
| `POST /api/v1/services/aigc/video-generation/video-synthesis`, `POST /api/v1/services/aigc/image2video/video-synthesis` | Creates a video with DashScope's own body. | Wan video models |
| `POST /api/v1/services/aigc/multimodal-generation/generation` | Creates images with DashScope's own body and answers when they are ready. | `wan2.7-image-pro`, `wan2.7-image`, `wan2.6-image`, `wan2.6-t2i`, Qwen-Image and Z-Image models |
| `POST /api/v1/services/aigc/image-generation/generation`, `POST /api/v1/services/aigc/text2image/image-synthesis`, `POST /api/v1/services/aigc/image2image/image-synthesis` | Creates or edits images as a task with DashScope's own body. | Wan image models |
| `GET /api/v1/tasks/{task_id}` | Reads a DashScope task. | Wan video and image models |

The Wan video models are `wan3.0-video`, `wan3.0-video-prime`, `wan2.7-t2v`, `wan2.7-i2v`, `wan2.6-t2v`, `wan2.6-t2v-us`, `wan2.6-i2v`, `wan2.6-i2v-flash`, `wan2.6-i2v-us`, `wan2.5-t2v-preview`, `wan2.5-i2v-preview`, `wan2.2-t2v-plus`, `wan2.2-i2v-flash`, `wan2.2-i2v-plus`, `wan2.2-kf2v-flash`, `wan2.2-s2v`, `wanx2.1-t2v-plus`, `wanx2.1-t2v-turbo`, `wanx2.1-i2v-plus` and `wanx2.1-i2v-turbo`.
The Wan image models are `wan2.7-image-pro`, `wan2.7-image`, `wan2.6-image`, `wan2.6-t2i`, `wan2.5-t2i-preview`, `wan2.2-t2i-flash`, `wan2.2-t2i-plus`, `wanx2.1-t2i-turbo`, `wanx2.1-t2i-plus`, `wanx2.0-t2i-turbo`, `wan2.5-i2i-preview` and `wanx2.1-imageedit`.
The Qwen-Image models are `qwen-image`, `qwen-image-plus`, `qwen-image-max`, `qwen-image-2.0`, `qwen-image-2.0-pro`, `qwen-image-3.0`, `qwen-image-3.0-pro`, `qwen-image-edit`, `qwen-image-edit-plus` and `qwen-image-edit-max`. The Z-Image model is `z-image-turbo`.
Dated snapshots such as `wan2.7-t2v-2026-04-25` and `qwen-image-2.0-pro-2026-03-03` behave like their base model.

Video clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains gateway links to the video or images. Background responses are read with `GET /v1/responses/{response_id}`.
DashScope's own paths also work under `/ali`.

## Billing

Wan video models have these price fields.

- "Video generation unit price" is charged per second of video. The gateway reads the duration from the request and uses 5 seconds when it is missing. `wan2.2-s2v` holds 20 seconds, because its length follows the audio. Wan 3.0 holds 30 seconds for a duration of `-1` or a reference video, and bills input plus output video seconds. When the video is ready, the charge is settled on the seconds DashScope reports.
- "Output video resolution" gives each resolution the model offers its own price. The gateway reads `resolution` or `size`, uses the model's default when both are missing, and settles on the resolution DashScope reports.
- "Whether audio is generated" gives `wan2.6-i2v-flash` videos with and without audio their own prices. Audio counts as on unless the request sets `audio` to `false`. The charge is settled on DashScope's report.

Image models have these price fields.

- "Image generation unit price" is charged per image. The gateway holds the requested `n`, or `max_images` for interleaved output, and settles on the number of images DashScope returns.
- "Whether prompt rewriting is enabled" gives `z-image-turbo` requests that turn on `prompt_extend` their own price.
- "Output image tier" gives 1K and 2K output their own prices on `qwen-image-3.0` and `qwen-image-3.0-pro`. The gateway estimates the tier from `size`, holds 2K when `size` is missing, and settles on DashScope's report.
- "Input image unit price" is charged per input image on `qwen-image-3.0` and `qwen-image-3.0-pro`. The gateway counts the request's images and settles on the count DashScope reports.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "wan2.6-t2v", "prompt": "A tram climbing a hillside street in the rain", "seconds": 5, "resolution": "1080P", "ratio": "16:9"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Each Wan video model has its own durations and resolutions. Requests outside them are refused before they reach DashScope.
- A duration of `-1` lets Wan 3.0 choose the length. Other models refuse it.
- `/v1/videos` takes images as URLs. File uploads are refused.
- The OpenAI Images API answers with all images at once and refuses `stream`. An uploaded image is at most 10 MiB.
- Qwen-Image and Z-Image models use the synchronous `multimodal-generation` path. The earlier Wan image models on `text2image` and `image2image` run as tasks.
