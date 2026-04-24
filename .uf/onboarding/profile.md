---
version: "1.0.0"
status: experimental
created: 2026-04-21
updated: 2026-04-21
---

## Inspiration

Secure, spec-governed code. Inspired by the OpenSSF ecosystem — SLSA, Complytime, Carabiner-dev Ampel, Gemara. Admires the culture of composable security tooling and community-driven standards. Core values: developer experience, reliability, and security.

## Interests

Software supply chain security and compliance tooling. Go-centric stack with Sigstore, in-toto, and CUE. Builds CLI tools, web UIs, and libraries following policy-as-code and attestation pipeline patterns. Active contributor to OpenSSF and OSCAL. Strategic adopter of AI Native Development.

## Objectives

1. Ship fuzzybunny v0.1.0 — new supply chain security/compliance tool
2. Achieve SLSA Build Level 3 across repositories
3. Conference talks and demos — hard deadlines driven by conference schedule
4. Improve CRAP scores and code quality via static analysis
5. Reduce time spent on line-by-line code review
6. Keep UX front-and-center — biggest risk is losing user experience focus

Success metrics: SLSA build levels, CRAP scores, community adoption and reuse.

## Hero Mapping

| Objective | Hero | How to Leverage |
|-----------|------|-----------------|
| Ship fuzzybunny v0.1.0 | Muti-Mind + Cobalt-Crush | Muti-Mind prioritizes the backlog and drafts specs via `/workflow seed`; Cobalt-Crush implements from specs with Go conventions and TDD |
| Achieve SLSA Build Level 3 | The Divisor (Adversary) + Cobalt-Crush | Adversary audits for supply chain risks and SLSA compliance gaps; Cobalt-Crush implements Sigstore signing, in-toto attestations, provenance generation |
| Conference talks & demos | Mx F + Muti-Mind | Mx F manages sprint cadence toward conference deadlines with velocity tracking; Muti-Mind gates what's demo-ready via acceptance decisions |
| Improve CRAP scores | Gaze | Track CRAP scores and coverage trends with machine-parseable quality reports feeding into Mx F metrics |
| Reduce code review time | The Divisor (full council) | 5-persona automated review (Guard, Architect, Adversary, SRE, Testing) via `/review-council` — biggest leverage point |
| Keep UX front-and-center | Muti-Mind + The Divisor (Guard) | Muti-Mind enforces UX-focused acceptance criteria; Guard detects intent drift from spec |
| Community adoption & reuse | Gaze | Observable quality evidence — reproducible, machine-parseable reports build community trust |
