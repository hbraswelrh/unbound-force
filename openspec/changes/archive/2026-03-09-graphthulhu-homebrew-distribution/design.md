## Context

`opencode.json` configures graphthulhu as an MCP server with
`"command": ["graphthulhu", "serve", ...]`. The binary must be on
`$PATH` for OpenCode to launch it. Currently the `unbound` Homebrew
cask installs only the `unbound` binary; graphthulhu must be obtained
separately. This creates a broken first-run experience.

GoReleaser auto-generates `Casks/unbound.rb` in `unbound-force/homebrew-tap`
on each release. The generated cask has no `depends_on` field by default.
GoReleaser v2 supports a `cask_template` path in the `homebrew_casks`
stanza, which lets us inject arbitrary Ruby into the generated cask —
including `depends_on cask: "graphthulhu"`.

graphthulhu v0.4.0 (2026-02-06) has pre-built binaries for all four
supported platforms. The project (`skridlevsky/graphthulhu`) has no
upstream Homebrew tap, so we host the cask in our own tap.

## Goals / Non-Goals

### Goals

- `brew install unbound-force/tap/unbound` installs both `unbound` and
  `graphthulhu` with no extra steps
- graphthulhu cask is pinned to a known-good version with verified checksums
- The dependency survives every GoReleaser release automatically via a
  template file checked into the repo
- A best-effort upstream PR is opened to `skridlevsky/graphthulhu` to
  facilitate eventual ownership transfer

### Non-Goals

- Updating graphthulhu automatically when new versions release (manual
  cask bump or a future workflow can handle this)
- Supporting Windows (graphthulhu is a CLI tool; Windows Homebrew support
  is out of scope for now)
- Blocking the `unbound` release process on upstream graphthulhu acceptance

## Decisions

### Decision 1: GoReleaser cask template over post-release patch

**Chosen**: Add `.goreleaser-cask.rb.tmpl` and reference it via
`homebrew_casks[].cask_template` in `.goreleaser.yaml`.

**Rationale**: GoReleaser v2 supports `cask_template` as a first-class
field. The template is checked into the source repo, versioned with the
code, and applied automatically on every release. No second pipeline step
(e.g., a patching Action) is required. The template file makes the
dependency explicit and auditable.

**Alternatives considered**:
- Post-release GitHub Action: adds a second pipeline step that could fail
  independently; harder to audit; more moving parts.
- Manual cask management: removes all automation; high maintenance burden.

### Decision 2: Host graphthulhu cask in unbound-force/homebrew-tap

**Chosen**: Add `Casks/graphthulhu.rb` directly to our tap.

**Rationale**: `skridlevsky/graphthulhu` has no Homebrew tap. We cannot
`depends_on` a cask that doesn't exist somewhere. Hosting it in our tap
gives us full control and unblocks the install story immediately. The
cask is straightforward — version + 4 checksums.

**Alternatives considered**:
- Wait for upstream: blocks delivery indefinitely; no guarantee of
  upstream cooperation.
- Bundle graphthulhu binary inside the unbound archive: violates
  Composability First (forces a coupling between unbound and graphthulhu
  versions at build time) and bloats the unbound release.

### Decision 3: graphthulhu version pinning strategy

**Chosen**: Pin to `0.4.0` with hardcoded checksums in `graphthulhu.rb`.
Update manually (or via a future workflow) when new versions release.

**Rationale**: Homebrew casks are always version-pinned. A livecheck
stanza could automate version tracking, but adds complexity not needed
now. The cask includes a comment documenting where to get updated
checksums.

### Decision 4: Open upstream PR to skridlevsky/graphthulhu

**Chosen**: Open a best-effort PR proposing a `Casks/graphthulhu.rb` in
an upstream tap. Not a blocker.

**Rationale**: If skridlevsky adds their own tap in the future, we can
switch the `depends_on` to point there. This keeps the door open for
proper upstream ownership without blocking our delivery.

## Risks / Trade-offs

**Risk**: GoReleaser `cask_template` field behavior changes in a future
version.
→ Mitigation: pin GoReleaser version in CI; test release dry-run in CI
with `goreleaser release --snapshot`.

**Risk**: graphthulhu releases a new version and our cask becomes stale.
→ Mitigation: document the update process in the cask file header; a
future workflow can automate bumping.

**Risk**: `skridlevsky/graphthulhu` adds a competing cask in a different
tap, causing a conflict.
→ Mitigation: if they publish their own tap, switch `depends_on` to point
to theirs and remove our copy. Low risk near-term.

**Trade-off**: Composability First (principle II) says heroes MUST NOT
require another hero as a hard prerequisite. graphthulhu is not a hero —
it is an infrastructure tool. The `depends_on` is a packaging dependency,
not a runtime coupling. No hero artifact formats or inter-hero protocols
are affected. This is consistent with the principle's intent.
