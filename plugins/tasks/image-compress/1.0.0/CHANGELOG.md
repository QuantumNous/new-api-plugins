---
changelogVersion: 1
plugin: "image-compress"
version: "1.0.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.0.0]

### Added

- Compress large images in requests before they are sent upstream.
- Use lossless mode, the default, to turn large PNG and GIF images into WebP without losing any pixels.
- Use lossy mode to re-encode large images as JPEG or PNG and scale down images above a maximum side length.
- Choose the minimum image size, the maximum side length, the output format and the JPEG quality.
