# Research: Agent Customization Onboarding

## R1: Profile Storage Location and Format

**Decision**: Store the user profile as a Markdown file
with YAML frontmatter at `.uf/onboarding/profile.md`.
Profile history snapshots stored at
`.uf/onboarding/history/YYYY-MM-DDTHH-MM-SS.md`.

**Rationale**: This follows the established pattern used
by Muti-Mind's backlog items (`.uf/muti-mind/backlog/
BI-NNN.md` per Spec 004) and Mx F's impediments
(`.uf/mx-f/impediments/IMP-NNN.md` per Spec 007). YAML
frontmatter provides machine-parseable metadata while the
Markdown body holds the human-readable narrative content
for each dimension (inspiration, interests, objectives).
The `.uf/` directory convention (per Spec 025) is the
standard location for per-project runtime data.

**Alternatives considered**:
- JSON file: Less human-readable for narrative content.
  Inspiration and interests are inherently freeform text
  that benefits from Markdown formatting.
- YAML-only file: Would work but loses the ability to
  include rich narrative descriptions with headings,
  lists, and emphasis.
- SQLite/database: Overkill for a single user profile
  per project. Introduces unnecessary complexity.

## R2: Profile Injection Mechanism

**Decision**: Inject the user profile into hero context
through AGENTS.md instructions. Add a "User Profile"
section to AGENTS.md that directs all agents to read
`.uf/onboarding/profile.md` if it exists and use its
content to contextualize their decisions. Each hero's
existing agent file already reads AGENTS.md as part of
its "Source Documents" step.

**Rationale**: This is the least-invasive approach. All
hero agents already read AGENTS.md as their primary
context source (verified in Cobalt-Crush's "Source
Documents" section, Muti-Mind's directives, and all
Divisor persona files). By adding profile-reading
instructions to AGENTS.md, all heroes automatically
gain profile awareness without any changes to their
individual agent files. This preserves FR-003 (inject
without modifying hero agent files) and FR-006 (backward
compatible -- instruction to read a file that doesn't
exist is a no-op).

The `/uf-init` command already handles AGENTS.md guidance
injection (per Spec 030). A new guidance block for the
user profile section fits naturally into this pattern.

**Alternatives considered**:
- Modifying each hero's agent file: Violates FR-003 and
  requires touching 10+ files. Creates maintenance burden.
- Environment variable injection: Not persistent, not
  project-scoped, hard to version control.
- MCP tool: Overkill; the profile is a single file read.
- Separate context file included via OpenCode config:
  OpenCode doesn't support arbitrary file inclusion
  beyond AGENTS.md and agent files.

## R3: Onboarding Agent Design

**Decision**: Create the onboarding agent as a standard
OpenCode agent file (`onboarding.md`) with a `/onboard`
slash command for invocation. The agent is **user-owned**
(not tool-owned) per the scaffold engine's ownership
model, allowing users to customize the interview flow.

**Rationale**: This follows the exact pattern of existing
hero agents like `cobalt-crush-dev.md` (per Spec 006) and
`mx-f-coach.md` (per Spec 007). User-owned files are
skipped if they already exist during `uf init`, so user
customizations are preserved across scaffold updates.

The agent uses temperature 0.4 (same as Cobalt-Crush)
because it needs creativity for interview conversation
while maintaining consistency in profile generation.
Tools: `read: true` (to read existing profile and agent
files for capability mapping), `write: true` (to persist
profile), `edit: true` (to update existing profiles).
No `bash` access needed.

**Alternatives considered**:
- Tool-owned agent: Would prevent users from customizing
  the interview questions, tone, or flow. Since the
  onboarding experience is inherently personal, user
  ownership is the right choice.
- CLI command (Go binary): The interview is a
  conversational AI interaction, not a deterministic
  operation. A Go CLI would need to shell out to an AI
  model anyway. The agent file pattern is purpose-built
  for this.

## R4: Hero Capability Map

**Decision**: The hero capability map is defined inline
within the onboarding agent file as a static reference
table. It is NOT generated dynamically from agent files.

**Rationale**: The five heroes and their capabilities are
well-known and change infrequently (hero additions are
architectural decisions per Specs 004-007). A static map
in the agent file is simpler, faster, and more reliable
than parsing agent files at runtime. The agent can always
be updated when a new hero is added (user-owned file).

The map structure covers:
- Hero name
- Agent file path
- Core capabilities (1-3 bullet points)
- Example user objectives that map to this hero

This is sufficient for FR-009 (map objectives to heroes)
and FR-004 (route users to existing heroes).

**Alternatives considered**:
- Dynamic discovery via `fs.WalkDir` on agent files: 
  Fragile (depends on file naming conventions), slower,
  and requires bash access. Agent file content would need
  to be parsed to extract capabilities, which is error-
  prone.
- Separate capability manifest file: Adds another file
  to maintain without significant benefit. The agent
  already has the context it needs.

## R5: Multi-User Profiles

**Decision**: Per-user profiles are identified by a
`user` field in the YAML frontmatter. The default
profile path is `.uf/onboarding/profile.md` for the
primary user. For multi-user setups, the profile can be
parameterized by username, but v1 targets single-user.

**Rationale**: The spec states "each user's profile is
independent and applies only to their own sessions."
For v1, the simplest approach is one profile per project.
The YAML frontmatter includes a `user` field for future
multi-user expansion, but the reading mechanism in
AGENTS.md points to a single path. Multi-user routing
(e.g., by `$USER` environment variable) can be added in
a later iteration.

**Alternatives considered**:
- Immediate multi-user with `$USER` routing: Adds
  complexity without immediate value. Most projects start
  with a single developer interacting with the AI swarm.
- Global profile (outside project): Violates the per-
  project assumption in the spec and prevents different
  projects from having different objectives.

## R6: Scaffold Integration

**Decision**: Add two new embedded assets to the scaffold
engine:
1. `opencode/agents/onboarding.md` — user-owned
2. `opencode/command/onboard.md` — tool-owned

Update `expectedAssetPaths` in tests (33 → 35).
No new `initSubTools()` delegation needed -- the files
are deployed by the standard asset walk.

**Rationale**: This follows the exact pattern used by
every prior agent and command addition (Specs 005, 006,
007, 018, 019, 026). The scaffold engine's `fs.WalkDir`
automatically picks up new files under
`internal/scaffold/assets/`. Tool ownership for the
command file ensures the invocation instructions stay
canonical. User ownership for the agent file allows
interview customization.

**Alternatives considered**:
- External tool delegation (like `specify init`): The
  onboarding agent is part of the core Unbound Force
  ecosystem, not a third-party tool. Embedding is
  correct.
- No scaffold integration (manual creation): Violates
  the zero-friction onboarding principle. Users shouldn't
  have to create files manually.

## R7: Profile Injection via /uf-init

**Decision**: Add a new AGENTS.md guidance block (Step 10
or integrated into existing Step 9) that injects a "User
Profile" section with instructions for heroes to read
`.uf/onboarding/profile.md`. This follows the Spec 030
guidance injection pattern.

**Rationale**: Spec 030 established the pattern for
injecting behavioral guidance into AGENTS.md via
`/uf-init`. The profile-reading instruction is another
behavioral guidance block: "If `.uf/onboarding/
profile.md` exists, read it and use the user's stated
inspiration, interests, and objectives to contextualize
your decisions." This block is idempotent (can be
re-injected safely) and additive (repos without profiles
are unaffected).

**Alternatives considered**:
- Manual AGENTS.md editing: Error-prone, not scalable,
  inconsistent across repos.
- OpenCode config injection: OpenCode's config doesn't
  support arbitrary context file references.
