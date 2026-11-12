---
translations:
  zh-CN: README.zh-CN.md
---
# Incho

Generates songs with vocals from a text prompt or lyrics with Incho's music model.

## Channel setup

- Choose the Task Plugin (61) channel type and bind the Incho plugin.
- The key is an API key from the [Incho platform](https://platform.yinchaoyongxian.com).
- Leave the API address empty to use `https://open.yinchaoyongxian.com`.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /incho/submit/music` | Generates songs from Incho's request fields: `prompt`, `lyric`, `n` and `model`. | `incho_music` |
| `GET /incho/fetch/{task_id}` | Reads a task. | `incho_music` |
| `POST /v1/responses` | Generates songs through the Responses API, as a stream, a blocking call or a background call. `lyric`, `n` and `model` go in `metadata`. | `incho_music` |

`incho_music` uses Incho v4.0. Send `model` as `v3.5` to use Incho v3.5.

Clients poll `GET /incho/fetch/{task_id}`. A finished task lists each song with its title, lyrics and audio link.
A finished Responses result contains the lyrics and gateway links to the audio. Background responses are read with `GET /v1/responses/{response_id}`.

## Billing

"Song generation unit price" is charged per song. The gateway holds `n` songs when the task starts, 2 when `n` is missing. When the task ends, the charge is settled on the songs that finished with audio.

## Example

```bash
curl https://your-gateway.example.com/incho/submit/music \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"prompt": "An upbeat city pop song about a summer night drive", "n": 1}'
```

The reply carries the task ID in `data`. Poll it until the songs are ready:

```bash
curl https://your-gateway.example.com/incho/fetch/{task_id} \
  -H "Authorization: Bearer sk-..."
```

## Limits

- A request needs a `prompt` or `lyric`. `prompt` is at most 1,000 characters and `lyric` at most 3,000.
- `n` is 1 or 2.
- The plugin uses Incho's prompt-to-song mode. Reference, extend and lyrics-only generation are unavailable.
- Incho is unavailable on a New API (60) channel.
