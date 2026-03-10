## Why

`brew install unbound-force/tap/unbound` does not install graphthulhu,
the MCP knowledge graph server that `opencode.json` expects on `$PATH`.
Users must manually download and place the binary — a broken out-of-box
experience that violates the Composability First principle.

## What Changes

- Add `Casks/graphthulhu.rb` to `unbound-force/homebrew-tap` (v0.4.0,
  all four platform checksums already verified)
- Add a GoReleaser cask template (`.goreleaser-cask.rb.tmpl`) that
  injects `depends_on cask: "graphthulhu"` into every generated
  `Casks/unbound.rb` release
- Reference the template from `.goreleaser.yaml` via `cask_template`
- Open a best-effort PR to `skridlevsky/graphthulhu` proposing an
  upstream Homebrew tap so ownership can be transferred later

## Capabilities

### New Capabilities

- `graphthulhu-cask`: graphthulhu v0.4.0 installable via
  `brew install unbound-force/tap/graphthulhu`
- `unbound-graphthulhu-dependency`: installing `unbound` automatically
  pulls graphthulhu as a cask dependency — zero extra steps

### Modified Capabilities

- `unbound-cask`: GoReleaser now uses a cask template that adds
  `depends_on cask: "graphthulhu"` to every release of the unbound cask

### Removed Capabilities

None.

## Impact

- `unbound-force/homebrew-tap`: new file `Casks/graphthulhu.rb`
- `unbound-force/unbound-force`: new file `.goreleaser-cask.rb.tmpl`,
  updated `.goreleaser.yaml` (`homebrew_casks[].cask_template` field)
- `opencode.json`: no change needed; graphthulhu will now be on `$PATH`
  after install
- `skridlevsky/graphthulhu`: upstream PR (best-effort, no blocker)

## Constitution Alignment

Assessed against the Unbound Force org constitution v1.1.0.

### I. Autonomous Collaboration

**Assessment**: PASS

This change is purely a distribution/packaging concern. It does not
alter any hero artifact formats, inter-hero communication protocols,
or the artifact envelope structure. Heroes continue to collaborate
through well-defined files. graphthulhu is a read-only MCP server that
heroes query independently — no runtime coupling is introduced.

### II. Composability First

**Assessment**: PASS

This change directly advances Composability First. Previously, `unbound`
was not independently installable in a useful state (graphthulhu had to
be manually sourced). After this change, a single install command
produces a fully functional environment. The cask dependency is additive
— graphthulhu can still be installed and used standalone.

### III. Observable Quality

**Assessment**: N/A

This change does not affect hero output formats, provenance metadata, or
JSON artifact schemas. It is a packaging-only change.

### IV. Testability

**Assessment**: N/A

Homebrew cask files have no unit-testable logic. Correctness is verified
by `brew audit --cask graphthulhu` (linting) and `brew install --cask`
smoke tests. No application code is modified.
