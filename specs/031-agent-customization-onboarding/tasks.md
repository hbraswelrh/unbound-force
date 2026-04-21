# Tasks: Agent Customization Onboarding

**Input**: Design documents from `/specs/031-agent-customization-onboarding/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/profile-schema.md, quickstart.md

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Verify Baseline)

**Purpose**: Confirm the project builds and tests pass
before any changes

- [x] T001 Run `go test -race -count=1 ./...` and confirm all tests pass (baseline)
- [x] T002 Verify `go build ./...` succeeds
- [x] T003 Confirm current asset count: 33 files in `internal/scaffold/assets/`, 33 entries in `expectedAssetPaths` in `internal/scaffold/scaffold_test.go`, `"33 files processed"` assertion in `cmd/unbound-force/main_test.go`

**Checkpoint**: Baseline green. All changes from this
point forward are additive.

---

## Phase 2: US1 — Onboarding Agent and /onboard Command (Priority: P1) MVP

**Goal**: Create the conversational onboarding agent that
captures inspiration, interests, and objectives, persists
them as a structured user profile, and generates a hero
capability mapping. This is the core interview flow.

**Independent Test**: Run `/onboard` in a project with no
existing profile. Complete the interview across all three
dimensions. Verify `.uf/onboarding/profile.md` exists
with `status: complete`, YAML frontmatter with version/
user/timestamps, and all four body sections (Inspiration,
Interests, Objectives, Hero Mapping) populated.

### Create the Onboarding Agent

- [x] T004 [US1] Create `.opencode/agents/onboarding.md` with YAML frontmatter: `description: "Onboarding interview agent — captures user inspiration, interests, and objectives to customize the Unbound Force hero swarm."`, `mode: subagent`, `model: google-vertex-anthropic/claude-opus-4-6@default`, `temperature: 0.4`, `tools: { read: true, write: true, edit: true, bash: false, webfetch: false }` (per research R3, data-model.md)
- [x] T005 [US1] Add `# Role: Onboarding Guide` heading with mission statement: "You conduct a structured interview to capture the user's inspiration, interests, and objectives, then synthesize these into a user profile that contextualizes all Unbound Force heroes." Include explicit constraint: "You MUST NOT create new agents or functionality that duplicates existing hero capabilities" (per FR-004, spec scope)
- [x] T006 [US1] Add `## Hero Capability Reference` section with the static hero capability map table from data-model.md section 3: 5 heroes (Muti-Mind, Cobalt-Crush, Gaze, The Divisor, Mx F) with agent files, capabilities, and example objectives (per research R4, FR-009)
- [x] T007 [US1] Add `## Interview Flow` section defining the three-phase structured interview: Phase 1 (Inspiration — "What does success look like?", "What projects or teams inspire you?", "What values drive your work?"), Phase 2 (Interests — "What domain are you working in?", "What technologies excite you?", "What architectural patterns do you prefer?"), Phase 3 (Objectives — "What are the top 3 things you want to accomplish?", "What timelines are you working with?", "How will you know you've succeeded?") with instructions to ask targeted follow-up questions when answers are vague (per spec US1 acceptance scenario 2, FR-007)
- [x] T008 [US1] Add `## Profile Synthesis` section instructing the agent to: (a) after completing all three dimensions, synthesize the user's inputs into the profile schema defined in `contracts/profile-schema.md`, (b) generate the Hero Mapping table by matching each objective to the hero capability reference table, (c) present the complete profile to the user for review before saving (per spec US1 acceptance scenario 3, FR-009)
- [x] T009 [US1] Add `## Profile Persistence` section with instructions to: (a) create `.uf/onboarding/` directory if it does not exist, (b) write the profile to `.uf/onboarding/profile.md` using the YAML frontmatter schema from `contracts/profile-schema.md` (version, user from `$USER`, status, created_at, modified_at), (c) set `status: draft` during interview, transition to `status: complete` only after user confirms all three dimensions (per FR-002, FR-007, data-model state transitions)
- [x] T010 [US1] Add `## Edge Case Handling` section covering: (a) abandoned interview — persist partial profile with `status: draft`, resume from last completed section on next invocation, (b) contradictory inputs — surface the tension and ask user to prioritize or reconcile, (c) missing heroes — note which heroes are available and which are missing, suggest `uf setup` (per spec edge cases)
- [x] T011 [US1] Add `## Out of Scope` section with explicit boundaries: does not modify existing hero agents, does not create new agents to fulfill unmet objectives, does not manage cross-repository profiles, does not provide authentication or access control (per spec Out of Scope)

