## Why

The OSS Scout hero is currently named "Pinkman" across
all spec artifacts, agent files, scaffold assets, Dewey
tags, data directories, and documentation. The name needs
to change to "Snoopy" to better reflect the hero's
investigative scouting persona.

This is a mechanical rename with no behavioral changes.
The agent's four modes (discover, trend, audit, report),
its `/scout` command routing, its Dewey learning
conventions, and its license compatibility classification
pipeline all remain functionally identical. Only the
identity string changes.

Because Pinkman exists exclusively on unmerged feature
branches (`032-pinkman-oss-scout`, `032-pinkman-clean`,
`opsx/pinkman-dewey-enrichment`,
`opsx/pinkman-license-compatibility`), the rename must
be applied to the working branch that will be merged to
`main`. The `main` branch has zero Pinkman references
today.

## What Changes

Rename all occurrences of "Pinkman" / "pinkman" to
"Snoopy" / "snoopy" across agent files, scaffold assets,
spec artifacts, OpenSpec changes, slash commands,
documentation, Dewey tags, data paths, Go test
assertions, and `.gitignore` patterns.

## Capabilities

### New Capabilities

- None. This is a rename, not a feature change.

### Modified Capabilities

- `OSS Scout identity`: Agent identity changes from
  "Pinkman" to "Snoopy" across all touchpoints (agent
  file, scaffold asset, slash command routing, Dewey
  tags, artifact envelope producer field, data
  directory paths).

### Removed Capabilities

- None.

## Impact

~405 text references across ~25 files on the Pinkman
feature branches. Affected categories:

1. **File/directory renames** (3 renames):
   - `.opencode/agents/pinkman.md` -> `snoopy.md`
   - `internal/scaffold/assets/opencode/agents/pinkman.md`
     -> `snoopy.md`
   - `.uf/pinkman/` -> `.uf/snoopy/` (data directory in
     `.gitignore` and agent file references)

2. **Agent identity text** (~28 refs): Role name, persona
   description in agent file (live + scaffold copy).

3. **Slash command routing** (~18 refs):
   `subagent_type: "pinkman"` -> `"snoopy"` in
   `scout.md` (live + scaffold copy).

4. **Data/reports directory paths** (~8 refs):
   `.uf/pinkman/reports/` -> `.uf/snoopy/reports/`.

5. **Dewey tags** (~30+ refs): `pinkman-discover`,
   `pinkman-trend`, `pinkman-audit`, `pinkman-report`
   -> `snoopy-discover`, `snoopy-trend`, `snoopy-audit`,
   `snoopy-report`.

6. **Artifact envelope producer field** (~2 refs):
   `producer: pinkman` -> `producer: snoopy`.

7. **Documentation** (~23 refs): `AGENTS.md`,
   `unbound-force.md`.

8. **Go test assertions** (~5 refs):
   `scaffold_test.go` (asset paths, `isToolOwned`),
   `main_test.go` (file count comment).

9. **Spec artifacts** (~175+ refs): All prose across
   `specs/032-pinkman-oss-scout/` (8 files) and
   `openspec/changes/pinkman-*/` (10 files).

10. **`.gitignore`** (~3 refs): Comment and path
    patterns.

11. **Git branch names** (4 branches): Historical --
    cannot and should not be renamed. New branches will
    use "snoopy".

12. **Dewey stored learnings** (~10 entries): Tagged
    with `pinkman-*` -- these are historical records
    and should not be modified.

**Not affected**: No Go business logic changes. No
behavioral changes. No schema changes. No CI changes.
No new dependencies.

## Constitution Alignment

Assessed against the Unbound Force org constitution.

### I. Autonomous Collaboration

**Assessment**: N/A

This change does not alter artifact-based communication
patterns. The artifact envelope `producer` field value
changes from `pinkman` to `snoopy`, but the envelope
format and self-describing output structure remain
identical.

### II. Composability First

**Assessment**: N/A

No dependency changes. The OSS Scout remains
independently installable via `uf init`. The rename is
confined to identity strings with no impact on
composability or integration points.

### III. Observable Quality

**Assessment**: PASS

Dewey learning tags change from `pinkman-*` to
`snoopy-*`, maintaining the structured tag convention.
Machine-parseable output format is unchanged. Provenance
metadata (artifact envelope producer field) is updated
consistently.

### IV. Testability

**Assessment**: PASS

Go test assertions (`expectedAssetPaths`, `isToolOwned`
test cases) will be updated to reference `snoopy.md`
instead of `pinkman.md`. Test isolation is unaffected.
The scaffold drift detection test will verify the renamed
asset is correctly embedded.
