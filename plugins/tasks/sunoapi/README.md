---
translations:
  zh-CN: README.zh-CN.md
---
# SunoAPI

Generates songs and lyrics through a [SunoAPI](https://github.com/Suno-API/Suno-API) service that you deploy yourself.

## Channel setup

- Choose the SunoAPI (36) channel type.
- The key is the access key your SunoAPI service expects. The plugin sends it as a Bearer token.
- Enter the address of your SunoAPI service as the API address. This channel type has no default address.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs the SunoAPI plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /suno/submit/music` | Generates songs from a description or from your own lyrics. | `suno_music` |
| `POST /suno/submit/lyrics` | Generates lyrics from a prompt. | `suno_lyrics` |
| `POST /v1/responses` | Generates songs or lyrics through the Responses API, as a stream, a blocking call or a background call. The input text becomes the song description or the lyrics prompt. | `suno_music`, `suno_lyrics` |

Clients poll `GET /suno/fetch/{task_id}` for one task, or send `POST /suno/fetch` with `{"ids": [...]}` for several.
A finished Responses result contains the lyrics and a gateway link to each song's audio.

## Billing

- "Song or lyrics generation unit price" is charged per generated clip.
- "Generate songs or lyrics" gives songs and lyrics their own prices.

The gateway reserves 2 clips for a song request and 1 clip for a lyrics request. When the task finishes, it settles on the number of clips SunoAPI returned.

## Example

```bash
curl https://your-gateway.example.com/suno/submit/music \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{"gpt_description_prompt": "A cheerful acoustic song about a cat cafe on a rainy afternoon"}'
```

The reply's `data` is the task ID. Poll it until the status is `SUCCESS`:

```bash
curl https://your-gateway.example.com/suno/fetch/{task_id} \
  -H "Authorization: Bearer sk-..."
```

## Limits

- A song request without `mv` uses `chirp-v3-0`. A request without `make_instrumental` generates vocals.
- A lyrics request needs a `prompt`.
- On `/v1/responses`, fields in `metadata` are passed to SunoAPI as request fields.
