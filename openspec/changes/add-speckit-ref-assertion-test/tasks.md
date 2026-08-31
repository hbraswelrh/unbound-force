<!--
  [P] marks tasks eligible for parallel execution.
  Tasks without [P] run sequentially first, then [P]
  tasks run in parallel.
-->

## 1. Anchor to the ADR-002 source-of-truth layer (#546)

- [x] 1.1 Confirm ADR-002
  (`docs/decisions/002-uf-refs-source-of-truth-layer.md`,
  merged via PR #550) pins **Option B — extend Step 6** as the
  authoritative injection layer, and that the in-scope set is the
  five upstream files (specify, plan, tasks, implement,
  constitution).
- [x] 1.2 Identify the exact Step 6 injection block / marker in
  the embedded canonical asset
  `internal/scaffold/assets/opencode/commands/uf.init.md` that
  defines required refs per speckit file, and confirm it is
  reachable via `assetContent("opencode/commands/uf.init.md")`.

## 2. Author the content-assertion test

- [x] 2.1 In `internal/scaffold/scaffold_test.go`, add
  `TestSpeckitTemplates_RequiredReferences` (or similarly named)
  modeled on `TestGuardrailTemplates_CommandSpecificContent`
  (~6589-6692). Read the durable Step 6 layer via
  `assetContent("opencode/commands/uf.init.md")`.
- [x] 2.2 Define the per-file required-reference table
  (`[]struct{ file string; mustContain, mustNotContain
  []string }`) enumerating the five in-scope files and encoding
  the D4 / ADR-002 matrix: Dewey (`dewey_semantic_search` OR
  `dewey_search`) required for all five; `.specify/memory/
  constitution.md` required for all five; `/uf.review-council`
  as `mustContain` for `speckit.implement.md` and
  `mustNotContain` for the other four in-scope files.
- [x] 2.3 Implement per-file `t.Run(file, ...)` subtests with
  block-scoped matching (design D3) so failure output names the
  exact missing/leaked reference and the exact file. Use only
  stdlib `testing` and `strings`.

## 3. Verify red-first behavior and gates

- [x] 3.1 Run `go test -race -count=1 ./internal/scaffold/` and
  confirm the new test fails RED against the current
  pre-reconciliation Step 6 surface, naming the missing Dewey /
  `/uf.review-council` refs for the in-scope files.
- [x] 3.2 Verify no changes were made to `knownNonEmbeddedFiles`,
  embedding expectations, coverage thresholds, or any other
  quality/governance gate (Gatekeeping Integrity).
- [x] 3.3 Record the #549 coordination constraint (design D5:
  this test must merge with or immediately before the #549
  reconciliation so `main` is never persistently red) for
  `/uf.finale`.

## 4. Constitution alignment verification

- [x] 4.1 Verify the change honors the proposal's Constitution
  Alignment: Composability First (PASS — self-contained, no
  mandatory hero dependency), Observable Quality (PASS —
  machine-parseable per-file/per-ref failure output), and
  Testability (PASS — isolated, no external services,
  regression coverage added). Autonomous Collaboration is N/A.

## 5. CI parity

- [x] 5.1 Run `make check` (or the equivalent derived from
  `.github/workflows/`: `go vet ./...`, `golangci-lint run`,
  `go test -race -count=1 ./...`) locally before marking the
  change complete. Note: full-suite green requires the #549
  reconciliation; scope local verification accordingly.
<!-- scaffolded by uf vdev -->
<!-- spec-review: passed -->
<!-- code-review: passed -->
