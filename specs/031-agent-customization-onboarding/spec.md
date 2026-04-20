# Feature Specification: Agent Customization Onboarding

**Feature Branch**: `031-agent-customization-onboarding`
**Created**: 2026-04-20
**Status**: Draft
**Input**: User description: "Create an agent that allows
a user to input inspiration, interest, and objectives to
customize and leverage the agents in unbound-force, but
with the constraint that they should be using the
unbound-force projects agents and not re-creating the same
functionality"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Initial Onboarding Interview (Priority: P1)

A new user installs Unbound Force and wants the agent
swarm to work in a way that aligns with their project
goals and personal working style. They invoke the
onboarding agent, which conducts a structured
conversational interview to capture three dimensions:
**inspiration** (what motivates them, what "great" looks
like for their project), **interests** (their domain,
technology preferences, areas of focus), and
**objectives** (concrete goals, timelines, success
metrics). The agent synthesizes these inputs into a
user profile that becomes the contextual lens through
which all existing Unbound Force heroes operate.

**Why this priority**: P1 because without capturing the
user's context, none of the downstream customization can
happen. This is the foundational data-gathering step that
all other stories depend on.

**Independent Test**: Can be fully tested by invoking the
onboarding agent, completing the interview, and verifying
that a user profile is persisted with all three dimensions
populated.

**Acceptance Scenarios**:

1. **Given** a user has Unbound Force installed but no
   user profile exists, **When** they invoke the
   onboarding agent, **Then** the agent conducts a
   structured interview covering inspiration, interests,
   and objectives, and persists a user profile.
2. **Given** the interview is in progress, **When** the
   user provides vague or incomplete answers, **Then** the
   agent asks targeted follow-up questions to elicit
   actionable detail before moving to the next dimension.
3. **Given** the interview is complete, **When** the user
   reviews the synthesized profile, **Then** they can
   confirm, edit, or restart any section before
   finalizing.

---

### User Story 2 - Hero Customization via Profile (Priority: P1)

After the onboarding interview produces a user profile,
the existing Unbound Force heroes automatically adjust
their behavior based on the profile. Muti-Mind
prioritizes backlog items aligned with the user's stated
objectives. Cobalt-Crush applies coding conventions
that match the user's domain and interests. The
Divisor emphasizes review criteria relevant to the
user's quality priorities. Gaze focuses test coverage
on areas the user identified as high-risk. Mx F tracks
metrics aligned with the user's success criteria.

The heroes themselves are NOT modified or replaced.
Instead, the user profile is injected as additional
context that each hero's existing agent file can
reference when making decisions.

**Why this priority**: P1 because this is the core value
proposition -- the user's input must actually change how
the existing agents operate, otherwise the onboarding
interview is pointless.

