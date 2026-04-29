## 1. Rebase onto Pinkman spec branch

Before renaming, the working branch must include all
Pinkman files from the spec branch. The rename operates
on these files.

- [x] 1.1 Identify the correct Pinkman source branch
  (`032-pinkman-clean` is the merge-ready branch with
  all consolidated changes)
- [x] 1.2 Rebase `opsx/rename-pinkman-to-snoopy` onto
  `032-pinkman-clean` to bring in all Pinkman files
- [x] 1.3 Resolve any merge conflicts from the rebase

## 2. Rename agent file and scaffold asset

File renames (git mv) for the two agent file copies.

- [x] 2.1 `git mv .opencode/agents/pinkman.md
  .opencode/agents/snoopy.md`
- [x] 2.2 `git mv
  internal/scaffold/assets/opencode/agents/pinkman.md
  internal/scaffold/assets/opencode/agents/snoopy.md`

## 3. Update agent file content

Replace all identity references within both agent file
copies (live + scaffold asset).

- [x] 3.1 Replace `Pinkman` -> `Snoopy` (title case)
  in `.opencode/agents/snoopy.md`
- [x] 3.2 Replace `pinkman` -> `snoopy` (lowercase)
  in `.opencode/agents/snoopy.md` (tags, paths,
  producer field)
- [x] 3.3 Synchronize changes to
  `internal/scaffold/assets/opencode/agents/snoopy.md`
  (scaffold asset must exactly match live copy)

## 4. Update slash command routing

Update `/scout` command files to route to `snoopy`.

- [x] 4.1 Replace `pinkman` -> `snoopy` in
  `.opencode/command/scout.md`
  (`subagent_type`, any prose references)
- [x] 4.2 Synchronize changes to
  `internal/scaffold/assets/opencode/command/scout.md`

## 5. Update `.gitignore`

- [x] 5.1 Replace `# Pinkman scouting reports` comment
  with `# Snoopy scouting reports`
- [x] 5.2 Replace `.uf/pinkman/` paths with
  `.uf/snoopy/`

## 6. Update documentation

- [x] 6.1 Replace Pinkman -> Snoopy in `AGENTS.md`
  (Utility Agents table, Project Structure tree,
  Active Technologies, Recent Changes entries)
- [x] 6.2 Replace Pinkman -> Snoopy in
  `unbound-force.md` (OSS Scout section header and
  description paragraphs)

## 7. Update Go test assertions

- [x] 7.1 Update `expectedAssetPaths` in
  `internal/scaffold/scaffold_test.go`: replace
  `"opencode/agents/pinkman.md"` with
  `"opencode/agents/snoopy.md"`
- [x] 7.2 Update `isToolOwned` test cases in
  `internal/scaffold/scaffold_test.go` to reference
  `snoopy.md`
- [x] 7.3 Update file count comment in
  `cmd/unbound-force/main_test.go` to reference Snoopy
- [x] 7.4 Add regression test
  `TestScaffoldOutput_NoPinkmanReferences` that scans
  all scaffold asset content for stale "pinkman"
  strings (per design decision R3)

## 8. Update spec artifact prose

Update hero identity references within spec files.
Directory names are preserved per design decision D1.

- [x] 8.1 Replace `Pinkman` -> `Snoopy` and
  `pinkman` -> `snoopy` in
  `specs/032-pinkman-oss-scout/spec.md`
- [x] 8.2 Replace in
  `specs/032-pinkman-oss-scout/plan.md`
- [x] 8.3 Replace in
  `specs/032-pinkman-oss-scout/tasks.md`
- [x] 8.4 Replace in
  `specs/032-pinkman-oss-scout/data-model.md`
- [x] 8.5 Replace in
  `specs/032-pinkman-oss-scout/quickstart.md`
- [x] 8.6 Replace in
  `specs/032-pinkman-oss-scout/research.md`
- [x] 8.7 Replace in
  `specs/032-pinkman-oss-scout/contracts/agent-interface.md`
- [x] 8.8 Replace in
  `specs/032-pinkman-oss-scout/checklists/requirements.md`

## 9. Update OpenSpec change artifact prose

Update hero identity references within completed
OpenSpec change files. Directory names are preserved
per design decision D1.

- [x] 9.1 Replace `Pinkman` -> `Snoopy` and
  `pinkman` -> `snoopy` in
  `openspec/changes/pinkman-dewey-enrichment/proposal.md`
- [x] 9.2 Replace in
  `openspec/changes/pinkman-dewey-enrichment/design.md`
- [x] 9.3 Replace in
  `openspec/changes/pinkman-dewey-enrichment/specs/dewey-integration.md`
- [x] 9.4 Replace in
  `openspec/changes/pinkman-dewey-enrichment/tasks.md`
- [x] 9.5 Replace in
  `openspec/changes/pinkman-license-compatibility/proposal.md`
- [x] 9.6 Replace in
  `openspec/changes/pinkman-license-compatibility/design.md`
- [x] 9.7 Replace in
  `openspec/changes/pinkman-license-compatibility/specs/license-compatibility.md`
- [x] 9.8 Replace in
  `openspec/changes/pinkman-license-compatibility/tasks.md`

## 10. Verification

- [x] 10.1 Run `go build ./...` -- verify build passes
- [x] 10.2 Run `go test -race -count=1 ./...` -- verify
  all tests pass including new regression test
- [x] 10.3 Run `grep -ri "pinkman" --include="*.go"
  --include="*.md" --include="*.yaml" .` on tracked
  files -- verify zero hits outside of:
  (a) this OpenSpec change directory
  (b) spec directory name components
  (c) OpenSpec change directory name components
  (d) historical git branch name references
- [x] 10.4 Verify constitution alignment: Observable
  Quality (Dewey tags updated to `snoopy-*`),
  Testability (regression test added and passing)
