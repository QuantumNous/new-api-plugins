---
translations:
  zh-CN: README.zh-CN.md
---
# Google Veo & Gemini Omni (Vertex AI)

Generates videos from text or images with Google Veo and Gemini Omni Flash on Google Cloud Vertex AI.

## Channel setup

- Choose the Vertex AI (41) channel type and set the Vertex AI key format to JSON.
- The key is a service account key in JSON form, created in the Google Cloud console for a project with Vertex AI enabled. Requests run in the key's `project_id`.
- The deployment region sets where requests run, for example `us-central1`. It can also be JSON that maps model names to regions, with `default` for the rest; the lookup uses the model name the client sent. This needs plugin 1.3.1 or later on a new-api release after v1.0.0-rc.49. Earlier versions run every request in `global`.
- Leave the API address empty to use the Vertex AI address of the region, `https://aiplatform.googleapis.com` for `global`. Enter another address only when you reach Vertex AI through a different host, and leave `/v1` off the end.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs this plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/videos` | Creates a video from a prompt. For image to video, upload the image as the `input_reference` file in a multipart request, or send it in `images` in JSON. | All models below |
| `POST /v1/responses` | Creates a video through the Responses API, as a stream, a blocking call or a background call. | All models below |
| `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning` | Creates a video with Vertex AI's own request body. | Veo models |
| `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:fetchPredictOperation` | Reads the operation that `predictLongRunning` returned. | Veo models |

The Veo models are `veo-3.0-generate-001`, `veo-3.0-fast-generate-001`, `veo-3.1-generate-preview` and `veo-3.1-fast-generate-preview`.
The Gemini Omni models are `gemini-omni-1.2-flash`, `gemini-omni-1.1-flash-preview` and `gemini-omni-flash-preview`.

Clients poll `GET /v1/videos/{task_id}` and download the file from `GET /v1/videos/{task_id}/content`.
A Responses result is a text message saying the video is ready. Use `/v1/videos` when the client needs the file.
The Vertex AI paths also work with `/v1beta1/` in place of `/v1/` and under the `/vertex-ai` prefix. They take the gateway token as `Authorization: Bearer sk-...` and run in the channel's project, whatever project and location the path names.

## Billing

Veo models have three price fields. The gateway takes all three values from the request.

- "Video generation unit price" is charged per second of video. A request without a duration counts as 8 seconds.
- "Output video resolution" gives `720p`, `1080p` and `4k` their own prices. The gateway reads `resolution`, or derives it from `size`, and uses `720p` otherwise.
- "Whether audio is generated" lets videos with and without audio have different prices. Audio counts as on unless the request sets `generateAudio` to `false`.

Gemini Omni models are priced per token, in three fields.

- "Input token unit price" and "Text output token unit price" use the counts Vertex AI reports when the video is ready. Text output includes thinking.
- "Video output token unit price" is estimated at submission from the duration and resolution, then settled on the video and audio output tokens Vertex AI reports. The estimate counts 1,931 tokens per second at 360p, 5,792 at 720p, 8,688 at 1080p and 17,376 at 4k.

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

- Requests fail on a channel whose key format is API Key.
- Veo takes 4, 6 or 8 seconds, the resolutions `720p`, `1080p` and `4k`, and one input image. Each request generates one video.
- Gemini Omni takes whole seconds from 3 to 15 and the resolutions `360p`, `720p`, `1080p` and `4k`. On `/v1/responses` it takes several images: two set the first and last frames, and more act as subject references.
- Input images are JPEG or PNG, sent as base64 or data URLs. An uploaded file is at most 20 MiB.
- A `size` taller than it is wide gives a 9:16 video. Other sizes give 16:9.
- A Vertex AI request holds one instance with `prompt` and an optional `image`.
- Gemini Omni is unavailable on a New API (60) channel.