**Independent Test**: Can be tested by creating a user
profile with specific objectives (e.g., "optimize for
security"), then invoking each hero and verifying their
output reflects the profile context (e.g., Divisor
emphasizes security findings, Gaze prioritizes security
test coverage).

**Acceptance Scenarios**:

1. **Given** a completed user profile exists with
   security-focused objectives, **When** the Divisor
   reviews code, **Then** the Divisor's review includes
   heightened attention to security concerns as specified
   in the profile.
2. **Given** a completed user profile exists with
   performance-focused interests, **When** Muti-Mind
   prioritizes the backlog, **Then** performance-related
   items receive higher priority weight.
3. **Given** no user profile exists, **When** any hero
   is invoked, **Then** it operates with its default
   behavior unchanged (backward compatible).

---

### User Story 3 - Profile Update and Evolution (Priority: P2)

Over time, a user's inspiration, interests, and
objectives change. The user can re-invoke the onboarding
agent to update their profile without starting from
scratch. The agent shows the current profile, lets the
user modify specific dimensions, and propagates the
changes to all heroes immediately.

**Why this priority**: P2 because initial onboarding
(US1) and hero customization (US2) must work first.
Profile evolution is important for long-term retention
but is not required for the first usable version.

**Independent Test**: Can be tested by creating a profile,
modifying one dimension (e.g., changing objectives from
"speed" to "reliability"), and verifying all heroes
reflect the updated context on their next invocation.

**Acceptance Scenarios**:

1. **Given** a user profile already exists, **When** the
   user invokes the onboarding agent, **Then** the agent
   displays the current profile and offers to update
   specific sections rather than restarting entirely.
2. **Given** the user updates their objectives, **When**
   they finalize the changes, **Then** all heroes use the
   updated profile on their next invocation.
3. **Given** the user wants to track how their priorities
   have shifted, **When** they view their profile history,
   **Then** they can see previous versions with timestamps.

---

### User Story 4 - Guardrail Against Agent Duplication (Priority: P2)

During onboarding, a user describes objectives that
overlap with existing hero capabilities (e.g., "I want
an agent that reviews my code for security issues").
Instead of creating a new agent, the onboarding agent
recognizes this as The Divisor's Adversary persona's
responsibility and routes the user to configure The
Divisor's existing security review capabilities. The
onboarding agent acts as a guide to the existing
ecosystem, never as a creator of redundant functionality.

**Why this priority**: P2 because this guardrail is
essential to the design philosophy, but only becomes
relevant after the core interview and customization
flows work (US1, US2).

**Independent Test**: Can be tested by describing
objectives that match existing hero capabilities and
verifying the agent routes to the correct hero rather
than suggesting new agent creation.

**Acceptance Scenarios**:

1. **Given** a user states "I need something to check my
   test coverage," **When** the onboarding agent
   processes this, **Then** it identifies Gaze as the
   existing capability and guides the user to configure
   Gaze's coverage thresholds.
2. **Given** a user states "I want automated code
   reviews," **When** the onboarding agent processes
   this, **Then** it identifies The Divisor Council and
   explains the five review personas available.
3. **Given** a user describes an objective that has no
   existing hero match, **When** the onboarding agent
   processes this, **Then** it records the objective in
   the profile as a custom goal and suggests which
   existing hero is closest to fulfilling it.

---

### Edge Cases

- What happens when a user abandons the interview
  midway? The agent persists a partial profile and
  resumes from the last completed section on next
  invocation.
- What happens when a user provides contradictory
  inputs across dimensions (e.g., objectives say "move
  fast" but inspiration says "zero-defect culture")?
  The agent surfaces the tension and asks the user to
  prioritize or reconcile.
- What happens when the user profile references hero
  capabilities that are not yet installed? The agent
  notes which heroes are available and which are missing,
  and suggests installation steps.
- What happens when multiple team members each have
  different profiles in the same repository? Each
  user's profile is independent and applies only to
  their own sessions.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide a conversational
  onboarding agent that captures user input across three
  dimensions: inspiration, interests, and objectives.
- **FR-002**: System MUST persist the user profile as a
  structured file within the project's `.uf/` directory
  so that all heroes can access it.
- **FR-003**: System MUST inject the user profile as
  contextual input to each existing hero agent without
  modifying the hero's agent file or capabilities.
- **FR-004**: System MUST NOT create new agents or
  functionality that duplicates existing hero
  capabilities. The onboarding agent routes users to
  existing heroes for any request that matches an
  existing hero's role.
- **FR-005**: System MUST support updating an existing
  profile without requiring the user to redo the
  entire onboarding interview.
- **FR-006**: System MUST maintain backward
  compatibility -- heroes operate with default behavior
  when no user profile exists.
- **FR-007**: System MUST validate that the user profile
  contains actionable content in all three dimensions
  before marking it as complete.
- **FR-008**: System SHOULD persist profile history so
  users can see how their priorities evolved over time.
- **FR-009**: System MUST map user-described objectives
  to existing hero capabilities and present the mapping
  to the user during onboarding.
- **FR-010**: System SHOULD detect when a user describes
  an objective that no existing hero covers and record
  it as an unmet need in the profile.

### Key Entities

- **User Profile**: Represents a user's customization
  context. Contains three dimensions (inspiration,
  interests, objectives), a status (draft, complete),
  timestamps, and a hero capability mapping.
- **Hero Capability Map**: A reference structure that
  maps user-described objectives to existing hero
  agents, including the hero name, relevant agent file,
  and the specific capability that matches.
- **Profile History Entry**: A timestamped snapshot of a
  previous profile version, enabling users to track how
  their priorities have evolved.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users complete the onboarding interview
  and produce a valid profile in under 10 minutes.
- **SC-002**: 100% of user-described objectives that
  match existing hero capabilities are correctly mapped
  to the corresponding hero during onboarding.
- **SC-003**: All five heroes (Muti-Mind, Cobalt-Crush,
  Gaze, The Divisor, Mx F) reflect user profile context
  in their output when a profile exists.
- **SC-004**: Zero new agents are created by the
  onboarding process that duplicate existing hero
  functionality.
- **SC-005**: Users can update any single dimension of
  their profile in under 3 minutes without affecting
  the other dimensions.
- **SC-006**: Heroes operate identically to their
  default behavior when no user profile is present
  (zero regressions).

## Dependencies and Assumptions

### Dependencies

- **Spec 004** (Muti-Mind Architecture): The onboarding
  agent must understand Muti-Mind's prioritization
  capabilities to route objectives correctly.
- **Spec 005** (The Divisor Architecture): The onboarding
  agent must understand the five Divisor personas to
  map review-related objectives.
- **Spec 006** (Cobalt-Crush Architecture): The
  onboarding agent must understand Cobalt-Crush's
  coding conventions integration.
- **Spec 007** (Mx F Architecture): The onboarding agent
  must understand Mx F's metrics and coaching
  capabilities.
- **Spec 008** (Swarm Orchestration): The user profile
  must integrate with the hero lifecycle workflow.
- **Spec 030** (AGENTS.md Guidance): The onboarding
  agent's behavioral guidance should be injectable via
  the established `/uf-init` pattern.

### Assumptions

- The user profile is stored per-user within the
  project repository (not globally across projects).
  Each project may have different objectives.
- The onboarding agent is implemented as an OpenCode
  agent file (Markdown-based persona), consistent with
  the existing hero agent pattern.
- Profile injection into heroes is achieved by making
  the user profile file available in the agent's
  context (referenced from AGENTS.md or agent files),
  not by modifying the hero's core logic.
- The onboarding agent has read access to all existing
  agent files to build its hero capability map.
- Interview persistence uses the same `.uf/` directory
  convention established in Spec 025.

## Scope

### In Scope

- Conversational onboarding interview agent
- User profile schema and persistence
- Profile injection mechanism for existing heroes
- Hero capability mapping and routing
- Profile update and history tracking
- Guardrails against agent duplication

### Out of Scope

- Modifying any existing hero agent's core logic or
  persona definition
- Creating new hero agents to fulfill unmet objectives
- Cross-repository profile sharing or synchronization
- Team-level aggregated profiles or organizational
  dashboards
- Authentication or access control for user profiles
