## Tasks

### Part 1: Revert convention pack move (scope reduction)

The pack move to `.agents/packs/` was abandoned. These
tasks revert Part 1 changes from the original scope.

#### Scaffold engine (Go code)
- [x] Revert pack files from `internal/scaffold/assets/agents/packs/` back to `internal/scaffold/assets/opencode/uf/packs/`
- [x] Remove `"agents/"` from `knownAssetPrefixes` in scaffold.go
- [x] Remove `agents/` → `.agents/` mapping from `mapAssetPath`
- [x] Revert `isConventionPack` prefix from `agents/packs/` back to `opencode/uf/packs/`
- [x] Revert `scaffold_test.go` -- restore original expected asset paths and test assertions
- [x] Revert `internal/doctor/checks.go` -- pack directory check back to `.opencode/uf/packs/`
- [x] Revert `internal/doctor/doctor_test.go` -- pack directory test fixtures
- [x] Revert `internal/schemas/packvalidator_test.go` -- pack file path references

#### Scaffold asset references (Markdown files deployed to other repos)
- [x] Revert all Divisor agent files in `internal/scaffold/assets/opencode/agents/` -- restore `.opencode/uf/packs/` references
- [x] Revert cobalt-crush-dev.md asset -- pack path references
- [x] Revert command files in `internal/scaffold/assets/opencode/command/` -- pack path references

#### Live repo files (unbound-force's own copies)
- [x] Revert all Divisor agent files in `.opencode/agents/` -- restore `.opencode/uf/packs/` references
- [x] Revert all other agent files in `.opencode/agents/` -- restore pack path references
- [x] Revert command files in `.opencode/command/` -- restore pack path references
- [x] Revert skill files in `.opencode/skills/` -- restore pack path references
- [x] Move live pack files back from `.agents/packs/` to `.opencode/uf/packs/`

#### AGENTS.md pack path references
- [x] Revert project structure section in AGENTS.md to show `.opencode/uf/packs/` path
- [x] Revert any Recent Changes / Active Technologies entries that reference `.agents/packs/`

### Part 2: Clean AGENTS.md / constitution overlap

These tasks are retained from the original scope.

- [x] Remove duplicated governance rules from AGENTS.md (CI Parity Gate, Review Council as PR Prerequisite, Spec-First Development) and replace with constitution reference
- [x] Verify AGENTS.md still contains repo-specific operational details (project structure, build commands, technology choices, pipeline commands) and does not remove non-duplicated content

### Verification
- [x] Run `make test` or `go test ./...` to verify all tests pass
- [x] Run `make lint` or lint commands to verify no lint issues
