## Tasks

### Part 1: Revert convention pack move (scope reduction)

The pack move to `.agents/packs/` was abandoned. These
tasks revert Part 1 changes from the original scope.

#### Scaffold engine (Go code)
- [ ] Revert pack files from `internal/scaffold/assets/agents/packs/` back to `internal/scaffold/assets/opencode/uf/packs/`
- [ ] Remove `"agents/"` from `knownAssetPrefixes` in scaffold.go
- [ ] Remove `agents/` → `.agents/` mapping from `mapAssetPath`
- [ ] Revert `isConventionPack` prefix from `agents/packs/` back to `opencode/uf/packs/`
- [ ] Revert `scaffold_test.go` -- restore original expected asset paths and test assertions
- [ ] Revert `internal/doctor/checks.go` -- pack directory check back to `.opencode/uf/packs/`
- [ ] Revert `internal/doctor/doctor_test.go` -- pack directory test fixtures
- [ ] Revert `internal/schemas/packvalidator_test.go` -- pack file path references

#### Scaffold asset references (Markdown files deployed to other repos)
- [ ] Revert all Divisor agent files in `internal/scaffold/assets/opencode/agents/` -- restore `.opencode/uf/packs/` references
- [ ] Revert cobalt-crush-dev.md asset -- pack path references
- [ ] Revert command files in `internal/scaffold/assets/opencode/command/` -- pack path references

#### Live repo files (unbound-force's own copies)
- [ ] Revert all Divisor agent files in `.opencode/agents/` -- restore `.opencode/uf/packs/` references
- [ ] Revert all other agent files in `.opencode/agents/` -- restore pack path references
- [ ] Revert command files in `.opencode/command/` -- restore pack path references
- [ ] Revert skill files in `.opencode/skills/` -- restore pack path references
- [ ] Move live pack files back from `.agents/packs/` to `.opencode/uf/packs/`

#### AGENTS.md pack path references
- [ ] Revert project structure section in AGENTS.md to show `.opencode/uf/packs/` path
- [ ] Revert any Recent Changes / Active Technologies entries that reference `.agents/packs/`

### Part 2: Clean AGENTS.md / constitution overlap

These tasks are retained from the original scope.

- [x] Remove duplicated governance rules from AGENTS.md (CI Parity Gate, Review Council as PR Prerequisite, Spec-First Development) and replace with constitution reference
- [ ] Verify AGENTS.md still contains repo-specific operational details (project structure, build commands, technology choices, pipeline commands) and does not remove non-duplicated content

### Verification
- [ ] Run `make test` or `go test ./...` to verify all tests pass
- [ ] Run `make lint` or lint commands to verify no lint issues
