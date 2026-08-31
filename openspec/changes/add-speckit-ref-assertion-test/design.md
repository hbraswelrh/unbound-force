## Context

The upstream `speckit.*.md` command templates are non-embedded
(listed in `knownNonEmbeddedFiles`,
`internal/scaffold/scaffold_test.go` ~1242-1251) and injected
per-repo via `/uf.init`. They must carry required UF references
(Dewey hooks, constitution awareness, and `/uf.review-council`
for the implement phase only). No test asserts these references
today, so drift re-propagates downstream on every `/uf.init`.

The source-of-truth layer is now settled. **ADR-002**
(`docs/decisions/002-uf-refs-source-of-truth-layer.md`, Accepted
2026-08-27, resolving issue #546) pins **Option B — extend
Step 6 of `/uf.init` (Speckit Command Guardrails)** as the
authoritative layer for durable UF customization of the speckit
command files. The sibling edit (#549) targets Step 6 in the
canonical **embedded** asset
`internal/scaffold/assets/opencode/commands/uf.init.md` (which
IS embedded and drift-tested), NOT the installed
`.opencode/commands/` copies.

### In-scope files (per ADR-002)

ADR-002 scopes the required-reference guarantee to the **five
upstream** speckit files that `specify init` regenerates:

- `speckit.specify.md`
- `speckit.plan.md` (compliant exemplar)
- `speckit.tasks.md`
- `speckit.implement.md`
- `speckit.constitution.md`

The remaining four `knownNonEmbeddedFiles` entries
(`speckit.analyze.md`, `speckit.checklist.md`,
`speckit.clarify.md`, `speckit.taskstoissues.md`) are UF-custom
commands created by Step 5 and are **out of scope** for this
test, consistent with ADR-002. (`speckit.testreview.md`, present
in the working tree, is not in `knownNonEmbeddedFiles` and is
also out of scope.)

### Injection layers (verified in `uf.init.md`)

| Step | Role | Mutating? |
|------|------|-----------|
| Step 5 (`uf.init.md:452`) | Creates the 4 UF-custom commands | Yes (create) |
| Step 6 (`uf.init.md:546`) | **Authoritative (ADR-002):** idempotent, marker-based inject/correct into speckit commands | Yes (inject/correct) |
| Step 7 (`uf.init.md:677`) | Read-only verifier — reports missing refs, does not modify | No (`uf.init.md:704`) |

This change (issue #548) adds the regression test that gates the
reconciliation. The sibling edit (#549) extends Step 6 to inject
the references, making the templates pass. The decision
dependency (#546) is closed COMPLETED and its ADR-002 is merged
(PR #550, mergeCommit `7708bada`). ADR-002 is accepted and
authoritative but may not yet be replicated to `origin/main`
(merge-queue/replication lag); inspect it via
`git show 7708bada:docs/decisions/002-uf-refs-source-of-truth-layer.md`
if it is not present in the working tree. The proposal's
Constitution Alignment (PASS on
Composability, Observable Quality, Testability; N/A on
Autonomous Collaboration) governs this design.

## Goals / Non-Goals

### Goals
- Add a table-driven, per-file/per-reference content-assertion
  test in `internal/scaffold/scaffold_test.go`, modeled on
  `TestGuardrailTemplates_CommandSpecificContent` (~6589-6692).
- Assert against the **durable source-of-truth layer** — Step 6
  of the embedded canonical
  `internal/scaffold/assets/opencode/commands/uf.init.md` (per
  ADR-002) — not a working-tree-only copy, so the guarantee
  survives downstream `/uf.init`.
- Produce precise failure output naming the exact missing
  reference on the exact template file.
- Encode conditional per-file required references
  (`/uf.review-council` implement-only, with a MUST-NOT-appear
  assertion elsewhere; Dewey + constitution for all five
  in-scope files).
- Be red before the #549 reconciliation and green after it.

### Non-Goals
- Performing the reconciliation edit itself (that is #549).
- Modifying `knownNonEmbeddedFiles`, embedding expectations,
  coverage thresholds, or any other quality/governance gate.
- Adding a full regenerate-cycle (`specify init` → `/uf.init`)
  integration test — the content assertion is scoped to the
  Step 6 injection-instruction surface; the end-to-end
  regenerate test is a separate MEDIUM acceptance criterion on
  the sibling edit issue (#549) per ADR-002 follow-up.
- Introducing any new third-party dependency.
- Asserting references for the four out-of-scope UF-custom
  speckit commands.

## Decisions

**D1 — Assertion target: Step 6 of the embedded canonical
`uf.init.md` (per ADR-002, Option B).**
Per ADR-002, Step 6 (Speckit Command Guardrails) is the single
authoritative, idempotent, marker-based mutating locus for
durable UF customization of speckit command files. The test
asserts that the Step 6 injection surface in the embedded
canonical asset defines the required references for each in-scope
speckit file, read via
`assetContent("opencode/commands/uf.init.md")` — the same
embedded-asset pattern already used by
`TestGuardrailTemplates_CommandSpecificContent`. This makes the
assertion durable: if the Step 6 required-ref surface is ever
weakened, the test fails. Step 7 (read-only verifier) is NOT the
target; it cannot guarantee presence downstream.

Rationale for the red→green contract: Step 6 does not yet carry
the Dewey and `/uf.review-council` references — the sibling edit
(#549) adds them to Step 6. The test is therefore legitimately
red now (Step 6 lacks those refs for the in-scope files) and
turns green once #549 extends Step 6, exactly matching ADR-002's
declared edit target.

**D2 — Table-driven, per-file/per-reference structure.**
A `[]struct{ file string; mustContain []string; mustNotContain
[]string }` table keyed by speckit filename, restricted to the
five in-scope upstream files. Each entry runs as a
`t.Run(file, ...)` subtest so failures are isolated and name the
file; each missing/leaked reference is reported with `t.Errorf`
naming the exact reference. This mirrors the existing
`mustContain`/`mustNotContain` guardrail pattern. The table MUST
be an explicit enumerated list of the five in-scope filenames,
NOT derived from a `speckit.*.md` filesystem glob, so the scope
cannot drift as files are added and the required-reference matrix
(D4) stays authoritative.

**D3 — Reference matching scoped to instruction context.**
To avoid green-passing on drifted-but-substring-present content,
matching scopes to the relevant Step 6 injected block for each
speckit file (as the guardrail test scopes to a fenced/labeled
block) rather than a naive whole-file `strings.Contains`. Dewey
is satisfied by presence of `dewey_semantic_search` OR
`dewey_search`.

Block granularity (coordination with #549): the current Step 6
"Spec-phase guardrails block" is a single labeled/fenced block
covering multiple files, so per-file assertions are not yet
extractable. #549 MUST restructure Step 6 so the required
references are injected at a granularity the test can resolve
per in-scope file — either (a) per-file labeled/fenced blocks,
or (b) per-in-scope-command reference sub-blocks that carry the
file name — so the "names the exact missing reference and the
exact template file" failure output is satisfiable. The test
author extracts by the label/fence #549 introduces; until #549
lands, the extraction target for the newly-required Dewey and
`/uf.review-council` references does not exist, which is the
concrete driver of the red-first state. This block shape is
best finalized jointly with #549's Step 6 edit.

**D4 — Required-reference matrix (per ADR-002).**
Directly transcribed from the ADR-002 reference matrix:

| Command | Dewey | Constitution awareness | `/uf.review-council` |
|---------|-------|------------------------|----------------------|
| `speckit.specify.md` | required | required | — |
| `speckit.plan.md` | required | required | — |
| `speckit.tasks.md` | required | required | — |
| `speckit.implement.md` | required | required | required |
| `speckit.constitution.md` | required | required | — |

`/uf.review-council` is `mustContain` for `speckit.implement.md`
only and `mustNotContain` for the other four in-scope files,
locking the implement-only boundary. `speckit.plan.md` is the
compliant exemplar defining the target reference shape.

**D5 — Red-first coordination.**
The test is expected to fail until #549 lands. To avoid leaving
`main` red, the test MUST merge together with (or immediately
before, in the same PR chain as) the #549 reconciliation. This
is a sequencing constraint recorded for `/uf.finale`, not a code
concern.

## Risks / Trade-offs

- **Substring brittleness (MEDIUM).** Naive `Contains` can green-
  pass on semantically-broken content. Mitigation: D3 scopes
  matching to the Step 6 injected block context.
- **False failures from uniform assertions (LOW).** A blanket
  "all refs in all files" table would false-red where a ref is
  intentionally absent (e.g., `/uf.review-council` outside
  implement). Mitigation: D4 conditional per-file matrix drawn
  directly from ADR-002.
- **Red-build window (LOW).** If merged before #549, `main`
  breaks. Mitigation: D5 coordination constraint.
- **Trade-off accepted:** scoping to the Step 6 injection-
  instruction surface (not a full regenerate-cycle e2e) keeps the
  test fast, isolated, and dependency-free at the cost of not
  exercising the literal `specify init` → `/uf.init` path end to
  end. Per ADR-002, the regenerate-cycle regression test is a
  separate acceptance criterion on #549. The content assertion
  still fails if the Step 6 source-of-truth surface drifts, which
  is the primary regression vector.
<!-- scaffolded by uf vdev -->
