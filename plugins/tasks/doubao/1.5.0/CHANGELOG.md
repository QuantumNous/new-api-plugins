---
changelogVersion: 1
plugin: "doubao"
version: "1.5.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.5.0]

### Added

- Generate and edit images with Seedream 5.5, `doubao-seedream-5-5-261028`. It makes 1K, 1.5K and 2K images and takes up to 10 reference images.
- Seedream 5.5 makes transparent images from text or from several reference images. Send `background: "auto"` to let the model decide, or `background: "transparent"` without the single-reference rule of Seedream 5.0 pro and flash.
- Seedream 5.5 layer decomposition takes `layer_decomposition_options.max_images`, from 2 to 100 images including the base image. Without it, Ark returns up to 30.

### Changed

- Layer decomposition on Seedream 5.5 holds the requested `max_images`, or 30 images, until the images arrive. Seedream 5.0 pro and flash still hold 17.
- Ark endpoint IDs that the plugin does not recognize as a listed model now accept `background: "auto"` and `layer_decomposition_options`, and their layer decomposition holds 30 images by default instead of 17.

### Migration

- Set prices for `doubao-seedream-5-5-261028` before you enable it for users. Its price table has the same fields as Seedream 5.0 pro: both image generation tiers, the input image price and the layer decomposition price.
