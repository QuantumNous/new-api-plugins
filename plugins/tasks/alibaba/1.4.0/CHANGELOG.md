---
changelogVersion: 1
plugin: "alibaba"
version: "1.4.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.4.0]

### Added

- Add Qwen-Image, Plus, Max, 2.0/2.0-Pro, 3.0/3.0-Pro, Edit, Edit-Plus, Edit-Max, their declared dated snapshots, and Z-Image-Turbo to native image generation and OpenAI Responses. These models use synchronous upstream calls; output and reference-image limits follow the selected model.
- Add OpenAI-compatible image generation at `/v1/images/generations` and editing at `/v1/images/edits` for supported image models. Generation accepts JSON; editing accepts JSON or multipart image uploads. Requests default to one output and return the completed image list, including when the upstream uses an asynchronous task.
- Support image URLs, Base64 data URLs, multipart reference images, and `url` or `b64_json` output on the OpenAI image endpoints. Each uploaded input image is limited to 10 MiB; unsupported streaming requests are rejected.
- Add `wan2.5-i2i-preview` with one to three reference images and `wanx2.1-imageedit` with a single base image at the native `/ali/api/v1/services/aigc/image2image/image-synthesis` route, Responses, and the OpenAI image endpoints.
- Support the ten `wanx2.1-imageedit` operations, including instruction editing, masked editing, stylization, expansion, super-resolution, colorization, and watermark removal. OpenAI edits default to instruction editing, or masked editing when a mask is supplied; operation-specific parameters are forwarded.
- Accept `prompt_extend_mode` (`direct` or `agent`) and `enable_thinking` for image requests, including vendor `input` and `parameters` objects passed through the OpenAI image endpoints.
- Support an upstream New API gateway through a type-60 channel with this plugin enabled on both gateways. Image and video requests use the plugin native routes; synchronous interleaved output is returned as complete JSON.

### Changed

- Report Qwen-Image-3.0 output image count and 1K/2K output tier, plus input image count, using the upstream usage values when valid. Report Z-Image prompt rewriting as a billing condition; per-call pricing applies a factor of two when rewriting is enabled.
- Report whether `wan2.6-i2v-flash` generates audio, defaulting to enabled when omitted and using the completion audio flag when provided.
- Synchronous native image results are returned only in the create response and are no longer available through later result or artifact queries.

### Migration

- **Price configuration:** Configure prices for newly enabled image models. Qwen-Image-3.0 expressions should price `u("image_count")` by `u("output_image_type")` and account for `u("input_image_count")`; Z-Image expressions should account for `u("prompt_extend")`. Existing expressions are not automatically rewritten. Z-Image per-call pricing doubles the image charge when prompt rewriting is enabled, so verify the configured base price before enabling it.
- **Video pricing:** To charge different rates for audio and silent `wan2.6-i2v-flash` output, update its expression to branch on `u("audio")`. Omitting the field keeps a single rate. Check saved resolution conditions against each enabled model's supported tiers; saved prices are not changed automatically.
- **Installation and gateway connections:** Upgrade to a new-api build supporting the [current plugin contract](../../../../docs/plugin-api/v1.md) before installing 1.4.0. For a type-60 connection, enable this plugin on both gateways and configure the upstream gateway URL and token.
- **Image requests:** Save synchronous native results from the create response. Specify `n` explicitly when relying on a particular image count: OpenAI image requests default to one, while existing native asynchronous model defaults still apply. Use `function` and a mask only with compatible editing operations.