### Create the /onboard Command

- [x] T012 [P] [US1] Create `.opencode/command/onboard.md` with instructions to invoke the onboarding agent. The command file should: (a) describe the purpose ("Run the onboarding interview to create or update your user profile"), (b) instruct the main agent to use the Task tool with `subagent_type: "onboarding"` to delegate to the onboarding agent (follow the established pattern from `.opencode/command/review-council.md` which delegates to `divisor-*` agents via Task tool), (c) include the user's input text (if any was typed after `/onboard`) as the prompt context, (d) describe expected outcome: a profile file at `.uf/onboarding/profile.md` (per research R3, established command delegation pattern)

### Create Scaffold Asset Copies

- [x] T013 [P] [US1] Create scaffold asset copy at `internal/scaffold/assets/opencode/agents/onboarding.md` — exact copy of `.opencode/agents/onboarding.md` (canonical source is the live file; scaffold asset is the embedded copy deployed by `uf init`)
- [x] T014 [P] [US1] Create scaffold asset copy at `internal/scaffold/assets/opencode/command/onboard.md` — exact copy of `.opencode/command/onboard.md`

### Update Scaffold Engine Tests

- [x] T015 [US1] Update `expectedAssetPaths` in `internal/scaffold/scaffold_test.go`: add `"opencode/agents/onboarding.md"` to the agents section (after `mx-f-coach.md`, before `divisor-adversary.md`) and `"opencode/command/onboard.md"` to the commands section (after `finale.md`, before `review-council.md`). Total count changes from 33 to 35.
- [x] T016 [US1] Update `cmd/unbound-force/main_test.go`: change the `"33 files processed"` assertion to `"35 files processed"` to match the new embedded asset count

### Verify

- [x] T017 [US1] Run `go test -race -count=1 ./internal/scaffold/...` and verify all scaffold tests pass (expected paths, drift detection, file ownership: `onboarding.md` agent is user-owned, `onboard.md` command is tool-owned)
- [x] T018 [US1] Run `go test -race -count=1 ./cmd/unbound-force/...` and verify the file count assertion passes with the updated 35 count
- [x] T019 [US1] Run `go build ./...` to confirm compilation

**Checkpoint**: The onboarding agent and `/onboard`
command are live, scaffold assets are embedded, and all
tests pass. US1 is complete — users can run `/onboard`
to create a profile.

---

## Phase 3: US2 — Hero Profile Injection (Priority: P1)

**Goal**: Inject user profile reading instructions into
AGENTS.md so all existing heroes automatically
contextualize their behavior based on the profile. No
hero agent files are modified.

**Independent Test**: Create a test profile at
`.uf/onboarding/profile.md` with security-focused
objectives. Invoke the Divisor and verify its review
mentions security context from the profile. Delete the
profile and verify the Divisor operates with default
behavior unchanged.

### Add Profile Section to AGENTS.md

- [x] T020 [US2] Add a `## User Profile` section to `AGENTS.md` (after the "Knowledge Retrieval" section, before "Testing Conventions") with the following content: (a) explanation that `.uf/onboarding/profile.md` contains the user's customization context, (b) instructions for all agents to check if the profile exists and read it if present, (c) per-dimension guidance on how to use each section (Inspiration → tone/values/quality bar, Interests → domain/technology/conventions, Objectives → priorities/timelines/focus areas, Hero Mapping → role-specific priorities), (d) explicit instruction: "If the profile does not exist, proceed with your default behavior unchanged" (per research R2, FR-003, FR-006, contracts/profile-schema.md Profile Read Contract)

### Update /uf-init for Profile Guidance Injection

- [x] T021 [US2] Add a new AGENTS.md guidance block to `.opencode/command/uf-init.md` for injecting the "User Profile" section into target repositories' `AGENTS.md` files during `/uf-init`. Follow the Spec 030 idempotent injection pattern: check for existing `## User Profile` heading before injecting, skip if already present (per research R7, Spec 030 pattern)
- [x] T022 [P] [US2] Update the scaffold asset copy at `internal/scaffold/assets/opencode/command/uf-init.md` to match the live `.opencode/command/uf-init.md` after the guidance block addition (scaffold sync)

### Verify

- [x] T023 [US2] Read `AGENTS.md` and verify the User Profile section exists with profile path, per-dimension guidance, and backward-compatibility instruction
- [x] T024 [US2] Run `go test -race -count=1 ./internal/scaffold/...` to verify scaffold drift detection passes after syncing `uf-init.md`

