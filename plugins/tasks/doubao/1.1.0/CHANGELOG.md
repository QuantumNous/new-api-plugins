---
changelogVersion: 1
plugin: "doubao"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Added

- Add Seedream 5.0 Pro, 5.0 Lite, 5.0, 4.5, and 4.0 for text-to-image generation and reference-image editing. Call them through `/doubao/api/v3/images/generations` or OpenAI Responses.
- Customize image size, prompt optimization, watermarks, seed, and output format where supported by the selected model.
- Generate a series of images in one request on compatible models. Reference images and generated images share a limit of 15; when enabled, group generation defaults to a maximum of 15 outputs.
- Split an image into a base image and up to 16 layers with Seedream 5.0 Pro. Choose 1K, 1.5K, 2K, or automatic sizing, with transparent backgrounds available through the native API for supported input and output formats.
- Use OpenAI Responses for synchronous, streaming, or background image generation, and retrieve the generated images later. The native image API returns image links when generation finishes and does not offer streaming or Base64 output.
- Use another New API server as the upstream for Seedream image generation and Seedance video generation.

### Changed

- Seedream supports separate image prices for 1.5K and below (up to 2.61 megapixels) and larger sizes, plus reference-image charges. Layer decomposition can use separate rates, with each generated layer counted at its own size.
- Seedance now rejects resolutions that the selected model does not support.
- Seedance 1.5 Pro supports separate prices for videos with and without audio. Audio generation is enabled by default.
- Reference-video pricing now applies only to Seedance 2.x models.

### Migration

- Configure prices for the Seedream models you enable. To charge by image size or for reference images, use billing expressions: `u("images_up_to_1_5k")` and `u("images_above_1_5k")` count outputs in each size group, while `u("input_images")` counts reference images. Per-call pricing counts only output images.
- For Seedream 5.0 Pro, use `max(u("input_images") - 1, 0)` to charge for reference images starting with the second. Use `u("layer_decomposition")` to select a separate layer-decomposition rate when needed.
- Remove `video_input` from Seedance 1.0 and 1.5 billing expressions that use it. To set separate audio and silent-video prices for 1.5 Pro, use `u("generate_audio")`.
- This release requires the [updated plugin API](../../../../docs/plugin-api/v1.md). Upgrade new-api to a build that supports it before installing the plugin.
- To connect to an upstream New API server, enable this plugin on both servers and configure a type-60 channel with the upstream address and API key.
- Save native image results from the creation response. Use OpenAI Responses if you need to retrieve images later.
