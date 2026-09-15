---
changelogVersion: 1
plugin: "hailuo"
version: "1.2.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.2.0]

### Changed

- Accept only the output resolutions a model can actually render: `768P` and `2K` for `MiniMax-H3`, `768P` and `1080P` for `MiniMax-Hailuo-2.3` and `MiniMax-Hailuo-2.3-Fast`, and `512P`, `768P`, and `1080P` for `MiniMax-Hailuo-02`. Asking one of these models for another resolution, such as `720P` or `2K` on `MiniMax-Hailuo-02`, now fails instead of quietly rendering and billing at `768P`. A pixel `size` such as `1280x720` still resolves to the closest supported resolution, and the 01-series models (`T2V-01`, `T2V-01-Director`, `I2V-01`, `I2V-01-Director`, `I2V-01-live`, and `S2V-01`) still accept any resolution.

### Migration

- **Price configuration:** Existing prices stay in effect and billed amounts do not change. Each model now offers only the prices it can charge: `MiniMax-H3` prices video seconds, output resolution, input images, and input video seconds; `MiniMax-Hailuo-2.3`, `MiniMax-Hailuo-2.3-Fast`, and `MiniMax-Hailuo-02` price video seconds and their own resolutions; the 01-series models price video seconds only. If you priced input images or input video on a model other than `MiniMax-H3`, fold those amounts into its per-second or per-resolution price the next time you save that model's price, and drop prices for resolutions the model cannot render.
- **Installation:** Install 1.2.0 on a new-api version that supports per-model billing parameters; an older gateway rejects the plugin.
