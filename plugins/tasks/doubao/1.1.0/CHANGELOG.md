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

- Add Seedream 5.0 Pro, 5.0 Lite, 5.0, 4.5, and 4.0 image generation, including text prompts and reference-image editing, through `/doubao/api/v3/images/generations` and OpenAI Responses.
- Support model-specific resolution presets and pixel sizes, reference-image limits, prompt optimization, watermark, seed, and output format. Preserve explicit zero and false values and reject unsupported parameter combinations before submission.
- Support sequential image groups on compatible models with `max_images` from 1 to 15, defaulting to 15 when group generation is enabled. Reference images and generated images share the 15-image limit.
- Support Seedream 5.0 Pro layer decomposition from one input image, with one base image and up to 16 layers. Layer sizes accept `1K`, `1.5K`, `2K`, or `auto`; transparent backgrounds are available on the native route with the supported input and format combination.
- Return complete URL-based image results from native calls and image links from Responses in stream, sync, and background modes. Responses retains image artifacts for later retrieval; native image creation returns its result once and does not support upstream event streaming or Base64 image output.
- Support an upstream New API gateway for image and video requests through a type-60 channel with this plugin enabled on both gateways.

### Changed

- Report generated images separately for output at or below 2.61 megapixels and output above that boundary, plus reference-image count and the layer-decomposition condition. Successful layer outputs are counted individually at their own size tier.
- Seedance requests reject resolution tiers unsupported by the selected model. Seedance 1.5 Pro reports the audio-generation condition, defaulting to enabled; reference-video usage is reported for Seedance 2.x models.

### Migration

- **Image pricing:** Configure prices for the newly enabled Seedream models. Expressions should price `u("images_up_to_1_5k")`, `u("images_above_1_5k")`, and `u("input_images")`, with a separate rate when `u("layer_decomposition")` is true if required. For Seedream 5.0 Pro reference-image charging from the second input, use `max(u("input_images") - 1, 0)`. Per-call pricing charges by output-image count and does not distinguish these conditions. Existing prices are not rewritten automatically.
- **Video pricing:** Review existing Seedance expressions before upgrading. Only Seedance 2.x reports `video_input`; remove that field from expressions for 1.0 and 1.5 models. Use `u("generate_audio")` for separate audio/silent prices on 1.5 Pro, and restrict resolution conditions to each model's supported tiers. Saved expressions are not migrated automatically.
- **Installation and gateway connections:** Upgrade to a new-api build supporting the [current plugin contract](../../../../docs/plugin-api/v1.md) before installing 1.1.0. For a type-60 connection, enable this plugin on both gateways and configure the upstream gateway URL and token.
- **Image requests:** Save native image results from the create response. Use Responses when later artifact retrieval is needed. The native route accepts only URL output and a non-streaming request; the native transparency string `background` is distinct from the Responses boolean background-execution flag.
