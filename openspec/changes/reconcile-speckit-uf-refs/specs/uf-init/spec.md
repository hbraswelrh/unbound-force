# Delta Spec: uf.init Step 6 Speckit guardrail references

## MODIFIED Requirements

### Requirement: Step 6 injects UF references into speckit guardrail blocks

Step 6 of `uf.init.md` (Speckit Command Guardrails) MUST inject
the Unbound Force required-reference matrix into the guardrail
blocks it applies to the five in-scope upstream speckit command
templates. The references MUST be present in the durable
injection layer (Step 6 fenced `markdown` blocks), not only in
the Step 7 read-only verifier.

Previously: the Step 6 guardrail blocks carried only
phase-boundary file-scope rules; Dewey and constitution
references were absent from the spec-phase, implement, and
constitution blocks, and `/uf.review-council` was absent
entirely. Constitution awareness existed only in the Step 7
read-only verifier.

The required-reference matrix is:

| Block (files)                         | Dewey | constitution.md | /uf.review-council |
|---------------------------------------|-------|-----------------|--------------------|
| Spec-phase (specify, plan, tasks)     | MUST  | MUST            | MUST NOT           |
| Implement (implement.md)              | MUST  | MUST            | MUST               |
| Constitution (constitution.md)        | MUST  | MUST            | MUST NOT           |

The Dewey reference requirement is satisfied by
`dewey_semantic_search` OR `dewey_search`. The constitution
reference MUST be the literal path `.specify/memory/constitution.md`.

The canonical source `.opencode/commands/uf.init.md` and the
embedded asset
`internal/scaffold/assets/opencode/commands/uf.init.md` MUST
remain byte-identical.

#### Scenario: Spec-phase block carries Dewey and constitution references

- **GIVEN** the embedded `uf.init.md` Step 6 region
- **WHEN** the "Spec-phase guardrails block" fenced content is
  extracted
- **THEN** it MUST contain `dewey_semantic_search` OR
  `dewey_search`
- **AND** it MUST contain `.specify/memory/constitution.md`
- **AND** it MUST NOT contain `/uf.review-council`

#### Scenario: Implement block carries all three references

- **GIVEN** the embedded `uf.init.md` Step 6 region
- **WHEN** the "Implement guardrails block" fenced content is
  extracted
- **THEN** it MUST contain `dewey_semantic_search` OR
  `dewey_search`
- **AND** it MUST contain `.specify/memory/constitution.md`
- **AND** it MUST contain `/uf.review-council`

#### Scenario: Constitution block carries Dewey and constitution references

- **GIVEN** the embedded `uf.init.md` Step 6 region
- **WHEN** the "Constitution guardrails block" fenced content is
  extracted
- **THEN** it MUST contain `dewey_semantic_search` OR
  `dewey_search`
- **AND** it MUST contain `.specify/memory/constitution.md`
- **AND** it MUST NOT contain `/uf.review-council`

#### Scenario: Canonical and embedded copies stay in sync

- **GIVEN** the edited `.opencode/commands/uf.init.md`
- **WHEN** the drift test `TestEmbeddedAssets_MatchSource` runs
- **THEN** the embedded asset MUST be byte-identical to the
  canonical source

#### Scenario: Red-first gate turns green

- **GIVEN** the red-first test
  `TestSpeckitTemplates_RequiredReferences` (#548) failing before
  this change
- **WHEN** the Step 6 reference injection is applied to both
  copies
- **THEN** `go test -race -count=1 ./internal/scaffold/...` MUST
  pass
- **AND** `TestGuardrailTemplates_CommandSpecificContent` MUST
  remain green (no phase-boundary leak)
