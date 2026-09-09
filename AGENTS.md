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
- Include `Migration` when a change affects pricing, billing calculations, or required configuration. Its first entries MUST explain the pricing impact and explicitly state whether administrators need to reconfigure prices, resolution tiers, billing mode, or billing expressions. Identify the affected models/settings, any automatic migration, and the concrete required action before describing other compatibility steps. State that no price reconfiguration is required only when supported by the change; this targeted statement is allowed. Keep validation reports and generic release procedures out of the changelog.
- When synchronizing a plugin, synchronize its changelog too. If the source repository has no changelog, write one from the verified changes. A new version without its sibling changelog is incomplete.
- Apply these requirements to new and unpublished versions. Do not modify immutable published version directories just to backfill or reformat changelogs.
- Commit the new plugin, its changelog and translations, and the generated `index.json` together. Never edit the index by hand. Before completion, run `go run . check ../..` from `tools/pluginindex` and validate the changelog metadata and Markdown structure against the documented format. The index checker does not validate changelogs.
