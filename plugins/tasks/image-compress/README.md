---
translations:
  zh-CN: README.zh-CN.md
---
# Image compression

Compresses large images in requests before the gateway sends them upstream. Smaller images upload faster, and lossy mode can also scale down images that are bigger than the model needs. The gateway keeps an image as it is when the new version would not be smaller.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Image compression** and fill in its **Settings**.
3. Under **Where it runs**, choose **All requests**, **Chosen groups** or **Channels that add it**. For the last, add the instance in each channel's settings under **Request processing plugins**. Under **Models**, you can limit it to some models.
4. Save. **Try it** in the editor runs the plugin on a sample request with an image and shows the original image next to the one sent upstream.

## Settings

- **Compression mode**: **Lossless** (the default) or **Lossy**.
  - **Lossless** keeps every pixel. It turns PNG and GIF images into WebP and never resizes them. JPEG and WebP images are left as they are. The upstream must accept WebP images.
  - **Lossy** re-encodes JPEG, PNG, GIF and WebP images and scales down those that are too large.
- **Minimum image size**: images at least this large are compressed. The default is 512 KB.
- **Maximum side length** (lossy only): an image whose longer side is above this is scaled down to fit, even when it is smaller than the minimum size. The default is 2048 pixels.
- **Output format** (lossy only): **Automatic** (the default) uses PNG for images with transparency and JPEG for the rest. You can also choose **JPEG** or **PNG**.
- **JPEG quality** (lossy only): from 1 to 100, 85 by default. Higher keeps more detail and saves less space.

## What clients see

- Requests are never rejected by this plugin. The model receives the compressed images.
- The usage log lists each image that was replaced, with its format and size before and after.
- If the plugin fails on a request, the images are sent as they were.

## Limits

- Only images sent inside the request as base64 data are compressed. Images given by URL or by file ID are left alone.
- It does not run on task submissions, such as video generation, or on requests the gateway sends upstream as a form, such as image edits.
- It does not run when the gateway or the channel passes the request body through unchanged.
- Animated GIFs, images over 20 MiB and images over 40 megapixels are sent as they are.
- Up to 32 images per request are compressed. The rest are sent as they are.
- Requires new-api v1.0.0-rc.50 or later.
