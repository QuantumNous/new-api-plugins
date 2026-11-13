---
translations:
  zh-CN: README.zh-CN.md
---
# Automatic reasoning effort

Sets the reasoning effort of each request from its size. You write rules such as "requests with tools think harder", and the gateway applies the effort that the first matching rule names. The gateway turns that effort into the right field for each model's API, the same way an `@effort` model modifier does.

## Setup

1. In the console, open **Request processing** and choose **Add**.
2. Pick **Automatic reasoning effort** and fill in its **Settings**.
3. Under **Where it runs**, choose **All requests**, **Chosen groups** or **Channels that add it**. For the last, add the instance in each channel's settings under **Request processing plugins**.
4. Under **Models**, choose **Chosen models** and pick models that support reasoning. Other models may refuse a request that carries a reasoning effort.
5. Save. **Try it** in the editor runs the plugin on a sample request with the current settings, and **Dry run** on the **Request processing** page runs the saved instances.

## Settings

- **Rules**: checked from top to bottom. The first rule whose conditions all hold sets the effort. A rule without conditions matches every request. Up to 50 rules.
- Each rule picks a **Reasoning effort** and can set these conditions:
  - **Minimum messages**: the request has at least this many messages.
  - **Minimum tools**: the request has at least this many tools.
  - **Last user message length**: the last user message has at least this many characters.
  - **Total text length**: all text in the request together has at least this many characters.
  - **Images only**: when on, only requests with images match.
- The default rules give **High** effort to requests with at least one tool and to requests with at least 4,000 characters of text.
- **Default effort**: the effort for requests that no rule matches. It is **Not set** by default, so those requests keep the effort they have.
- **Lowest effort** and **Highest effort**: keep the effort the plugin sets within these limits. Both are **Not set** by default.
- **Replace the request's effort**: off by default, so a request that already sets an effort, in its body or with an `@effort` model modifier, keeps it. Turn it on to always apply the plugin's effort.

The efforts, from weakest to strongest, are **Off** (thinking off), **Minimal**, **Low**, **Medium**, **High**, **Extra high** and **Max**.

## What clients see

- The plugin never rejects a request. It only adds or changes the reasoning effort.
- The effort the plugin set shows in the request's usage log.
- If the plugin fails on a request, the request goes on with the effort it already had.

## Limits

- It runs on Chat Completions, Responses, Claude Messages and Gemini requests.
- It does not run when the gateway or the channel passes the request body through unchanged.
- A model without reasoning support may refuse the request, so limit the plugin to reasoning models.
- Requires new-api v1.0.0-rc.50 or later.
