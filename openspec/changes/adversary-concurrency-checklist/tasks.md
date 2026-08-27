<!--
Tasks are grouped and numbered. `[P]` marks tasks that are
parallel-eligible (no ordering dependency on sibling tasks in
the same group). Checkboxes are marked `[x]` immediately on
completion. Implementation tasks are pre-marked complete
because this change is an OpenSpec correction: the edits were
implemented directly before the spec was written, and the spec
now formalizes the already-verified change.
-->

# Tasks: Adversary Concurrency Correctness Checklist

## 1. Specification

- [x] 1.1 Write `proposal.md` (Why, What Changes, Capabilities,
  Impact, Constitution Alignment for all five org principles)
- [x] 1.2 Write `design.md` (decisions D1–D5, risks,
  verification)
- [x] 1.3 Write `specs/adversary-concurrency-checklist.md`
  spec delta (ADDED Requirements with FR-NNN + scenarios)

## 2. Implementation

- [x] 2.1 Add "Concurrency Correctness" section to
  `.opencode/agents/divisor-adversary.md` after "Error Handling
  and Resilience", with a production-scope note and five
  auditable interrogative items
- [x] 2.2 [P] Renumber subsequent Adversary checklist sections
  (Path/Injection, Adversarial Input, Language-Specific,
  Gate Tampering)
- [x] 2.3 [P] Add Adversary/SRE concurrency-boundary row to the
  Out of Scope table
- [x] 2.4 Sync the change to
  `internal/scaffold/assets/opencode/agents/divisor-adversary.md`
  via drift remediation (`cp` canonical → mirror)

## 3. Verification

- [x] 3.1 Run `go test -race -count=1 ./internal/scaffold/`;
  confirm `TestEmbeddedAssets_MatchSource` and
  `TestEmbeddedAssets_SingleMarker` pass
- [x] 3.2 Confirm the new section contains ≥ 5 items and is
  scoped to production code
- [x] 3.3 Confirm existing Out of Scope boundaries are
  preserved
- [x] 3.4 Constitution alignment verification: re-confirm all
  five org principles (I–V) assess PASS as documented in
  `proposal.md`

## 4. Governance (pre-PR)

- [ ] 4.1 Run `/uf.review-council` and resolve any
  REQUEST CHANGES before opening the PR
- [ ] 4.2 File follow-up issue for reciprocal `divisor-sre.md`
  Out of Scope row (deferred Option A)
- [ ] 4.3 File cross-repo website documentation issue for the
  user-facing agent capability change

<!-- spec-review: pending -->
<!-- code-review: pending -->
