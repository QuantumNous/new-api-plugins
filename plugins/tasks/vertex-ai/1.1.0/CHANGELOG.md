---
changelogVersion: 1
plugin: "vertex-ai"
version: "1.1.0"
locale: "en"
translations:
  zh-CN: "CHANGELOG.zh-CN.md"
---
# Changelog

## [1.1.0]

### Added

- Use another New API server as the upstream for Google Veo video generation on Vertex AI. Enable this version on both servers and bind the plugin to a New API channel that holds the upstream address and API key. The upstream server runs every request in its own Google Cloud project.
- Send requests in the Vertex AI format to `/vertex-ai/v1/projects/{project}/locations/{location}/publishers/google/models/{model}:predictLongRunning` with a prompt and an optional first-frame `image`, and poll with `:fetchPredictOperation` on the same path. The returned operation name ends with the task ID and repeats the project and location from your request path.

### Migration

- **Installation:** Install 1.1.0 on new-api v1.0.0-rc.39 or later; older versions reject the plugin.
