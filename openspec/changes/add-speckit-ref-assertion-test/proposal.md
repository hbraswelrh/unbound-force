## Why

The upstream `speckit.*.md` command templates are listed in
`knownNonEmbeddedFiles` (`internal/scaffold/scaffold_test.go`
around lines 1242-1251). They are injected per-repo on `/uf.init`
rather than embedded in the binary. Per **ADR-002**
(`docs/decisions/002-uf-refs-source-of-truth-layer.md`, accepted
2026-08-27, merged via PR #550 / mergeCommit `7708bada`,
resolving decision #546; the ADR is authoritative but may not yet
be replicated to `origin/main` — inspect via
`git show 7708bada:docs/decisions/002-uf-refs-source-of-truth-layer.md`),
the in-scope set is the
**five upstream** files that `specify init` regenerates:
`speckit.specify.md`, `speckit.plan.md`, `speckit.tasks.md`,
`speckit.implement.md`, and `speckit.constitution.md`. (The four
UF-custom commands — analyze, checklist, clarify, taskstoissues —
are created by Step 5 and are out of scope; `speckit.testreview.md`
is not in `knownNonEmbeddedFiles` and is likewise out of scope.)
These templates carry load-bearing agent instructions: Dewey
knowledge-retrieval hooks (`dewey_semantic_search` /
`dewey_search`), constitution awareness
(`.specify/memory/constitution.md`), and the mandatory
`/uf.review-council` gate (implement phase only).

There is currently **no test** that asserts these references are
present. The nearest analog,
`TestGuardrailTemplates_CommandSpecificContent`
(`scaffold_test.go` ~6589-6692), asserts on `uf.init.md`
guardrail blocks, not on the speckit templates' reference
surface. As a result, the references can silently drift or
disappear — and because the files are injected per-repo, a
regression re-propagates to every downstream repo on the next
`/uf.init`, with no automated protection.

ADR-002 pins **Step 6 of `/uf.init`** (Speckit Command
Guardrails) in the embedded canonical asset
`internal/scaffold/assets/opencode/commands/uf.init.md` as the
authoritative injection layer. Step 6 does not yet inject the
Dewey and `/uf.review-council` references for the in-scope files;
`speckit.plan.md` already carries the compliant Dewey +
constitution pattern (the exemplar). This change adds the
regression gate; the sibling mechanical reconciliation edit
(issue #549) extends Step 6 to inject the missing references,
making the templates pass it. This change corresponds to issue
#548 (split from #445, depends on the now-closed source-of-truth
decision #546).

## What Changes

- Add a table-driven, per-file/per-reference content-assertion
  test in `internal/scaffold/scaffold_test.go` that asserts each
  of the five in-scope upstream `speckit.*.md` templates has its
  required references, modeled on the existing
  `TestGuardrailTemplates_CommandSpecificContent` pattern.
- The test asserts against the **durable source-of-truth layer**
  pinned by ADR-002 — Step 6 of the embedded canonical
  `internal/scaffold/assets/opencode/commands/uf.init.md`, read
  via `assetContent` — not a working-tree-only copy, so it
  provides real regression protection against downstream
  re-drift.
- Encode a per-file required-reference table so failure output
  names the exact missing reference on the exact file. Required
  references are conditional per file: Dewey and constitution
  awareness are required for all five in-scope files;
  `/uf.review-council` is required in `speckit.implement.md` and
  MUST NOT appear in the other four in-scope files.
- The test is a red-first gate: it fails against the current
  Step 6 injection surface and passes once the #549
  reconciliation edit extends Step 6.

## Capabilities

### New Capabilities
- `speckit-ref-assertion-test`: A regression test that verifies
  required UF references (Dewey, constitution awareness,
  `/uf.review-council`) are present in the speckit command
  templates at the durable injection layer, with per-file/
  per-reference precision in its failure output.

### Modified Capabilities
- None.

### Removed Capabilities
- None.

## Impact

- **Files**: `internal/scaffold/scaffold_test.go` (test-only
  addition). No production source changes.
- **Behavior**: Adds a new failing test that gates the #549
  reconciliation. Until #549 lands, this test is expected to
  fail red — it MUST be introduced in coordination with #549 (or
  the reconciliation performed in the same merge) so `main` is
  never left with a persistently red build.
- **Gates**: MUST NOT modify `knownNonEmbeddedFiles`, embedding
  expectations, coverage thresholds, or any other quality/
  governance gate.
- **Dependencies**: No new third-party dependencies; uses stdlib
  `testing` and `strings` only, consistent with the existing
  test conventions.

## Constitution Alignment

Assessed against the Unbound Force org constitution.

### I. Autonomous Collaboration

**Assessment**: N/A

This change adds a test and does not alter inter-hero artifact
formats or communication. It does not introduce runtime coupling.

### II. Composability First

**Assessment**: PASS

The test is self-contained within the scaffold test suite and
introduces no mandatory dependency on any other hero. It reads
static template/injection content only. Standalone functionality
is unaffected.

### III. Observable Quality

**Assessment**: PASS

The per-file/per-reference table produces machine-parseable,
reproducible failure output that names the exact missing
reference on the exact file. This backs the quality claim
("required references are present") with automated, re-runnable
evidence, directly serving the Observable Quality principle.

### IV. Testability

**Assessment**: PASS

The change is itself a test that verifies an observable property
(presence of required references) in isolation, without external
services or network access. It uses `t.Run` subtests for
per-file isolation and asserts on the durable source-of-truth
layer so the guarantee holds across downstream `/uf.init` runs.
It adds regression coverage where none existed and respects all
coverage/gate boundaries.
<!-- scaffolded by uf vdev -->
