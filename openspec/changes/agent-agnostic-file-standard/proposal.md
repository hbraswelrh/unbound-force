## Why

AGENTS.md currently restates governance rules already
defined in the org constitution
(`.specify/memory/constitution.md`). Duplicated sections
include CI parity requirements, review council
prerequisites, and spec-first development workflow rules.
This creates two places to maintain and risks divergence
when the constitution is amended.

AGENTS.md should reference the constitution for governance
rules and provide only repo-specific operational details:
project structure, build commands, technology choices, and
pipeline command reference.

## Scope Reduction Notice

This change was originally scoped to also move convention
packs from `.opencode/uf/packs/` to `.agents/packs/`.
After analysis, the pack move was abandoned because:

- `.agents/` is not auto-discovered by any tool (OpenCode,
  Claude Code, or Cursor). The move adds migration cost
  without improving discovery.
- Cross-tool visibility is better solved by bridge files
  in each tool's native discovery path. See the
  `cross-tool-bridge` change for that approach.

This change now covers only the AGENTS.md governance
deduplication. Convention packs remain at
`.opencode/uf/packs/`.

## What Changes

### 1. Remove duplicated governance sections from AGENTS.md

Remove sections that restate rules already in the
constitution's Development Workflow and Core Principles:

- **CI Parity Gate**: Restates CI requirements already
  covered by the constitution's "Continuous Integration"
  rule and the `/unleash` command's CI derivation logic.
- **Review Council as PR Prerequisite**: Restates code
  review requirements already covered by the
  constitution's "Code Review" rule and the
  `/review-council` command documentation.
- **Spec-First Development**: Restates the spec-driven
  development workflow already covered by the
  constitution's "Spec-Driven Development" rule and the
  Specification Framework section in AGENTS.md.

### 2. Add constitution reference

Replace removed sections with a concise reference
directing agents to the constitution for governance
rules.

## Capabilities

### New Capabilities

None.

### Modified Capabilities

- `AGENTS.md`: Deduplicated -- references constitution
  instead of restating governance rules. Behavioral
  Constraints section retained (gatekeeping, phase
  boundaries) as these are repo-specific elaborations,
  not duplications.

### Removed Capabilities

None. All governance rules remain in effect via the
constitution. No behavioral change for agents.

## Impact

- No code changes. Markdown-only.
- No migration needed for consuming repositories.
- Agents that read AGENTS.md continue to get governance
  rules via the constitution reference and the
  constitution file itself (already loaded by OpenCode
  via `.specify/memory/constitution.md`).
- Convention pack paths are NOT changed by this
  proposal. Packs remain at `.opencode/uf/packs/`.

## Constitution Alignment

Assessed against the Unbound Force org constitution.

### I. Autonomous Collaboration

**Assessment**: N/A

No change to artifact communication.

### II. Composability First

**Assessment**: N/A

No change to hero independence or integration.

### III. Observable Quality

**Assessment**: N/A

No change to output formats or provenance.

### IV. Testability

**Assessment**: N/A

No code changes. Documentation only.
