# Add Concurrency Correctness Checklist to The Adversary

## Why

Production-code concurrency correctness is an unowned review
gap across the Divisor review panel. A grep of all nine
`divisor-*.md` agent files confirms that no persona reviews
production goroutine lifecycle, lock ordering, unsafe
concurrent access, `context.Context` propagation, or
`sync.WaitGroup` misuse:

- The Adversary owns "Error Handling and Resilience" but its
  checklist contains zero concurrency-correctness items.
- The SRE owns efficiency/performance and reliability, but has
  no goroutine-leak, deadlock, or lock-ordering items.
- The Tester (`divisor-testing.md`) covers only *test-code*
  concurrency ("tests compatible with concurrent execution").

The `-race` CI flag only detects data races on code paths that
are actually exercised at runtime; static review of concurrency
patterns is a distinct, complementary control that no persona
currently performs. Goroutine leaks (resource exhaustion → DoS)
and deadlocks (service unavailability) are genuine
security/resilience failure modes that fit The Adversary's
declared "Security & Resilience" domain.

Discovered via gaze PR #207 and corroborated by a 5-agent
Divisor triage panel (4/5 VALID, unanimous `enhancement`,
objective) on issue #469 (parent epic #465).

## What Changes

- Add a **Concurrency Correctness** checklist section to The
  Adversary agent, placed within its Security & Resilience
  domain, immediately after "Error Handling and Resilience."
- The section contains **five auditable interrogative review
  items** covering goroutine leaks/cancellation, deadlock and
  lock-ordering/channel blocking, unsafe concurrent map/slice/
  shared access, missing `context.Context` propagation, and
  `sync.WaitGroup` misuse.
- The section is scoped to **production code only**; test-code
  concurrency remains owned by The Tester.
- Add an **Out of Scope** row demarcating the Adversary/SRE
  boundary: concurrency *efficiency* (lock contention,
  throughput tuning) → The SRE, while concurrency *correctness*
  (leaks, deadlocks, unsafe access) is owned by The Adversary.
- Synchronize the change to the embedded scaffold mirror so the
  drift-detection regression test remains green.

## Capabilities

### New

- Static production-code concurrency-correctness review as an
  explicit, auditable capability of The Adversary persona.

### Modified

- `divisor-adversary.md` review checklist gains a numbered
  Concurrency Correctness section; subsequent sections are
  renumbered.
- `divisor-adversary.md` Out of Scope table gains an
  Adversary/SRE concurrency-boundary row.

### Removed

- None.

## Impact

### New files

- `openspec/changes/adversary-concurrency-checklist/` (this
  change's proposal, design, tasks, and spec delta).

### Modified files

- `.opencode/agents/divisor-adversary.md` (canonical source).
- `internal/scaffold/assets/opencode/agents/divisor-adversary.md`
  (embedded scaffold mirror; kept byte-identical).

### Tests

- `internal/scaffold` drift-detection tests
  (`TestEmbeddedAssets_MatchSource`,
  `TestEmbeddedAssets_SingleMarker`) MUST remain green,
  enforcing byte-identical parity between the two copies.

### Out of scope

- Reciprocal edit to `divisor-sre.md` Out of Scope table
  (deferred to a possible follow-up issue).
- Test-code concurrency review (owned by The Tester).
- Concurrency efficiency/performance tuning (owned by The SRE).

## Constitution Alignment

Assessed against the Unbound Force org constitution v1.2.0.

### I. Autonomous Collaboration

**Assessment**: PASS. The change is a self-contained edit to a
single agent persona plus its mirror. It clarifies artifact
ownership boundaries (Adversary vs. SRE vs. Tester), reducing
cross-hero finding fragmentation. No runtime coupling is
introduced.

### II. Composability First

**Assessment**: PASS. The Adversary remains independently
usable. The new checklist adds value when the persona reviews
Go code standalone and composes cleanly with SRE and Tester
scopes via the explicit Out of Scope boundary.

### III. Observable Quality

**Assessment**: PASS. Each new item is phrased as an auditable
interrogative review question, producing observable
pass/fail review signal. The drift-detection test provides
machine-verifiable evidence that both copies stay in sync.

### IV. Testability

**Assessment**: PASS. The change is guarded by an existing
automated regression gate (byte-identical drift test) that
blocks the build on divergence. Verification is deterministic
and requires no external services.

### V. Security by Default

**Assessment**: PASS. The change directly strengthens the
security posture of reviewed code by adding static detection of
concurrency-driven DoS/availability failure modes (goroutine
leaks, deadlocks, unsafe shared access) that the runtime `-race`
gate cannot fully cover. No dependencies are added; no
privilege or input-validation surface changes.