**Checkpoint**: All heroes now read the user profile
when it exists. US2 is complete — hero customization
via profile is active. Combined with US1, this is the
full MVP.

---

## Phase 4: US3 — Profile Update and Evolution (Priority: P2)

**Goal**: Enable users to update their profile without
starting from scratch, and maintain a history of previous
profiles.

**Independent Test**: Create a profile via `/onboard`.
Re-invoke `/onboard` and update only the Objectives
dimension. Verify: (a) the previous profile was saved to
`.uf/onboarding/history/`, (b) the updated profile has
new objectives but unchanged Inspiration and Interests,
(c) `modified_at` is updated while `created_at` is
preserved.

### Add Update Mode to Onboarding Agent

- [x] T025 [US3] Add `## Update Mode` section to `.opencode/agents/onboarding.md`: when a profile already exists at `.uf/onboarding/profile.md`, the agent MUST (a) read and display the current profile, (b) ask the user which dimension(s) they want to update rather than restarting entirely, (c) preserve unchanged dimensions, (d) create a history snapshot before overwriting (per spec US3 acceptance scenarios 1-2, FR-005, contracts/profile-schema.md Profile Write Contract)
- [x] T026 [US3] Add `## Profile History` section to `.opencode/agents/onboarding.md`: (a) before overwriting `profile.md`, copy current version to `.uf/onboarding/history/<modified_at timestamp>.md`, (b) history files are immutable once created, (c) when user asks to view history, list files in `.uf/onboarding/history/` with timestamps (per FR-008, data-model section 2, spec US3 acceptance scenario 3)

### Sync Scaffold Asset

- [x] T027 [P] [US3] Update scaffold asset copy at `internal/scaffold/assets/opencode/agents/onboarding.md` to match the live `.opencode/agents/onboarding.md` after US3 additions

### Verify

- [x] T028 [US3] Run `go test -race -count=1 ./internal/scaffold/...` to verify scaffold drift detection passes after syncing `onboarding.md`

**Checkpoint**: US3 complete. Users can update profiles
incrementally and view history.

---

## Phase 5: US4 — Guardrail Against Agent Duplication (Priority: P2)

**Goal**: When a user describes an objective that overlaps
with an existing hero's capabilities, the onboarding
agent routes them to the existing hero rather than
suggesting new agent creation.

**Independent Test**: Run `/onboard` and state "I need
something to check my test coverage." Verify the agent
identifies Gaze as the existing capability and does NOT
suggest creating a new agent. Try "I want automated code
reviews" and verify the agent identifies The Divisor
Council.

### Enhance Objective Processing

- [x] T029 [US4] Add `## Objective-to-Hero Routing` section to `.opencode/agents/onboarding.md`: when the user describes an objective during the interview, the agent MUST (a) compare the objective against the Hero Capability Reference table, (b) if a match is found, explain which hero handles this and how to configure it rather than creating new functionality, (c) if no match is found, record the objective as a custom goal in the profile and suggest which existing hero is closest to fulfilling it (per spec US4 acceptance scenarios 1-3, FR-004, FR-010)

### Sync Scaffold Asset

- [x] T030 [P] [US4] Update scaffold asset copy at `internal/scaffold/assets/opencode/agents/onboarding.md` to match the live `.opencode/agents/onboarding.md` after US4 additions

### Verify

- [x] T031 [US4] Run `go test -race -count=1 ./internal/scaffold/...` to verify scaffold drift detection passes after syncing `onboarding.md`

**Checkpoint**: US4 complete. The agent routes users to
existing heroes instead of duplicating functionality.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Documentation, final sync, and validation

### Documentation Updates

- [x] T032 [P] Update `AGENTS.md` "Project Structure" section: add `onboarding.md` to the agents list and `onboard.md` to the commands list in the file tree under `.opencode/`
- [x] T033 [P] Update `AGENTS.md` "The Heroes" table or add an entry for the onboarding agent (note: it is NOT a hero — it is a utility agent that guides users to heroes). Add a row to the appropriate section or add a new subsection for non-hero agents.
- [x] T034 [P] Update `AGENTS.md` "Recent Changes" section: add a `031-agent-customization-onboarding` entry summarizing all changes (new agent, new command, AGENTS.md profile section, scaffold asset additions, test count updates)
- [x] T035 Update `AGENTS.md` active technologies section if the agent context script did not already add the relevant entries

### Website Documentation Issue (Constitution MUST)

