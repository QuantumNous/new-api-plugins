---
changelogVersion: 1
plugin: "alibaba"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Added

- Complete `wan2.2-kf2v-flash` and `wan2.2-s2v` support, including the native `/ali/api/v1/services/aigc/image2video/video-synthesis` submission route and digital-human results at `output.results.video_url`.
- Declare the existing Wan 2.6 regional variants and Wan 2.7 dated versions alongside the supported Wan 2.1 through 3.0 models.

### Changed

- Apply model-specific resolution, aspect-ratio, duration, and media rules after merging request parameters. Native protocol fields take precedence over compatibility fields.
- Accept pixel `size` values with `*` or `x` separators and resolution tiers in compatibility requests. Preserve the corresponding aspect ratio when the target model supports it; image-driven models retain the input media's aspect ratio.
- Default `wan2.7-t2v` to `resolution: 1080P` and `ratio: 16:9` instead of the legacy `size: 1280*720` parameter.
- Default `wan2.7-i2v` to `1080P` instead of `720P`.
- Default `wan2.6-t2v` and `wan2.6-t2v-us` to `size: 1920*1080` instead of `size: 1280*720`.
- Default `wan2.2-s2v` to `480P` instead of `720P`. Duration follows the audio rather than an assumed five seconds; reserve 20 seconds and read the actual fractional duration at completion.
- Retain the aspect ratio of explicit Wan3 pixel sizes instead of falling back to `adaptive`.
- Reserve 30 seconds for Wan3 requests with reference video or smart duration. Completion usage includes input and output video duration; Wan2.7 continuation uses the full returned video duration.

- Reject malformed or oversized completion usage through host validation instead of clamping oversized duration to 3600 seconds.

### Fixed

- Stop validating `wan2.7-t2v` resolution tiers as legacy pixel sizes when `size` is omitted.
- Preserve native `input` and `parameters`, including explicit `prompt_extend: false`, `watermark: false`, `audio: false`, and `seed: 0`.
- Correct precedence between metadata duration and unified request fields while preserving Wan3 smart duration.
- Use the digital-human result URL consistently for task results and artifact downloads.

### Migration

- **Price configuration:** This upgrade does not automatically change administrator prices or migrate billing expressions. Existing task-usage expressions based on `u("seconds")` and `u("resolution")` remain compatible and need no syntax rewrite solely for this upgrade; configure any missing prices for the models and resolution tiers enabled with this release.
- **Billing mode:** If you currently use per-call pricing and want actual-duration settlement for `wan2.2-s2v` or Wan3, switch those models to task-usage expressions based on `u("seconds")` and configure the applicable per-second prices. The plugin does not switch billing modes automatically.
- **Wan3 expressions:** `u("seconds")` now includes input and output video duration. If a custom expression already adds input-video duration separately, remove that extra calculation to avoid charging twice.
- **Resolution price tiers:** The default output for `wan2.7-t2v`, `wan2.7-i2v`, `wan2.6-t2v`, and `wan2.6-t2v-us` changes to 1080P; confirm the configured 1080P prices. To retain the previous 720P output, explicitly set `parameters.resolution: "720P"` in native Wan2.7 requests or `parameters.size: "1280*720"` in Wan2.6 text-to-video requests. Also confirm the 480P price for the new `wan2.2-s2v` default.
