## Tasks

### Part 1: Implement ensureAGENTSmdPackSection()

- [x] Add `agentsmdPackMarker` constant (heading `## Convention Packs`)
- [x] Implement `ensureAGENTSmdPackSection()` that appends Convention Packs section to AGENTS.md listing deployed packs -- idempotent via heading detection, uses opts.ReadFile/WriteFile
- [x] Add `collectDeployedPacks()` helper that enumerates which pack files were deployed based on resolved language (reuse shouldDeployPack logic)

### Part 2: Implement ensureCLAUDEmd()

- [x] Add `claudemdMarker` constant (`# Unbound Force — managed by uf init`)
- [x] Implement `ensureCLAUDEmd()` that creates or appends managed block to CLAUDE.md with @imports for AGENTS.md and deployed convention packs -- idempotent via marker detection, uses opts.ReadFile/WriteFile

### Part 3: Implement ensureCursorrules()

- [x] Add `cursorrulesMarker` constant (same marker pattern)
- [x] Implement `ensureCursorrules()` that creates or appends managed block to .cursorrules with pack reference instructions -- idempotent via marker detection, uses opts.ReadFile/WriteFile

### Part 4: Wire into Run() and printSummary()

- [x] Call ensureAGENTSmdPackSection(), ensureCLAUDEmd(), ensureCursorrules() after ensureGitignore() in Run(), append results to subResults
- [x] Verify printSummary() handles new subToolResult entries (existing format should work)

### Part 5: Tests

- [x] Add TestEnsureAGENTSmdPackSection (fresh, existing without section, existing with section, idempotent)
- [x] Add TestEnsureCLAUDEmd (fresh, existing without marker, existing with marker, idempotent)
- [x] Add TestEnsureCursorrules (fresh, existing without marker, existing with marker, idempotent)
- [x] Add TestCollectDeployedPacks for different languages (go, typescript, default)

### Part 6: Verification

- [x] Run `go test -race -count=1 ./...` to verify all tests pass
- [x] Run `golangci-lint run` to verify no lint issues
