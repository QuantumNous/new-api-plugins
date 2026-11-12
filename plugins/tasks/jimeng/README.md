---
translations:
  zh-CN: README.zh-CN.md
---
# Jimeng

Generates videos from text, one image, or a first and last frame with Volcengine Jimeng.

## Channel setup

- Choose the Jimeng (51) channel type.
- The key is a Volcengine access key pair from the Volcengine console, written as `AccessKeyID|SecretAccessKey`.
- Leave the API address empty to use `https://visual.volcengineapi.com`.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.
- On a Jimeng channel, a key that starts with `sk-` is read as a token for another New API gateway at the channel's API address.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video. Send one image as `input_reference` for image to video, or two in `images` for a first and last frame. A multipart request can upload one `input_reference` file. | `jimeng_vgfm_t2v_l20` |
| `POST /v1/responses` | Creates a video from text through the Responses API, as a stream, a blocking call or a background call. | `jimeng_vgfm_t2v_l20` |
| `POST /api/common/v3/submit_task_v2` | Submits a task with the body of Jimeng's `CVSync2AsyncSubmitTask` action. | The `req_key` in the body |
| `POST /api/common/v3/get_result_v2` | Reads a task with the body of Jimeng's `CVSync2AsyncGetResult` action. | The `req_key` in the body |

`jimeng_vgfm_t2v_l20` is Jimeng S2.0 Pro. A request with an image runs as image to video.
For Jimeng 3.0, set `metadata.req_key` to `jimeng_v30` (720p), `jimeng_v30_1080p` or `jimeng_v30_pro`, or map the model to one of these names in the channel. The plugin picks the text, first-frame or first-and-last-frame `req_key` from the number of images.
Jimeng's own paths use the `req_key` as the model name, so the channel's model list needs every `req_key` clients send.

Clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains a gateway link to the video. Background responses are read with `GET /v1/responses/{response_id}`.
Jimeng's own paths also work under `/jimeng`.

## Billing

- "Video generation unit price" is charged per second of video. S2.0 Pro videos count as 5 seconds. Jimeng 3.0 videos count as 10 seconds when the request asks for 241 frames or 10 seconds, and 5 seconds otherwise.
- "Product tier" gives S2.0 Pro, 3.0 720p, 3.0 1080p and 3.0 Pro their own prices. The tier follows the `req_key` sent to Jimeng.

The charge is based on the request. A `req_key` outside these tiers is refused before it reaches Jimeng.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "jimeng_vgfm_t2v_l20", "prompt": "A red kite rising over a wheat field", "seconds": 5, "size": "1280x720"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- S2.0 Pro makes 5-second videos. Jimeng 3.0 makes 5 or 10 seconds. Set `metadata.frames` to 241 for 10 seconds or 121 for 5.
- The aspect ratios are `16:9`, `9:16`, `1:1`, `4:3`, `3:4` and `21:9`. A `size` maps to the nearest one.
- `/v1/responses` takes text only.
- An uploaded `input_reference` file is at most 4.7 MiB.
- Result queries on Jimeng's own path refuse `req_json`, because the gateway answers from its last poll of the task.
