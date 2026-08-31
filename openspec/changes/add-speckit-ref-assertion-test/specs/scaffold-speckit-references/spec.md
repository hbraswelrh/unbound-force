# Speckit Template Reference Assertions

## ADDED Requirements

### Requirement: Speckit templates MUST assert required UF references

The scaffold test suite MUST include a table-driven test that
asserts, for each in-scope upstream `speckit.*.md` command
template, the presence of its required Unbound Force references.
The in-scope files are the five upstream templates that
`specify init` regenerates, as pinned by ADR-002
(`docs/decisions/002-uf-refs-source-of-truth-layer.md`):
`speckit.specify.md`, `speckit.plan.md`, `speckit.tasks.md`,
`speckit.implement.md`, and `speckit.constitution.md`. The four
UF-custom commands (`speckit.analyze.md`, `speckit.checklist.md`,
`speckit.clarify.md`, `speckit.taskstoissues.md`) are out of
scope.

Required references comprise Dewey knowledge-retrieval hooks
(`dewey_semantic_search` and/or `dewey_search`), constitution
awareness (`.specify/memory/constitution.md`), and the
`/uf.review-council` gate in the implement phase only.

The required-reference set MUST be encoded per file so that a
reference legitimately absent from one template does not cause a
false failure on another. The test MUST assert against the
durable source-of-truth injection layer pinned by ADR-002 —
Step 6 of the embedded canonical asset
`internal/scaffold/assets/opencode/commands/uf.init.md`, read via
`assetContent` — not a working-tree-only copy, so the guarantee
holds across downstream `/uf.init` runs.

The test MUST NOT modify `knownNonEmbeddedFiles`, embedding
expectations, coverage thresholds, or any other quality or
governance gate. The test MUST use only the standard library
`testing` and `strings` packages, consistent with existing
scaffold test conventions.

#### Scenario: Required reference present

- **GIVEN** an in-scope `speckit.*.md` template whose Step 6
  injection surface contains all of its required references
- **WHEN** the content-assertion test runs
- **THEN** the per-file subtest for that template passes

#### Scenario: Required reference missing

- **GIVEN** an in-scope `speckit.*.md` template whose Step 6
  injection surface is missing one or more of its required
  references
- **WHEN** the content-assertion test runs
- **THEN** the test fails and the failure output names the exact
  missing reference and the exact template file

### Requirement: review-council reference MUST be implement-scoped

The test MUST assert that the `/uf.review-council` reference is
present for `speckit.implement.md` and MUST assert that it is
absent from the other four in-scope `speckit.*.md` templates,
locking the implement-only boundary.

#### Scenario: review-council correctly scoped

- **GIVEN** the Step 6 surface carries `/uf.review-council` for
  `speckit.implement.md` and for no other in-scope template
- **WHEN** the content-assertion test runs
- **THEN** the review-council scope subtests pass

#### Scenario: review-council leaks outside implement

- **GIVEN** a non-implement in-scope `speckit.*.md` template
  whose Step 6 surface contains `/uf.review-council`
- **WHEN** the content-assertion test runs
- **THEN** the test fails and names the offending template

### Requirement: test MUST be a red-first regression gate

The test MUST fail against the current Step 6 injection surface
(in which the required Dewey and `/uf.review-council` references
are not yet injected for the in-scope upstream files) and MUST
pass once the sibling reconciliation edit (issue #549) extends
Step 6 to inject them. The test MUST be introduced in
coordination with that reconciliation so the main branch is never
left with a persistently failing build.

#### Scenario: red before reconciliation

- **GIVEN** the Step 6 surface in its pre-reconciliation state
- **WHEN** the content-assertion test runs
- **THEN** the test fails red, naming the missing references

#### Scenario: green after reconciliation

- **GIVEN** the Step 6 surface after the #549 reconciliation edit
- **WHEN** the content-assertion test runs
- **THEN** the test passes green

## MODIFIED Requirements

None.

## REMOVED Requirements

None.
<!-- scaffolded by uf vdev -->
