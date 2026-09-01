<!--
  [P] marks tasks eligible for parallel execution.
  scaffold.go and scaffold_test.go are distinct files;
  doctor_test.go (drift test) is a third file. Tasks that
  touch the same file are NOT marked [P].
-->

## 1. Implementation (internal/scaffold/scaffold.go)

- [x] 1.1 Extend the scan set of `warnStaleCommandRefs()`
  (`scaffold.go:360`) to include the root `AGENTS.md`
  (`filepath.Join(targetDir, "AGENTS.md")`) alongside the
  existing `.opencode/agents/*.md` glob, iterating both through
  the same read + read-error-continue loop. Reuse the existing
  `renamedCommands`-derived `refMap` (no duplication).
- [x] 1.2 Replace substring matching (`strings.Contains`,
  `scaffold.go:392`) with word-boundary-anchored matching using
  the standard-library `regexp` (pattern built from
  `regexp.QuoteMeta(oldRef)` with a boundary/negative-`uf.`
  guard), so `/uf.review-council` is NOT flagged as containing
  `/review-council`. Same file as 1.1 — sequential.
- [x] 1.3 Preserve warn-only behavior and per-file read-error
  resilience (log + continue); confirm no code path mutates
  `AGENTS.md` or any user-owned file. Same file — sequential.
- [x] 1.4 Update the GoDoc comment on `warnStaleCommandRefs()`
  to reflect the extended scan set and word-boundary matching.
  Also generalize the user-visible warning header
  (`scaffold.go:403`, currently "Stale command references in
  agent files:") so it accurately covers the root `AGENTS.md`
  (e.g., "in project files"), and ensure the existing
  `uf init --force` guidance does not imply `--force` will
  rewrite the user-owned `AGENTS.md` (it MUST NOT — that path is
  deferred). Same file — sequential.

## 2. Tests (parallel across distinct files)

- [x] 2.1 [P] Add a unit test to
  `internal/scaffold/scaffold_test.go` for a stale
  `/review-council` reference in a `t.TempDir()` root
  `AGENTS.md`, asserting the warning is emitted and the file is
  unmodified (model:
  `TestWarnStaleCommandRefs_StaleRef`, `scaffold_test.go:6800`;
  stdlib `testing`, `bytes.Buffer` writer).
- [x] 2.2 [P] Add a negative unit test to
  `internal/scaffold/scaffold_test.go` asserting that content
  containing only `/uf.review-council` produces NO warning
  (substring-safety guard for task 1.2). Same file as 2.1 —
  NOT parallel with 2.1; sequence 2.1 then 2.2.
- [x] 2.3 [P] Add a drift-detection test (in the appropriate
  `_test.go` file, e.g. `doctor_test.go`) that walks to
  `go.mod`, reads this repo's own root `AGENTS.md`, and asserts
  zero stale command names remain; skip gracefully when
  `AGENTS.md`/`go.mod` is absent (model:
  `TestCheckAgentContext_ProjectAGENTSmdHasAllTier1Sections`,
  `doctor_test.go:3713`). The drift check MUST iterate every
  `refMap` key (all renamed commands), not only `/review-council`,
  for full regression coverage.

<!-- Note: 2.1 and 2.2 share scaffold_test.go and MUST run
     sequentially relative to each other. 2.3 is in a different
     file and is genuinely parallel-eligible against 2.1/2.2. -->

## 3. Verification & documentation

- [x] 3.1 Run CI parity locally:
  `go test -race -count=1 ./internal/scaffold/...` and the
  drift test package, plus `go vet ./...` and
  `golangci-lint run` (derive exact commands from
  `.github/workflows/`). All MUST pass.
- [x] 3.2 Verify constitution alignment holds as declared in
  `proposal.md`/`design.md`: II (Composability) PASS,
  III (Observable Quality) PASS, IV (Testability) PASS,
  V (Security by Default — no mutation, stdlib-only, least
  privilege) PASS, I (Autonomous Collaboration) N/A.
- [x] 3.3 Update `CHANGELOG.md` with an entry for the extended
  `uf init` stale-reference detection. The entry MUST include a
  `Spec:` reference to this OpenSpec change
  (`openspec/changes/scaffold-agents-md-stale-refs/`) per the
  documentation gate.
- [x] 3.4 File a website documentation issue for the
  user-facing `uf init` warning-output change (documentation
  gate) before PR merge.
<!-- scaffolded by uf vdev -->

<!-- spec-review: passed -->
<!-- code-review: passed -->
