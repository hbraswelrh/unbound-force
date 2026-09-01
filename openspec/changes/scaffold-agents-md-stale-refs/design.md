## Context

`warnStaleCommandRefs()` (`internal/scaffold/scaffold.go:360`)
is invoked from `Run()` (`scaffold.go:257`) at the end of
`uf init`. It builds a `refMap` of old→new command names from
the `renamedCommands` variable (`scaffold.go:323`), globs
`.opencode/agents/*.md` (`scaffold.go:361`), and for each file
reports any `refMap` key found via
`strings.Contains(content, oldRef)` (`scaffold.go:392`). Output
is warn-only; per-file read errors are logged and skipped
(`scaffold.go:385-388`).

Two gaps motivate this change (see `proposal.md`):
1. The root `AGENTS.md` — user-owned per `isToolOwned()`
   (`scaffold.go:418-436`, returns `false`) and never
   auto-migrated — is outside the scan set, so stale refs there
   go unwarned (issue #454).
2. `strings.Contains` is substring-based: the `refMap` entry
   `/review-council → /uf.review-council` causes the substring
   `/review-council` to match inside the already-migrated
   `/uf.review-council`, producing false positives once
   `AGENTS.md` (saturated with `/uf.review-council`) enters the
   scan set.

## Goals / Non-Goals

### Goals
- Add root `AGENTS.md` to the `warnStaleCommandRefs()` scan set,
  reusing the existing `renamedCommands`-derived `refMap`
  (single source of truth, DRY).
- Replace substring matching with word-boundary-anchored
  matching so migrated refs are not falsely flagged.
- Preserve warn-only behavior and per-file read-error
  resilience; honor the `isToolOwned()` user-owned contract
  (no mutation).
- Add unit, negative, and drift-detection regression tests.

### Non-Goals
- No `--force` auto-rewrite of `AGENTS.md` (deferred; mutates a
  user-owned file, carries corruption risk, only conditional in
  #567). Tracked separately if built. The existing
  `uf init --force` guidance in the warning output MUST NOT be
  extended to imply it rewrites `AGENTS.md`; the warning for
  `AGENTS.md` remains advisory (manual update only).
- No change to `renamedCommands` contents or to `isToolOwned()`
  ownership classification.
- No new JSON/machine-readable artifact — output remains
  interactive init-time terminal feedback.
- Bridge files (`CLAUDE.md`, `.cursorrules`) are optional; if
  included they reuse the identical mechanism, but they are not
  required to resolve #454.

## Decisions

**D1 — Scan set as an explicit file list plus the agents
glob.** Keep the existing `.opencode/agents/*.md` glob and
append the root `AGENTS.md` absolute path
(`filepath.Join(targetDir, "AGENTS.md")`) to the iterated
match set. Rationale: minimal, localized change; the existing
per-file read + read-error-continue loop is reused unchanged.
`filepath.Base()` in the warning output continues to render a
clean file label (`AGENTS.md`).

**D2 — Word-boundary matching via a compiled pattern per
ref.** Replace `strings.Contains(content, oldRef)` with a
match that requires the stale ref to be a standalone token —
i.e., not immediately preceded by the `uf.` prefix (or, more
generally, not part of a longer command token). Use
`regexp` from the standard library with a boundary-anchored
pattern built from `regexp.QuoteMeta(oldRef)`, or an
equivalent hand-rolled boundary check. The decisive test is
that `/uf.review-council` MUST NOT match `/review-council`.
Patterns MUST be compiled once per ref (outside the per-file
scan loop) so the two scan sources — `.opencode/agents/*.md`
and root `AGENTS.md` — do not recompile the same patterns
redundantly. Rationale: standard library only (Constitution V
— no new dependency); deterministic; directly testable.

**D3 — Warn-only, no mutation.** `warnStaleCommandRefs()`
continues to only write to the injected `io.Writer`. This
honors the `isToolOwned()` contract (root `AGENTS.md` is
user-owned) and keeps the change side-effect-free on disk.
Rationale: least privilege (Constitution V); avoids the
corruption risk that motivated deferring the `--force` path.

**D4 — Three regression tests.**
- Unit: stale `/review-council` in a `t.TempDir()` `AGENTS.md`
  → warning emitted (model:
  `TestWarnStaleCommandRefs_StaleRef`,
  `scaffold_test.go:6800`).
- Negative: only `/uf.review-council` present → no warning
  (guards D2).
- Drift: read this repo's own `AGENTS.md`, assert zero stale
  refs (model:
  `TestCheckAgentContext_ProjectAGENTSmdHasAllTier1Sections`,
  `doctor_test.go:3713`; walk to `go.mod`, skip gracefully if
  absent).
Rationale: Constitution IV — observable side effects via
injected writer, isolated with `t.TempDir()`, network-free.

## Risks / Trade-offs

- **False positives if D2 is incomplete.** If word-boundary
  matching does not account for the `uf.` prefix specifically,
  the AGENTS.md scan would warn on correct repos. Mitigation:
  the negative test (D4) is a hard gate; matching MUST be
  substring-safe before the scan extension is considered
  correct.
- **Regex vs. hand-rolled boundary check.** `regexp` adds
  compile cost per ref, but the ref set is tiny (~10 entries)
  and `uf init` is one-shot; performance impact is negligible
  (confirmed acceptable in triage). Accepted trade-off for
  clarity and correctness.
- **Drift test coupling to repo layout.** The self-reading
  drift test depends on locating `go.mod`; it must skip
  gracefully when run outside the repo (matching the
  `doctor_test.go` precedent) to avoid brittle failures.
- **User-facing output change.** `uf init` warning output
  changes → a website documentation issue MUST be filed before
  merge, and `CHANGELOG.md` updated (documentation gate).

Constitution alignment is inherited from `proposal.md`:
II (Composability) PASS — self-contained in `internal/scaffold`;
III (Observable Quality) PASS — stable warn output + automated
regression tests; IV (Testability) PASS — isolated, side-effect
tests; V (Security by Default) PASS — no mutation, stdlib-only,
least privilege; I (Autonomous Collaboration) N/A.
<!-- scaffolded by uf vdev -->