- [x] T036 [P] Create a GitHub issue in `unbound-force/website` to track documentation for the new `/onboard` command and user profile workflow: `gh issue create --repo unbound-force/website --title "docs: new /onboard command and user profile customization workflow" --body "Spec 031 adds a new /onboard slash command and onboarding agent that captures user inspiration, interests, and objectives into a profile at .uf/onboarding/profile.md. Heroes automatically read this profile to contextualize their behavior. Pages to create/update: getting started guide (mention /onboard), hero customization workflow, profile schema reference."` — this MUST be created before the implementing PR is merged (per constitution Development Workflow, Cross-Repo Documentation rule)

### Final Scaffold Sync Verification

- [x] T037 Run `go test -race -count=1 ./...` to verify all tests pass across the entire project
- [x] T038 Run `go build ./...` to confirm final compilation
- [x] T039 Validate quickstart.md scenarios: confirm `/onboard` command is listed in `.opencode/command/`, confirm `onboarding` agent is listed in `.opencode/agents/`, confirm `.uf/onboarding/` path conventions match the data-model

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — can start
  immediately
- **US1 (Phase 2)**: Depends on Setup — creates the
  agent, command, scaffold assets, and tests. BLOCKS
  all subsequent phases.
- **US2 (Phase 3)**: Depends on US1 — needs the
  agent and command to exist before injecting profile
  instructions into AGENTS.md
- **US3 (Phase 4)**: Depends on US1 — extends the
  agent with update mode and history
- **US4 (Phase 5)**: Depends on US1 — extends the
  agent with routing guardrails
- **Polish (Phase 6)**: Depends on US1-US4 completion

### User Story Dependencies

- **US1 (P1)**: Foundational — all other stories depend
  on the agent and command existing
- **US2 (P1)**: Depends on US1 — profile must exist
  before heroes can read it. Can run in parallel with
  US3 and US4 after US1 completes.
- **US3 (P2)**: Depends on US1 — update mode needs
  the initial interview flow. Can run in parallel with
  US2 and US4.
- **US4 (P2)**: Depends on US1 — routing needs the
  hero capability map. Can run in parallel with US2
  and US3.

### Within Each User Story

- Agent content tasks are sequential (building sections
  of the same file)
- Scaffold asset sync tasks ([P]) can run in parallel
  with each other
- Test verification tasks run after their story's
  implementation tasks

### Parallel Opportunities

- T012 (command file) and T013-T014 (scaffold copies)
  can run in parallel with each other after T004-T011
- T020 (AGENTS.md) and T021-T022 (uf-init) can run in
  parallel within US2
- US3 (Phase 4) and US4 (Phase 5) can run in parallel
  after US1 completes
- T032, T033, T034 (docs updates) can all run in
  parallel

---

## Parallel Example: After US1 Completes

```text
# US2, US3, and US4 can start in parallel:

Agent A (US2): T020 → T021 → T022 → T023 → T024
Agent B (US3): T025 → T026 → T027 → T028
Agent C (US4): T029 → T030 → T031

# Then Polish phase after all converge:
All: T032-T039
```

---

## Implementation Strategy

### MVP First (US1 + US2 Only)

1. Complete Phase 1: Setup (T001-T003)
2. Complete Phase 2: US1 (T004-T019) — agent + command
3. Complete Phase 3: US2 (T020-T024) — profile injection
4. **STOP and VALIDATE**: Test `/onboard` end-to-end,
   verify heroes read the profile
5. Deploy/demo if ready

### Incremental Delivery

1. US1 → agent exists, users can create profiles
2. US2 → heroes use profiles → **Full MVP**
3. US3 → users can update profiles without restart
4. US4 → guardrails prevent agent duplication
5. Polish → documentation and final validation

### Parallel Team Strategy

With multiple agents after US1:
- Agent A: US2 (AGENTS.md + uf-init)
- Agent B: US3 (update mode + history)
- Agent C: US4 (routing guardrails)
- All converge for Polish phase

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story
- **v1 scope: single-user only** — the profile path
  `.uf/onboarding/profile.md` supports one profile per
  project. The spec edge case about multiple team
  members with different profiles is acknowledged but
  deferred to a future iteration (per research R5).
  The YAML frontmatter `user` field is included for
  forward compatibility.
- Agent file (`onboarding.md`) is **user-owned** — not
  overwritten by `uf init` if it already exists
- Command file (`onboard.md`) is **tool-owned** — kept
  canonical by `uf init`
- No new Go packages or dependencies — only Markdown
  files and test assertion updates
- Profile files at `.uf/onboarding/` are NOT gitignored
  (they are user configuration, not runtime data)
- Total tasks: 39
- Total new embedded assets: 2 (33 → 35)
