## Context

The OSS Scout hero (Spec 032) is named "Pinkman" across
all artifacts on unmerged feature branches. The user has
requested a rename to "Snoopy." This is a purely
mechanical find-and-replace operation with no behavioral,
architectural, or dependency changes.

The rename must be applied to the working branch that
will be merged to `main`. Since `main` has zero Pinkman
references today, the rename is applied entirely within
branch-local files before merge.

Per the proposal's constitution alignment assessment:
all four principles receive PASS or N/A. The rename
preserves artifact structure (I), standalone
installability (II), machine-parseable output with
updated provenance (III), and test isolation with
updated assertions (IV).

## Goals / Non-Goals

### Goals

- Replace every occurrence of "Pinkman" / "pinkman"
  with "Snoopy" / "snoopy" in all affected files.
- Rename files and directories that contain "pinkman"
  in their path.
- Update Go test assertions to reference the new name.
- Update Dewey tag conventions from `pinkman-*` to
  `snoopy-*`.
- Preserve the `/scout` command name unchanged (it is
  role-based, not identity-based).
- Update `.gitignore` data directory paths.

### Non-Goals

- Renaming historical git branches (`032-pinkman-*`,
  `opsx/pinkman-*`). These are historical records.
- Modifying existing Dewey stored learnings. The
  `pinkman-*` tagged learnings are historical records
  with correct provenance at time of creation.
- Renaming the spec directory `specs/032-pinkman-oss-scout/`.
  Spec directories are numbered and named at creation
  time; renaming would break cross-references in
  completed spec artifacts and Dewey indexes.
- Renaming OpenSpec change directories
  (`openspec/changes/pinkman-*/`). These are completed
  archived changes with internal cross-references.
- Any behavioral changes to the agent's modes,
  prompts, or capabilities.

## Decisions

### D1: Preserve spec and OpenSpec directory names

**Decision**: Do NOT rename `specs/032-pinkman-oss-scout/`
or `openspec/changes/pinkman-*/` directories.

**Rationale**: These are historical spec artifacts. The
directory names serve as stable identifiers referenced
in Dewey indexes, git history, and AGENTS.md Recent
Changes entries. Renaming them would break these
references and violate the principle that completed
specs are immutable records. The text content within
these files SHOULD be updated to use "Snoopy" where it
refers to the current identity, but directory paths
remain unchanged.

### D2: Update content within spec files selectively

**Decision**: Update Pinkman -> Snoopy in the prose of
spec and OpenSpec files, but preserve references to
historical branch names and file paths as they were at
the time of writing.

**Rationale**: The spec files describe the hero's
identity and behavior. Updating the identity name
ensures consistency when reading specs as living
documentation. Historical references to branch names
(e.g., "Branch: 032-pinkman-oss-scout") are factual
records and should remain.

### D3: Case-sensitive replacement strategy

**Decision**: Apply two replacement patterns:
- `Pinkman` -> `Snoopy` (title case, used in prose
  and headings)
- `pinkman` -> `snoopy` (lowercase, used in file
  paths, tags, code identifiers)

**Rationale**: The codebase uses these two case variants
consistently. No UPPERCASE or mixed-case variants exist.

### D4: Dewey tag migration

**Decision**: Update tag patterns in agent file and
command file from `pinkman-{mode}` to `snoopy-{mode}`.
Do not modify existing stored learnings.

**Rationale**: New scouting operations will produce
learnings with `snoopy-*` tags. Historical learnings
retain `pinkman-*` tags as correct provenance. Semantic
search will still find all learnings regardless of tag
prefix.

## Risks / Trade-offs

### R1: Spec directory name divergence

**Risk**: The spec directory remains
`specs/032-pinkman-oss-scout/` while all content
references "Snoopy." This creates a naming inconsistency.

**Mitigation**: Accepted trade-off. The directory name
is a historical identifier, not a display name. The
spec's `title:` frontmatter and prose will use "Snoopy."
This mirrors the established pattern from Spec 013
(binary-rename) where the spec directory name preserved
the original decision context.

### R2: Dewey search fragmentation

**Risk**: Historical learnings tagged `pinkman-*` may
not surface when searching for `snoopy-*` by tag.

**Mitigation**: Dewey semantic search is content-based,
not tag-based. Searches for "OSS Scout" or "scouting
report" will find both old and new learnings. Tag-based
filtering (`dewey_semantic_search_filtered`) will only
return new learnings, which is acceptable since the
historical ones reflect a superseded identity.

### R3: Incomplete replacement on merge

**Risk**: If the rename branch is rebased onto a Pinkman
branch that has new references added after the rename
was prepared, stale "pinkman" references could survive.

**Mitigation**: Add a regression test
(`TestScaffoldOutput_NoPinkmanReferences`) that greps
all scaffold assets for "pinkman" and fails if any
remain. This catches incomplete renames at CI time.
