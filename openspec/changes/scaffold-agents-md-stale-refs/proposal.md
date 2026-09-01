## Why

`warnStaleCommandRefs()` (`internal/scaffold/scaffold.go:360`)
detects stale (pre-`uf.`-namespace) command references during
`uf init` but only scans `.opencode/agents/*.md`
(`scaffold.go:361`). It never scans the root `AGENTS.md`, which
`isToolOwned()` (`scaffold.go:418-436`) deliberately leaves
user-owned and therefore never auto-migrates. As a result, an
`AGENTS.md` left with old unprefixed command names (e.g.,
`/review-council` instead of `/uf.review-council`) produces no
warning — the exact defect reported in issue #454, of which
#567 is the durable-fix split (Option B).

A second, latent defect compounds this: the matcher at
`scaffold.go:392` uses `strings.Contains(content, oldRef)`.
Because `renamedCommands` (`scaffold.go:323`) yields a `refMap`
entry `/review-council → /uf.review-council`, the substring
`/review-council` matches *inside* the already-correct
`/uf.review-council`. Extending the scan to `AGENTS.md`
(which is saturated with `/uf.review-council`) on top of the
current substring logic would emit false-positive stale
warnings on every correctly-migrated repository. The
word-boundary fix is therefore a prerequisite that MUST land
together with the scan-set extension.

## What Changes

- Extend the scan set of `warnStaleCommandRefs()` to include
  root `AGENTS.md` alongside `.opencode/agents/*.md`, reusing
  the existing `renamedCommands`-derived `refMap` as the single
  source of truth (DRY, consistent with #419).
- Replace the substring matcher (`strings.Contains`,
  `scaffold.go:392`) with word-boundary-anchored matching so a
  correct `/uf.review-council` reference is not falsely flagged
  as containing the stale `/review-council`.
- Preserve the existing **warn-only** behavior and per-file
  read-error resilience (log + continue). No file is modified;
  the user-owned ownership contract enforced by `isToolOwned()`
  is honored.
- Add regression tests: a unit test for a stale ref located in
  `AGENTS.md`, an explicit negative test asserting a correct
  `/uf.review-council` ref produces no warning, and a
  self-reading drift-detection test asserting this repo's own
  `AGENTS.md` contains no stale command names.

Out of scope (deferred, per triage no-split decision): the
optional `--force` auto-rewrite path that would mutate the
user-owned `AGENTS.md`. It is only conditional in #567's body,
carries the sole data-corruption risk (double-prefixing,
non-idempotent rewrites), and will be tracked separately if and
when it is actually built (zero-waste).

## Capabilities

### New Capabilities
- `stale-ref-detection-agents-md`: `uf init` warns when the
  root `AGENTS.md` (and optionally bridge files) contains stale
  pre-namespace command references.
- `word-boundary-ref-matching`: stale-reference matching is
  token/word-boundary anchored, eliminating substring
  false positives such as `/review-council` inside
  `/uf.review-council`.

### Modified Capabilities
- `warnStaleCommandRefs`: scan set extended from
  `.opencode/agents/*.md` to also include root `AGENTS.md`;
  matcher hardened from substring to word-boundary.

### Removed Capabilities
- None.

## Impact

- **Files**: `internal/scaffold/scaffold.go`
  (`warnStaleCommandRefs`), `internal/scaffold/scaffold_test.go`
  (new unit + negative cases, modeled on
  `TestWarnStaleCommandRefs_StaleRef`,
  `scaffold_test.go:6800`), and a drift-detection
  test (modeled on
  `TestCheckAgentContext_ProjectAGENTSmdHasAllTier1Sections`,
  `doctor_test.go:3713`).
- **Behavior**: `uf init` gains one additional file read
  (`AGENTS.md`) and may emit additional warn-only output. No
  change to exit status; still non-blocking.
- **User-facing**: `uf init` warning output changes → file a
  website documentation issue before merge per the
  documentation gate. Update `CHANGELOG.md`.
- **Coordination**: shares the `renamedCommands` migration
  mechanism with #419 (closed) and resolves the reported symptom
  in #454 (parent, open).

## Constitution Alignment

Assessed against the Unbound Force org constitution (v1.2.0).

### I. Autonomous Collaboration

**Assessment**: N/A

This change affects local `uf init` scaffold behavior and its
terminal warning output. It introduces no new inter-hero
artifact exchange and does not alter the artifact envelope
format.

### II. Composability First

**Assessment**: PASS

The change is entirely contained within the standalone
`internal/scaffold` package and requires no other hero to be
present. `uf init` continues to deliver its core value alone;
no mandatory dependency is introduced.

### III. Observable Quality

**Assessment**: PASS

Warning output remains stable and human-readable at init time
(interactive scaffold feedback, not a machine-consumed
artifact). The change is backed by automated, reproducible
regression tests, and matching is deterministic across runs.

### IV. Testability

**Assessment**: PASS

The change verifies observable side effects (warning text via
an injected `io.Writer`) rather than implementation details.
All tests run in isolation using `t.TempDir()`; the
drift-detection test is self-locating and network-free. A
coverage strategy (unit, negative, drift) is defined here and
in `tasks.md`.

### V. Security by Default

**Assessment**: PASS

The change performs no file mutation and honors the user-owned
ownership contract (`isToolOwned()`), preserving least
privilege over user files. Word-boundary matching hardens input
handling against substring false positives. Existing read-error
resilience (log + continue) is preserved. No new dependency is
added — word-boundary matching uses the standard library.
<!-- scaffolded by uf vdev -->
