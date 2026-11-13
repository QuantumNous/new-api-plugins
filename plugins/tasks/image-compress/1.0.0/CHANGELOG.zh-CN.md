---
changelogVersion: 1
plugin: "image-compress"
version: "1.0.0"
locale: "zh-CN"
---
# Changelog

## [1.0.0]

### Added

- 在请求发往上游前压缩其中的大图片。
- 默认的无损模式把较大的 PNG 和 GIF 图片转为 WebP，不损失任何像素。
- 有损模式把较大的图片重新编码为 JPEG 或 PNG，并缩小超过最长边的图片。
- 可以设置最小图片大小、最长边、输出格式和 JPEG 质量。
