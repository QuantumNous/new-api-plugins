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

- Add Qwen-Image text-to-image and editing models, including Plus, Max, 2.0/Pro, 3.0/Pro, and the Edit series. They are available through the native image API and OpenAI Responses.
- Add Z-Image-Turbo text-to-image generation, with optional prompt rewriting.
- Generate and edit images through the OpenAI-compatible `/v1/images/generations` and `/v1/images/edits` endpoints. Both default to one image and wait for the finished images before returning.
- Upload reference images directly for editing, or provide image URLs or Base64 data URLs. OpenAI image requests support URL or Base64 results; uploaded images can be up to 10 MiB each.
- Add `wan2.5-i2i-preview` image editing with one to three reference images. It works through the native image API, OpenAI Responses, and the OpenAI image endpoints.
- Add `wanx2.1-imageedit` for editing a single image, with ten operations including masked editing, stylization, expansion, upscaling, colorization, and watermark removal. OpenAI edits use instruction editing by default and switch to masked editing when a mask is supplied.
- Allow image requests to select direct or agent-assisted prompt rewriting and enable thinking on models that support these options.
- Use another New API server as the upstream for image and video generation, including synchronous responses containing both text and images.

### Changed

- Qwen-Image-3.0 supports separate prices for 1K and 2K output images and for input images.
- Z-Image can be priced differently when prompt rewriting is enabled. With per-call pricing, enabling rewriting doubles the image charge.
- `wan2.6-i2v-flash` supports different prices for videos with and without audio. Audio generation is enabled by default.
- Synchronous native image results are now returned only when the image is created. They cannot be retrieved with a later query.

### Migration

- Configure prices for the image models you enable. For Qwen-Image-3.0, use `u("image_count")` for output count, `u("output_image_type")` to select the 1K or 2K rate, and `u("input_image_count")` for input count.
- For Z-Image, use `u("prompt_extend")` in billing expressions to select the prompt-rewriting rate. With per-call pricing, check the base price against the doubled charge before enabling rewriting.
- To set separate audio and silent-video prices for `wan2.6-i2v-flash`, use `u("audio")` in its billing expression.
- This release requires the [updated plugin API](../../../../docs/plugin-api/v1.md). Upgrade new-api to a build that supports it before installing the plugin.
- To connect to an upstream New API server, enable this plugin on both servers and configure a type-60 channel with the upstream address and API key.
- Save synchronous native image results from the creation response. When switching to the OpenAI image endpoints, set `n` if you need more than one image.
