---
translations:
  zh-CN: README.zh-CN.md
---
# TypeSafe

Answers typed questions about text, data or images with TypeSafe System One and the Jev models. The answers come back in the same response.

## Channel setup

- Choose the Task Plugin (61) channel type and bind the TypeSafe plugin.
- The key is a TypeSafe API key. See the [TypeSafe API docs](https://docs.typesafe.ai/api).
- Leave the API address empty to use `https://api.typesafe.ai`. An address ending in `/v1` also works.
- To use another New API gateway as the upstream, choose New API (60) and enter that gateway's address and one of its tokens. That gateway needs the TypeSafe plugin too.

## Models and endpoints

| Endpoint | What it does | Models |
| --- | --- | --- |
| `POST /v1/systemone` | Sends a System One request and returns TypeSafe's answers. | `jev-1.14.0`, `jev-1.13.0`, `jev-latest`, `jev-preview` |
| `POST /typesafe/v1/systemone` | Same as above, under the plugin's own prefix. | Same as above |

The official TypeSafe SDKs work with the gateway root as their base URL.
Questions use the types `choice`, `score`, `noul` and `box`.
Images go in `assets`, as base64 in JSON or as file parts that the SDKs upload. The plugin adds TypeSafe's image header to these requests.
Send `TypeSafe-Danger-Zone-Long-Context: true` to use TypeSafe's long context.
Errors from TypeSafe reach the client in TypeSafe's format.

## Billing

Each model has one price field, "Input token unit price".
When the request arrives, the gateway reserves quota for 65,792 input tokens, or 262,400 with the long context header.
When TypeSafe answers, the gateway settles the charge on the input tokens TypeSafe reports. Image tokens count as input tokens.

## Example

```bash
curl https://your-gateway.example.com/v1/systemone \
  -H "Authorization: Bearer sk-..." \
  -H "Content-Type: application/json" \
  -d '{
    "model": "jev-1.14.0",
    "state": {"ticket": "The package arrived crushed and the mug inside is broken."},
    "questions": {
      "request": {"type": "choice", "criteria": {"refund": "Wants the money back", "replacement": "Wants a new mug", "other": null}},
      "severity": {"type": "score", "criteria": ["Minor", "Moderate", "Severe"]}
    }
  }'
```

## Limits

- The request takes the top-level fields `model`, `state`, `questions` and `assets`. A request with any other top-level field is rejected.
- A `choice` question has 1 to 255 options. A `score` question has 1 to 10 levels.
