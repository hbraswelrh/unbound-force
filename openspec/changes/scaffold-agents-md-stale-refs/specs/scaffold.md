## ADDED Requirements

### Requirement: Stale command reference detection in root AGENTS.md

`warnStaleCommandRefs()` MUST include the root `AGENTS.md`
file in its scan set, in addition to `.opencode/agents/*.md`.
When root `AGENTS.md` contains a stale (pre-`uf.`-namespace)
command reference, `uf init` MUST emit a warning identifying
the file, the stale reference, and its migrated replacement.
The scan MUST remain warn-only: `warnStaleCommandRefs()` MUST
NOT modify `AGENTS.md` or any user-owned file, honoring the
ownership contract enforced by `isToolOwned()`.

#### Scenario: Stale reference present in AGENTS.md

- **GIVEN** a target directory whose root `AGENTS.md` contains
  the stale reference `/review-council`
- **WHEN** `warnStaleCommandRefs()` runs during `uf init`
- **THEN** a warning is written to the output writer naming
  `AGENTS.md`, the stale ref `/review-council`, and its
  replacement `/uf.review-council`
- **AND** the contents of `AGENTS.md` remain unmodified

#### Scenario: No stale reference in AGENTS.md

- **GIVEN** a target directory whose root `AGENTS.md` contains
  only migrated references such as `/uf.review-council`
- **WHEN** `warnStaleCommandRefs()` runs
- **THEN** no warning is written for `AGENTS.md`

#### Scenario: AGENTS.md absent or unreadable

- **GIVEN** a target directory with no root `AGENTS.md`, or an
  `AGENTS.md` that cannot be read
- **WHEN** `warnStaleCommandRefs()` runs
- **THEN** the read error (if any) is logged and the scan
  continues without aborting `uf init`
- **AND** no false warning is emitted for the missing file

### Requirement: Word-boundary-anchored reference matching

Stale command reference matching MUST be word-boundary
(token) anchored rather than substring-based. A migrated
reference that contains a stale reference as a substring
(for example `/uf.review-council` contains `/review-council`)
MUST NOT be flagged as stale.

#### Scenario: Migrated reference not flagged as substring match

- **GIVEN** file content containing `/uf.review-council` and no
  standalone `/review-council` token
- **WHEN** stale-reference matching runs
- **THEN** `/review-council` is NOT reported as a stale
  reference

#### Scenario: Standalone stale token is flagged

- **GIVEN** file content containing a standalone
  `/review-council` token not preceded by the `uf.` prefix
- **WHEN** stale-reference matching runs
- **THEN** `/review-council` IS reported as a stale reference
  with replacement `/uf.review-council`

## MODIFIED Requirements

### Requirement: warnStaleCommandRefs scan scope and matching

`warnStaleCommandRefs()` scans for stale command references
across `.opencode/agents/*.md` AND the root `AGENTS.md`, using
the `renamedCommands`-derived `refMap` as the single source of
truth, and matches references using word-boundary anchoring.

Previously: `warnStaleCommandRefs()` scanned only
`.opencode/agents/*.md` and matched references using
`strings.Contains` (substring), which never inspected the
user-owned root `AGENTS.md` and could falsely flag a migrated
reference as stale.

## REMOVED Requirements

_None._
<!-- scaffolded by uf vdev -->
