# Implementation Plan: Agent Customization Onboarding

**Branch**: `031-agent-customization-onboarding` | **Date**: 2026-04-21 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `specs/031-agent-customization-onboarding/spec.md`

## Summary

Create a conversational onboarding agent that captures
user inspiration, interests, and objectives, persists
them as a structured user profile in `.uf/onboarding/`,
and injects that profile as contextual input to all
existing heroes (Muti-Mind, Cobalt-Crush, Gaze,
The Divisor, Mx F) without modifying their core logic.
The agent maps user-described objectives to existing
hero capabilities and routes users to the correct hero
rather than creating duplicate functionality.

The implementation consists of four deliverables:
1. A new OpenCode agent file (`onboarding.md`) that
   conducts the structured interview
2. A user profile schema (Markdown + YAML frontmatter)
   persisted at `.uf/onboarding/profile.md`
3. A profile injection mechanism via AGENTS.md
   instructions that direct heroes to read the profile
4. An `/onboard` slash command to invoke the agent
5. Scaffold engine updates to deploy the new agent and
   command via `uf init`

## Technical Context

**Language/Version**: Go 1.24+ (scaffold engine),
Markdown (agent file, command file, profile schema)
**Primary Dependencies**: `embed.FS` (scaffold engine
asset embedding), existing scaffold engine at
`internal/scaffold/scaffold.go`
**Storage**: Markdown files with YAML frontmatter at
`.uf/onboarding/profile.md` and
`.uf/onboarding/history/` for profile snapshots
**Testing**: Standard library `testing` package;
scaffold drift detection tests
**Target Platform**: macOS, Linux (same as existing
`uf` CLI)
**Project Type**: CLI tool + OpenCode agent ecosystem
**Performance Goals**: Profile read/write under 100ms
(filesystem only, no network)
**Constraints**: No new Go dependencies; no
modification to existing hero agent files' core logic;
backward compatible (heroes unchanged when no profile)
**Scale/Scope**: Single user per project; profile file
< 10KB; 5 heroes to integrate with

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check
after Phase 1 design.*

### I. Autonomous Collaboration — PASS

The onboarding agent produces a user profile artifact
(Markdown + YAML frontmatter) at a well-known location
(`.uf/onboarding/profile.md`). Heroes consume this
artifact asynchronously -- they read the file when
invoked, not through runtime coupling. The profile is
self-describing: it contains the user identifier,
version, timestamp, and status metadata. If the
profile does not exist, heroes operate with default
behavior unchanged.

### II. Composability First — PASS

The onboarding agent is independently deployable via
`uf init`. It delivers value alone (capturing user
context) without requiring any hero to be present. When
heroes are present, the profile provides additive
value by customizing their behavior. No hero requires
the onboarding agent as a hard prerequisite. The
profile is consumed through a well-defined extension
point (file read at invocation time) rather than
requiring hero modification.

### III. Observable Quality — PASS

The user profile uses a structured format (YAML
frontmatter + Markdown body) that is machine-parseable.
Profile metadata includes provenance (creator agent,
version, timestamps). The hero capability mapping is a
structured data section that can be validated against
the known hero roster. Profile completeness status
(draft/complete) is a machine-readable field.

### IV. Testability — PASS

The onboarding agent is a Markdown file -- testable
through scaffold drift detection (same as all other
agents). Profile file I/O follows the established
injectable pattern (`ReadFile`/`WriteFile` on the
`Options` struct). The profile schema can be validated
by reading the YAML frontmatter. Hero profile
consumption is testable by creating a profile file and
verifying hero output references it. No external
services, network access, or shared mutable state
required.

**Coverage strategy**: Unit tests for scaffold engine
changes (new asset paths, file ownership classification).
Integration test: scaffold produces the onboarding
agent and command files. Drift detection: embedded
assets match their canonical sources. No new Go
business logic beyond scaffold registration -- the
onboarding agent is a Markdown persona file.

