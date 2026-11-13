---
translations:
  zh-CN: README.zh-CN.md
---
# Doubao

Generates videos with Doubao Seedance and images with Doubao Seedream on Volcengine Ark.

## Channel setup

- Choose the Doubao (54) channel type. A VolcEngine (45) channel works as well.
- The key is an API key from the Volcengine Ark console.
- Leave the API address empty to use `https://ark.cn-beijing.volces.com`.
- The channel's model mapping can send a model to an Ark endpoint ID. Prices and options still follow the model the client asked for.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video. Send reference images as URLs in `images`, and other references in `metadata.content`. Ark fields such as `resolution`, `ratio` and `generate_audio` go in `metadata`. | Seedance models |
| `POST /v1/images/generations` | Creates images and returns them when they are ready. | Seedream models |
| `POST /v1/images/edits` | Edits images. Reference images are URLs, data URLs or multipart uploads. | Seedream models |
| `POST /v1/responses` | Creates a video or images through the Responses API, as a stream, a blocking call or a background call. | All models below |
| `POST /api/v3/contents/generations/tasks` | Creates a video with Ark's own body. | Seedance models |
| `GET /api/v3/contents/generations/tasks/{task_id}` | Reads a video task. | Seedance models |
| `POST /api/v3/images/generations` | Creates images with Ark's own body. With `stream: true` the images stream back as Ark sends them. | Seedream models |

The Seedance models are `doubao-seedance-1-0-pro-250528`, `doubao-seedance-1-0-lite-t2v`, `doubao-seedance-1-0-lite-i2v`, `doubao-seedance-1-5-pro-251215`, `doubao-seedance-2-0-260128`, `doubao-seedance-2-0-fast-260128`, `doubao-seedance-2-0-mini-260615` and `doubao-seedance-2-5-260628`.
The Seedream models are `doubao-seedream-5-5-261028`, `doubao-seedream-5-0-pro-260628`, `doubao-seedream-5-0-flash-260915`, `doubao-seedream-5-0-lite-260128`, `doubao-seedream-5-0-260128`, `doubao-seedream-4-5-251128` and `doubao-seedream-4-0-250828`.

Video clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains gateway links to the video or images. Background responses are read with `GET /v1/responses/{response_id}`.
Ark's own paths also work under `/doubao`.

## Billing

Seedance models have these price fields.

- "Billing token unit price" is charged per token. When the task starts, the gateway estimates tokens from the duration and resolution. A request with a duration of `-1` or none holds the model's longest video: 12 seconds for Seedance 1.0 and 1.5 pro, 15 for Seedance 2.0, and 30 for Seedance 2.5. When the video is ready, the charge is settled on the tokens Ark reports.
- "Output video resolution" gives each resolution the model offers its own price. The gateway reads `metadata.resolution` or `size`, holds `1080p` when both are missing (`720p` on Seedance 2.0 fast and mini), and settles on the resolution Ark reports.
- "Reference video input" gives requests with and without a reference video their own prices on Seedance 2.0 and 2.5.
- "Whether audio is generated" gives videos with and without audio their own prices on Seedance 1.5 pro. Audio counts as on unless the request sets `generate_audio` to `false`.

Seedream models have these price fields.

- "Image generation unit price (1.5K and below)" and "Image generation unit price (above 1.5K)" charge per output image, split at 2.61 megapixels. Seedream 5.0 lite, 5.0 and 4.5 make only larger images and have a single "Image generation unit price". The gateway estimates from the requested size and count, and settles on the size of each delivered image.
- "Input image unit price" is charged per reference image. The gateway counts the request's images and settles on the count Ark reports.
- "Whether layer decomposition is enabled" gives layer decomposition its own price on Seedream 5.5, 5.0 pro and flash. Each layer counts as an output image.

Group generation and layer decomposition hold the largest image count the request allows until the images arrive. Layer decomposition holds 17 images on Seedream 5.0 pro and flash, and on Seedream 5.5 the `layer_decomposition_options.max_images` of the request, or 30 without it.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "doubao-seedance-2-0-260128", "prompt": "A hot air balloon drifting over green hills at sunrise", "seconds": 5, "metadata": {"resolution": "1080p", "ratio": "16:9"}}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Seedance 1.0, 1.5 pro and 2.5 make `480p`, `720p` and `1080p`. Seedance 2.0 adds `4k`. Seedance 2.0 fast and mini make `480p` and `720p`. Other resolutions are refused.
- A duration of `-1` lets Seedance 1.5 pro, 2.0 and 2.5 choose the length.
- `/v1/videos` takes references as URLs. File uploads are refused.
- On the OpenAI Images API, `n` is 1. Use `sequential_image_generation` for several images. Streaming and `mask` are unavailable there.
- A transparent background (`background: "transparent"`) on Seedream 5.0 pro and flash needs exactly one reference image. Seedream 5.5 has no such rule and also takes `background: "auto"`. Transparent images cannot use `output_format: "jpeg"`.
- Seedream 5.5, 5.0 pro and flash take up to 10 reference images, and the other Seedream models up to 14. An uploaded image is at most 30 MiB.
- Requests with `callback_url` are refused. Poll the task instead.
