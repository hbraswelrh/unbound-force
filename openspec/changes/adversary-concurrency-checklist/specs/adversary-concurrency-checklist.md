## ADDED Requirements

### Requirement: Adversary Concurrency Correctness Review

FR-001: The Adversary agent (`divisor-adversary.md`) MUST
include a dedicated Concurrency Correctness review section
within its Security & Resilience domain, placed immediately
after the Error Handling and Resilience section.

FR-002: The Concurrency Correctness section MUST contain at
least five review items, each phrased as an auditable
interrogative review question consistent with the existing
checklist style. The items MUST collectively cover: goroutine
leaks and cancellation/join paths; deadlock potential
(inconsistent lock ordering, locks held across blocking calls,
unbuffered channel blocking); unsafe concurrent access to
shared maps, slices, or structs; `context.Context` propagation
for cancellation and deadlines; and `sync.WaitGroup` misuse.

FR-003: The Concurrency Correctness section MUST be scoped to
production code and MUST explicitly state that test-code
concurrency (race-detector and parallel-runner compatibility)
is owned by The Tester (`divisor-testing.md`).

FR-004: The Adversary Out of Scope table MUST include a row
demarcating the Adversary/SRE concurrency boundary: concurrency
*efficiency* (lock contention, throughput tuning) is routed to
The SRE, while concurrency *correctness* (leaks, deadlocks,
unsafe access) is owned by The Adversary. All pre-existing Out
of Scope boundaries MUST be preserved.

FR-005: The change MUST be applied identically to both the
canonical source (`.opencode/agents/divisor-adversary.md`) and
the embedded scaffold mirror
(`internal/scaffold/assets/opencode/agents/divisor-adversary.md`).
The two copies MUST remain byte-identical so the scaffold
drift-detection tests pass.

#### Scenario: Concurrency section present and production-scoped

- **GIVEN** the maintainer opens `divisor-adversary.md`
- **WHEN** they read the Security & Resilience checklist
- **THEN** a Concurrency Correctness section appears
  immediately after Error Handling and Resilience
- **AND** the section contains at least five interrogative
  review items
- **AND** the section states test-code concurrency is owned by
  The Tester

#### Scenario: Adversary/SRE boundary demarcated

- **GIVEN** the maintainer reads the Adversary Out of Scope
  table
- **WHEN** they look for concurrency routing
- **THEN** concurrency efficiency (lock contention, throughput
  tuning) is routed to The SRE
- **AND** concurrency correctness (leaks, deadlocks, unsafe
  access) is stated to be owned by The Adversary
- **AND** all pre-existing Out of Scope rows remain unchanged

#### Scenario: Embedded mirror stays in sync

- **GIVEN** the canonical `.opencode/agents/divisor-adversary.md`
  has been edited
- **AND** the change has been copied to the embedded mirror
  under `internal/scaffold/assets/`
- **WHEN** `go test -race -count=1 ./internal/scaffold/` runs
- **THEN** `TestEmbeddedAssets_MatchSource` passes
- **AND** `TestEmbeddedAssets_SingleMarker` passes
