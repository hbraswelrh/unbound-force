# Tasks: reconcile-speckit-uf-refs

## 1. Edit canonical source (`.opencode/commands/uf.init.md`)

- [x] 1.1 Locate the Step 6 region (between `### Step 6:` and
  `### Step 7:`) and the three fenced `markdown` blocks:
  "Spec-phase guardrails block", "Implement guardrails block",
  "Constitution guardrails block".
- [x] 1.2 Inject into the **Spec-phase guardrails block**: a Dewey
  reference (`dewey_semantic_search` OR `dewey_search`) and
  `.specify/memory/constitution.md`. Do NOT add `/uf.review-council`.
- [x] 1.3 Inject into the **Implement guardrails block**: Dewey
  reference, `.specify/memory/constitution.md`, and
  `/uf.review-council`.
- [x] 1.4 Inject into the **Constitution guardrails block**: Dewey
  reference and `.specify/memory/constitution.md`. Do NOT add
  `/uf.review-council`.

## 2. Lockstep dual-write to embedded asset

- [x] 2.1 Mirror the exact edits into
  `internal/scaffold/assets/opencode/commands/uf.init.md` so it is
  byte-identical to the canonical source (D6).

## 3. Verify

- [x] 3.1 Run `go test -race -count=1 ./internal/scaffold/...` and
  confirm `TestSpeckitTemplates_RequiredReferences` (#548) is GREEN.
- [x] 3.2 Confirm `TestEmbeddedAssets_MatchSource` (drift) is GREEN.
- [x] 3.3 Confirm `TestGuardrailTemplates_CommandSpecificContent`
  is GREEN — no `/uf.review-council` leak into spec-phase or
  constitution blocks (implement-only boundary preserved).
- [x] 3.4 Run full CI parity locally: `make check` (or derive from
  `.github/workflows/`).

## 4. Constitution alignment verification

- [x] 4.1 Confirm no gatekeeping value was modified (thresholds,
  severity, CI flags, review limits) — this change satisfies an
  existing gate, it does not weaken one.
- [x] 4.2 Confirm Testability (IV): the change is verified by
  stdlib-only, isolated unit tests reading the embedded asset.

## 5. Co-merge bookkeeping (design D5)

- [x] 5.1 Ensure the #548 test commit and this #549 fix commit are
  on the same branch (`opsx/add-speckit-ref-assertion-test`).
- [x] 5.2 Update the PR description with the co-merge rationale and
  scope boundaries (pre-empt scope-creep objections). (PR text
  delivered as markdown and applied manually — environment
  forbids `gh` mutations.)
- [x] 5.3 Cross-link issues #548 and #549 and note the co-merge in
  both issue descriptions. (Updated issue text delivered as
  markdown and applied manually — environment forbids `gh`
  mutations.)
