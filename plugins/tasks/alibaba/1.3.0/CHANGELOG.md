---
changelogVersion: 1
plugin: "alibaba"
version: "1.3.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.3.0]

### Added

- Expand Alibaba Bailian from video tasks to image and video generation with ten image models. Add `wan2.7-image-pro` and `wan2.7-image` for text-to-image and reference-image editing, with up to nine reference images.
- Add `wan2.6-t2i` for text-to-image and `wan2.6-image` for image editing and interleaved text/image output. Ordinary editing accepts up to four reference images; interleaved generation accepts at most one.
- Add asynchronous text-to-image support for `wan2.5-t2i-preview`, `wan2.2-t2i-flash`, `wan2.2-t2i-plus`, `wanx2.1-t2i-turbo`, `wanx2.1-t2i-plus`, and `wanx2.0-t2i-turbo`.
- Add native asynchronous submission at `/ali/api/v1/services/aigc/image-generation/generation` for Wan 2.6/2.7 image models and `/ali/api/v1/services/aigc/text2image/image-synthesis` for the legacy image models. Requests use the vendor's `input` and `parameters` objects, with results available through `/ali/api/v1/tasks/:task_id`.
- Add native synchronous image submission at `/ali/api/v1/services/aigc/multimodal-generation/generation`, returning the completed image results as JSON in the submission response.
- Extend OpenAI Responses to all ten image models, including `stream`, `sync`, and `background` request modes, text and reference-image inputs, and image-generation parameters. Responses image requests use asynchronous upstream submission by default and can explicitly select synchronous upstream submission.
- Support synchronous ordinary generation of one to four images, with one image by default. Wan 2.7 sequential generation supports up to twelve images through `enable_sequential`; its asynchronous default is twelve images, while synchronous requests still default to one unless `n` is supplied.
- Support Wan 2.6 interleaved text/image generation with `enable_interleave` and a `max_images` limit of one to five, defaulting to five. Synchronous native requests return the complete text and image output in order as aggregate JSON.
- Support `1K` and `2K` size presets for ordinary `wan2.6-image` and Wan 2.7 image requests. `wan2.7-image-pro` additionally accepts `4K` for text-to-image without reference images or sequential generation. Explicit pixel sizes accept `width*height` or `widthxheight` and are checked against the selected model's pixel and aspect-ratio limits.
- Forward image controls including negative prompts, prompt extension, watermark, seed, sequential/interleaved generation, thinking mode, bounding boxes, and color palette. Preserve explicitly supplied `seed: 0` and false-valued options when combining compatibility fields with native parameters.
- Expose each successful image as a separate `image-1`, `image-2`, and subsequent task artifact, with individual content downloads. Responses renders text and image links in their original order.

### Changed

- Report image usage as the actual successful image count for synchronous and asynchronous tasks. Prefer upstream `usage.image_count`; if it is absent, count returned images. Partial-success responses count successful images without counting failed result entries.
- Validate image counts, interleaved limits, boolean options, reference-image counts, prompt structure, and size constraints before submission, including values supplied through `metadata.parameters`. Reject attempts to override the selected model through metadata, unsupported synchronous legacy requests, and upstream streaming outside synchronous Wan 2.6 interleaved generation.
- Reject incomplete or empty ordinary synchronous results. An asynchronous image task marked successful without image output becomes a failure; Wan 2.6 interleaved text-only results remain successful and report zero image usage.
- Reject malformed completion counts, including non-numeric, negative, fractional, or oversized values and zero counts accompanying image output, instead of reporting them as valid image usage.

### Migration

- **Price configuration:** Configure per-image task-usage pricing with `u("image_count")` for each image model you enable if you want billing by actual successful image count; switch from per-call pricing where necessary. The upgrade does not change saved prices or expressions automatically. Existing video prices and expressions require no changes solely for this release.
- **Installation:** This release requires a supporting new-api build. Upgrade unsupported installations before installing 1.3.0, using the [plugin API compatibility requirements](../../../../docs/plugin-api/v1.md) to check compatibility.
- **Call mode:** Responses images default to asynchronous upstream submission. Set `metadata.upstream_mode: "sync"` when a synchronous upstream call is needed; native synchronous calls use the new multimodal-generation route. Legacy image models support asynchronous submission only. Specify `n` explicitly when relying on a particular output count: synchronous calls and ordinary asynchronous Wan 2.7 calls default to one, ordinary asynchronous Wan 2.6 and legacy calls default to four, and asynchronous Wan 2.7 sequential calls default to twelve.
- **Interleaved output:** For synchronous Wan 2.6 interleaved generation, set `metadata.parameters.enable_interleave: true` on Responses requests or `parameters.enable_interleave: true` on the native synchronous route, and set `max_images` as needed. The plugin enables upstream streaming automatically. Do not set native top-level `stream: true`; the native response is aggregate JSON, and upstream frames are not forwarded directly to clients.
