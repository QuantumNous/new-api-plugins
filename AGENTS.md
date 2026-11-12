# AGENTS.md — Plugin Repository Rules

## No plugin test code

- Do NOT create, modify, copy, or generate plugin test code anywhere in this repository, including under `plugins/`, `tools/`, or inside CI workflows.
- Keep plugin test fixtures, test harnesses, and dependencies or CI jobs used solely for plugin tests outside this repository. Run plugin behavior and billing tests in the host repository or a temporary directory outside this repository.
- Existing marketplace index, schema, and changelog validation commands may still be run.

## Changelogs for new versions (required)

- Every new plugin or version, including a version synchronized from another repository, MUST include `CHANGELOG.md` beside that version's `plugin.js`:

  ```text
  plugins/<kind>/<key>/<version>/plugin.js
  plugins/<kind>/<key>/<version>/CHANGELOG.md
  ```

- Repository rules and the canonical `CHANGELOG.md` MUST be written in English. Optional translations belong in `CHANGELOG.<locale>.md` in the same version directory. They supplement, rather than replace, the complete English changelog.
- Changelogs MUST follow the spec's [changelog format](https://github.com/QuantumNous/new-api/blob/main/docs/plugin-api/v1.md#changelog). The `plugin` and `version` metadata MUST match the directory and `plugin.js`. The English changelog MUST list every translation in its `translations` map.
- Changelogs MUST describe actual changes and their required migration steps: additions, behavior changes, deprecations, removals, fixes, and security changes. Include compatibility, default-parameter, and billing impacts of those changes where applicable. Initial releases MUST describe their initial capabilities. Do not include validation or test reports, verification limits, generic upgrade instructions, or lists of unrelated unchanged behavior; put those details in commit or PR descriptions. Empty files and placeholders are not acceptable.
- Changelogs MUST be readable by plugin users and administrators without reading the source code. Use natural sentences, one main change per bullet, and familiar feature names. Lead with what users can do or what behavior changes. Avoid literal translations, internal terminology, parameter inventories, and long sentences that combine unrelated changes. Include exact model names, paths, or fields only when they help users select a feature, call it, or upgrade correctly. Review English and translated prose for both readability and equivalent meaning.
- Describe only actions users need to take. Omit statements such as "no price reconfiguration is required", "existing settings are unchanged", or "no action is needed" from changelogs and handoffs. If an area needs no action and has no relevant behavior change, omit it.
- Include `Migration` when the release requires a user action. Put required pricing changes first: name the affected models/settings, explain the price impact, and give the concrete changes to prices, resolution tiers, billing mode, or expressions. Then describe other required compatibility steps. Explain automatic migration only when it affects a decision or remaining action. Omit `Migration` when no action is required.
- When synchronizing a plugin, synchronize its changelog too. If the source repository has no changelog, write one from the verified changes. A new version without its sibling changelog is incomplete.
- Apply these requirements to new and unpublished versions. Do not modify immutable published version directories just to backfill or reformat changelogs.
- Commit the new plugin, its changelog and translations, and the generated `index.json` together. Never edit the index by hand. Before completion, run `go run . check ../..` from `tools/pluginindex` and validate the changelog metadata and Markdown structure against the spec. The index checker only checks that listed translations exist and are named correctly.

## Plugin READMEs (required)

- Every plugin has `plugins/<kind>/<key>/README.md` (English) and `README.zh-CN.md`. `README.md` lists the translation in its front matter.
- A new version that changes channel setup, models, endpoints or billing updates both files in the same commit.
- READMEs follow the spec's [usage guide rules](https://github.com/QuantumNous/new-api/blob/main/docs/plugin-api/v1.md#usage-guide-readme): no HTML, no images.
