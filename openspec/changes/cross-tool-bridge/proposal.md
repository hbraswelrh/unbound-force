## Why

`uf init` deploys convention packs and agent files under
`.opencode/`, which only OpenCode auto-discovers. Teams
where contributors use Claude Code or Cursor get no
benefit from convention packs unless they manually find
and reference the files. There is no bridge telling
non-OpenCode tools where the packs live.

The current `agent-agnostic-file-standard` change
proposes moving packs to `.agents/packs/` -- but no tool
auto-discovers that path either. The move adds migration
cost without solving the discovery problem. The actual
need is thin bridge files in each tool's native discovery
location pointing to the canonical OpenCode files.

Additionally, AGENTS.md lacks a "Convention Packs"
section that would tell any LLM -- regardless of tool --
where packs are and how to use them.

## What Changes

### 1. AGENTS.md "Convention Packs" section

`uf init` appends a "Convention Packs" section to
AGENTS.md (if it exists and lacks one). The section
lists deployed packs by path and instructs agents to
read them before writing or reviewing code.

Detection: marker heading `## Convention Packs`. If
present, skip. Same idempotent pattern as
`ensureGitignore()`.

### 2. CLAUDE.md bridge file

`uf init` creates or appends to `CLAUDE.md` a managed
block that imports AGENTS.md and each deployed convention
pack using Claude Code's `@path` import syntax.

**Generated content (example for a Go project):**

```
# Unbound Force — managed by uf init

@AGENTS.md

## Convention Packs

@.opencode/uf/packs/default.md
@.opencode/uf/packs/severity.md
@.opencode/uf/packs/go.md
```

The pack list is generated dynamically -- `uf init`
enumerates which packs were actually deployed (based on
`--lang` and `detectLang()`) and lists only those. This
avoids broken `@imports` for absent packs.

Detection: marker comment `# Unbound Force — managed by
uf init`. If present, skip. If CLAUDE.md exists without
the marker, append the block (preserving existing user
content above).

### 3. .cursorrules bridge file

`uf init` creates or appends to `.cursorrules` a managed
block that instructs Cursor's agent to read AGENTS.md
and the applicable convention packs.

**Generated content (example for a Go project):**

```
# Unbound Force — managed by uf init

This project follows coding conventions defined in
AGENTS.md and enforced through convention packs. Before
writing or reviewing code, read the applicable convention
pack(s) from .opencode/uf/packs/ and apply all rules
marked [MUST].

Available packs:
- .opencode/uf/packs/default.md (language-agnostic)
- .opencode/uf/packs/severity.md (severity definitions)
- .opencode/uf/packs/go.md (Go-specific)
```

Detection: same marker pattern. Idempotent.

### 4. Convention packs stay at .opencode/uf/packs/

No file move. Packs remain at their current canonical
location. The bridge files point to them. This avoids
migration cost for existing adopters and keeps uf
opinionated about OpenCode as the primary tool.

## Design Rationale

### Why not move packs to .agents/packs/?

`.agents/` is not auto-discovered by any tool (OpenCode,
Claude Code, or Cursor). Moving packs there adds
migration cost without improving discovery. The bridge
approach achieves the same cross-tool visibility without
moving files.

### Why not symlinks?

Claude Code supports symlinks in `.claude/rules/`, but
`@import` in CLAUDE.md is simpler and has no cross-
platform issues. Cursor requires `.mdc` files with
different frontmatter -- symlinks to `.md` files would
not be recognized. Symlinks also have Windows
compatibility concerns (`core.symlinks`, admin
privileges).

### Why bridges instead of native format conversion?

Convention packs are pure Markdown content with numbered
MUST/SHOULD rules. Any LLM can read and follow them
regardless of the tool that loads them. The UF-specific
frontmatter (`pack_id`, `language`, `version`) is
harmless metadata that non-OpenCode tools ignore. No
format conversion is needed -- only discovery.

### Stacking behavior

All three tools natively stack project-level and user-
level config files. A contributor using Claude Code gets
both their personal `~/.claude/CLAUDE.md` AND the
project's committed `CLAUDE.md`. The bridge files add
project conventions without overriding personal settings.

### Cursor bridge limitation

The `.cursorrules` bridge is weaker than the `CLAUDE.md`
bridge by design. Claude Code's `@path` syntax auto-
loads referenced files into context at session start.
Cursor loads `.cursorrules` into context but the agent
must then choose to read the referenced pack files --
auto-loading via `@import` is not supported. This means
Cursor users get the instruction to read packs, but
the packs are not guaranteed to be in context for every
interaction. This is a known limitation of Cursor's
rule system and cannot be solved without Cursor adding
an import mechanism.

## Capabilities

### New Capabilities

- `ensureCLAUDEmd()`: Creates or appends managed block
  to CLAUDE.md with @imports for AGENTS.md and deployed
  convention packs. Idempotent via marker detection.
- `ensureCursorrules()`: Creates or appends managed block
  to .cursorrules with pack reference instructions.
  Idempotent via marker detection.
- `ensureAGENTSmdPackSection()`: Appends Convention Packs
  section to AGENTS.md listing deployed packs.
  Idempotent via heading detection.

### Modified Capabilities

- `Run()` in `scaffold.go`: Calls the three new ensure
  functions after existing `ensureGitignore()`, before
  sub-tool delegation.
- `printSummary()`: Reports bridge file status (created,
  skipped, appended).

### Removed Capabilities

None.

## Impact

- No breaking changes. Existing `uf init` behavior is
  unchanged for OpenCode users.
- New files (CLAUDE.md, .cursorrules) are only created
  if absent. Existing files get a managed block appended.
- Bridge files should be committed to version control
  by maintainers so new contributors get convention
  packs out-of-box regardless of their tool choice.
- Supersedes `agent-agnostic-file-standard` (pack move
  abandoned) and `multi-agent-init` (bridge approach
  replaces --packs-only flag).

## Supersedes

- **agent-agnostic-file-standard**: The `.agents/packs/`
  move is abandoned. AGENTS.md deduplication (removing
  governance rules restated from the constitution) should
  be extracted as a separate, independent change.
- **multi-agent-init**: The `--packs-only` flag and
  AGENTS.md pack section are replaced by the bridge
  approach. The "Convention Packs" section concept
  migrates here.

## Constitution Alignment

Assessed against the Unbound Force org constitution.

### I. Autonomous Collaboration

**Assessment**: PASS

Convention packs remain self-describing artifacts.
Bridge files are thin pointers -- no runtime coupling
between tools. Each tool discovers packs independently
through its native mechanism.

### II. Composability First

**Assessment**: PASS

This change directly serves composability. Convention
packs become usable by Claude Code and Cursor without
requiring OpenCode. Each tool benefits independently.
The bridge files are optional -- removing them does not
break OpenCode functionality.

### III. Observable Quality

**Assessment**: N/A

No change to output formats or quality metrics.

### IV. Testability

**Assessment**: PASS

Each ensure function follows the established
ensureGitignore() pattern with injectable I/O. Tests
verify idempotency, marker detection, dynamic pack
enumeration, and append-without-overwrite behavior.
