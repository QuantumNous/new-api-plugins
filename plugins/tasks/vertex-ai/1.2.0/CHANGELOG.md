---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Call Vertex AI's own paths on the gateway: `POST /v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning` to start a Veo video and `:fetchPredictOperation` on the same path to query it. The same paths under `/v1beta1` work too, and the gateway also serves each path under `/vertex-ai`, so existing `/vertex-ai/v1/...` addresses keep working.
- Add Gemini Omni Flash video generation (`gemini-omni-1.1-flash-preview` and `gemini-omni-flash-preview`) through OpenAI Responses and the OpenAI video API, from text or from a first frame. Videos run 3 to 10 seconds at 360p, 720p, 1080p or 4k. Tasks are billed by the tokens Vertex AI reports: input, text output including thinking, and video output. Gemini Omni is not available on a New API upstream channel.

### Changed

- Native requests reach Vertex AI as written, and replies keep its format, status codes and error codes. The gateway rejects a native request itself only when the model is missing or unknown, the request cannot be priced, or it uses a feature the gateway does not support.
- Once the gateway has failed a task, for example after a timeout, a native query shows the operation as failed even if Vertex AI's last reply still showed it running.

### Fixed

- Finished Veo videos download again. A video Vertex AI returns inline is fetched from Vertex AI again when it is downloaded, instead of being kept in the task, and a video written to Cloud Storage keeps its `gs://` address. Through another New API gateway, the video downloads from that gateway's task.

### Migration

- **Price configuration:** Set input, text output and video output token prices for `gemini-omni-1.1-flash-preview` and `gemini-omni-flash-preview` before enabling them. Vertex AI lists $1.50, $9 and $17.50 per million tokens. The video output price is reserved at submit time from the requested length and resolution.
- **Channel:** Gemini Omni needs a service account JSON key and the channel API address `https://aiplatform.googleapis.com`.
- **Installation:** Install 1.2.0 on a new-api release after v1.0.0-rc.47. Older versions reject the plugin or cannot serve its Vertex AI paths.
