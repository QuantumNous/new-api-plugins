---
changelogVersion: 1
plugin: "doubao"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Generate and edit Seedream images through the OpenAI Images API at `/v1/images/generations` and `/v1/images/edits`.
- Upload reference images directly with multipart edit requests, including multiple files. Each uploaded image is limited to 30 MiB, subject to the gateway's download limit. JSON requests accept reference-image URLs or Base64 data URLs.
- Choose image links or Base64 output with `response_format: "url"` or `"b64_json"`. The response returns when image generation finishes; streaming and mask-based editing are not supported.
- Use Seedream group generation and layer decomposition through the OpenAI Images API on models that support them. `n` accepts only `1`; use `sequential_image_generation` to request multiple images.
- OpenAI Images calls report output-image counts by size tier, reference-image counts, and whether layer decomposition is enabled for usage-based pricing.

### Migration

- Before installing, upgrade new-api if your build does not support the [OpenAI Images plugin protocol](../../../../docs/plugin-api/v1.md#host-protocols).
