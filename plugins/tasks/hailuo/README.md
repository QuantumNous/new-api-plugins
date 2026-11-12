---
translations:
  zh-CN: README.zh-CN.md
---
# Hailuo Video

Generates videos with MiniMax Hailuo from text or images, and with `MiniMax-H3` from reference images, videos and audio.

## Channel setup

- Choose the MiniMax (35) channel type.
- The key is an API key from the MiniMax open platform.
- Leave the API address empty to use `https://api.minimax.chat`.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video from a prompt. For image to video, send `input_reference`, or upload it as a file in a multipart request. | All models below |
| `POST /v1/responses` | Creates a video through the Responses API, as a stream, a blocking call or a background call. A second image sets the last frame. | All models below |
| `POST /v1/video_generation` | Creates a video with MiniMax's own body. | Hailuo 2.3, 02 and the 01 series |
| `GET /v1/query/video_generation?task_id={task_id}` | Reads a task. | Hailuo 2.3, 02 and the 01 series |
| `GET /v1/files/retrieve?file_id={task_id}` | Returns the download link of a finished video. The file ID is the task ID. | Hailuo 2.3, 02 and the 01 series |
| `POST /v2/video_generation` | Creates a video with MiniMax's own H3 body. | `MiniMax-H3` |
| `GET /v2/query/video_generation/{task_id}` | Reads an H3 task. | `MiniMax-H3` |

The models are `MiniMax-H3`, `MiniMax-Hailuo-2.3`, `MiniMax-Hailuo-2.3-Fast`, `MiniMax-Hailuo-02`, and the 01 series: `T2V-01-Director`, `T2V-01`, `I2V-01-Director`, `I2V-01-live`, `I2V-01` and `S2V-01`.

Clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A finished Responses result contains a gateway link to the video. Background responses are read with `GET /v1/responses/{response_id}`.
MiniMax's own paths also work under `/hailuo`.

## Billing

Each model family has its own price fields.

- "Video generation unit price" is charged per second of video on every model. Hailuo 2.3, 02 and the 01 series take the duration from the request and count 6 seconds when it is missing. `MiniMax-H3` counts 5 seconds when it is missing, or 15 seconds on MiniMax's own path, and settles on the seconds MiniMax reports.
- "Output video resolution" gives each resolution its own price: `768P` and `2K` on `MiniMax-H3`, `768P` and `1080P` on Hailuo 2.3 and 2.3 Fast, and `512P`, `768P` and `1080P` on Hailuo 02. The gateway reads the resolution from the request, uses `768P` when it is missing (`2K` on MiniMax's own H3 path), and settles on the resolution of the finished video. The 01 series has no resolution price.
- "Input image unit price" is charged per input image on `MiniMax-H3`. The gateway counts the images in the request and settles on the count MiniMax reports.
- "Input video unit price" is charged per second of input video on `MiniMax-H3`. A request with a reference video holds 15 seconds, then settles on the seconds MiniMax reports.

## Example

```bash
curl https://your-gateway.example.com/v1/videos \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"model": "MiniMax-Hailuo-2.3", "prompt": "A fox trotting through fresh snow in a pine forest", "seconds": 6, "size": "1920x1080"}'
```

Poll the task with the `id` from the reply, then download the video:

```bash
curl https://your-gateway.example.com/v1/videos/{task_id} \
  -H "Authorization: Bearer sk-..."

curl -o video.mp4 https://your-gateway.example.com/v1/videos/{task_id}/content \
  -H "Authorization: Bearer sk-..."
```

## Limits

- Hailuo 2.3, 2.3 Fast and 02 make 6 or 10 seconds. A 10-second video is `768P`, or also `512P` for Hailuo 02 image to video. The 01 series makes 6-second videos.
- `MiniMax-H3` makes 4 to 15 seconds. It takes up to 9 images, 3 reference videos and 3 reference audio clips. A request uses either frame images or reference media.
- `MiniMax-Hailuo-2.3-Fast` makes image to video only.
- An uploaded `input_reference` file is at most 20 MiB.
- Requests with `callback_url` are refused. Poll the task instead.