**Coverage ratchet**: The `expectedAssetPaths` assertion
in `scaffold_test.go` serves as the coverage ratchet --
the count MUST increase from 33 to 35 and MUST NOT
decrease without explicit justification. The drift
detection test (`TestEmbeddedAssets_MatchSource`) ensures
100% sync between canonical and embedded copies.
`isToolOwned` assertions verify file ownership
classification (agent=user-owned, command=tool-owned).

## Project Structure

### Documentation (this feature)

```text
specs/031-agent-customization-onboarding/
├── spec.md              # Feature specification
├── plan.md              # This file
├── research.md          # Phase 0: design decisions
├── data-model.md        # Phase 1: profile schema
├── quickstart.md        # Phase 1: getting started
├── contracts/           # Phase 1: interface contracts
│   └── profile-schema.md
├── checklists/
│   └── requirements.md  # Quality checklist
└── tasks.md             # Phase 2 output (not created here)
```

### Source Code (repository root)

```text
.opencode/
├── agents/
│   └── onboarding.md              # NEW: Onboarding interview agent
├── command/
│   └── onboard.md                 # NEW: /onboard slash command
└── uf/
    └── packs/                     # No changes to convention packs

.uf/
└── onboarding/                    # NEW: User profile storage
    ├── profile.md                 # User profile (YAML frontmatter)
    └── history/                   # Profile version snapshots
        └── YYYY-MM-DDTHH-MM-SS.md

internal/scaffold/
├── scaffold.go                    # MODIFIED: Add onboarding assets
├── scaffold_test.go               # MODIFIED: Update expected paths
└── assets/
    └── opencode/
        ├── agents/
        │   └── onboarding.md      # NEW: Embedded asset copy
        └── command/
            └── onboard.md         # NEW: Embedded asset copy

AGENTS.md                          # MODIFIED: Add profile injection
                                   #   instructions for heroes
```

**Structure Decision**: This feature follows the
established scaffold asset pattern (per Spec 003, 005).
New Markdown files are added to the embedded assets
directory, deployed by `uf init`, and tested via drift
detection. The profile storage location follows the
`.uf/` directory convention (per Spec 025). No new Go
packages are needed -- the scaffold engine gains two
new asset entries.

## Constitution Re-Check (Post-Design)

*Re-evaluated after Phase 1 design artifacts were
produced.*

### I. Autonomous Collaboration — PASS (confirmed)

Profile artifact at `.uf/onboarding/profile.md` is
self-describing per the schema contract
(`contracts/profile-schema.md`). YAML frontmatter
contains version, user, status, and timestamps. Heroes
consume asynchronously via file read. No runtime
coupling introduced.

### II. Composability First — PASS (confirmed)

Onboarding agent deploys independently via `uf init`.
Profile read contract explicitly states: "if profile
doesn't exist, proceed with default behavior." No
mandatory dependencies between onboarding agent and
any hero. Each hero can independently consume the
profile or ignore it.

### III. Observable Quality — PASS (confirmed)

Profile uses structured YAML frontmatter
(machine-parseable). Hero Mapping table is a structured
data section. Profile status (`draft`/`complete`) is a
machine-readable quality signal. Profile history
provides timestamped audit trail.

### IV. Testability — PASS (confirmed)

All deliverables are testable in isolation:
- Scaffold engine: `expectedAssetPaths` assertion
  (33 → 35 files), file ownership classification
- Drift detection: embedded asset copies match live
  counterparts
- Profile schema: YAML frontmatter parseable with
  standard tools
- No external services, no network, no shared state

**Coverage strategy (post-design)**: Unit tests for
scaffold engine changes (2 new expected asset paths,
verify `isToolOwned` returns false for the agent and
true for the command). Drift detection tests verify
embedded copies stay in sync. No new Go business logic
requiring additional unit tests.

## Complexity Tracking

No constitution violations to justify. The design
uses only existing patterns:
- Agent file pattern (per Spec 006)
- Slash command pattern (per Spec 018)
- Scaffold asset embedding (per Spec 003)
- `.uf/` directory convention (per Spec 025)
- YAML frontmatter profile (per Spec 004 backlog pattern)
- AGENTS.md guidance injection (per Spec 030)
