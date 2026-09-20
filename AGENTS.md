# AGENTS.md — Plugin Repository Rules

## Changelogs for new versions (required)

- Every new plugin or version, including a version synchronized from another repository, MUST include `CHANGELOG.md` beside that version's `plugin.js`:

  ```text
  plugins/<kind>/<key>/<version>/plugin.js
  plugins/<kind>/<key>/<version>/CHANGELOG.md
  ```

- Repository rules and the canonical `CHANGELOG.md` MUST be written in English. Optional translations belong in `CHANGELOG.<locale>.md` in the same version directory. They supplement, rather than replace, the complete English changelog.
- Changelogs MUST follow the [release notes format](docs/marketplace.md#release-notes): UTF-8 Markdown, YAML front matter with stable English keys, one matching version heading, fixed English category headings, and bullet entries. Keep these structural elements in English even when the entry text is translated.
- Required metadata: `changelogVersion`, `plugin`, `version`, and `locale`. The plugin and version MUST match the directory and `plugin.js` metadata.
- The canonical changelog MUST list any translated files in its optional `translations` map so consumers can discover them without listing the directory. Translation metadata and release categories MUST match the English entry, except for `locale` and translated prose.
- Changelogs MUST describe actual changes and their required migration steps: additions, behavior changes, deprecations, removals, fixes, and security changes. Include compatibility, default-parameter, and billing impacts of those changes where applicable. Initial releases MUST describe their initial capabilities. Do not include validation or test reports, verification limits, generic upgrade instructions, or lists of unrelated unchanged behavior; put those details in commit or PR descriptions. Empty files and placeholders are not acceptable.
- Changelogs MUST be readable by plugin users and administrators without reading the source code. Use natural sentences, one main change per bullet, and familiar feature names. Lead with what users can do or what behavior changes. Avoid literal translations, internal terminology, parameter inventories, and long sentences that combine unrelated changes. Include exact model names, paths, or fields only when they help users select a feature, call it, or upgrade correctly. Review English and translated prose for both readability and equivalent meaning.
- Describe only actions users need to take. Omit statements such as "no price reconfiguration is required", "existing settings are unchanged", or "no action is needed" from changelogs and handoffs. If an area needs no action and has no relevant behavior change, omit it.
- Include `Migration` when the release requires a user action. Put required pricing changes first: name the affected models/settings, explain the price impact, and give the concrete changes to prices, resolution tiers, billing mode, or expressions. Then describe other required compatibility steps. Explain automatic migration only when it affects a decision or remaining action. Omit `Migration` when no action is required.
- When synchronizing a plugin, synchronize its changelog too. If the source repository has no changelog, write one from the verified changes. A new version without its sibling changelog is incomplete.
- Apply these requirements to new and unpublished versions. Do not modify immutable published version directories just to backfill or reformat changelogs.
- Commit the new plugin, its changelog and translations, and the generated `index.json` together. Never edit the index by hand. Before completion, run `go run . check ../..` from `tools/pluginindex` and validate the changelog metadata and Markdown structure against the documented format. The index checker does not validate changelogs.
