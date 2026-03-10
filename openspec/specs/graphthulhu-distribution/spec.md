---
capability: graphthulhu-distribution
introduced: 2026-03-09
change: graphthulhu-homebrew-distribution
---

## Requirements

### Requirement: graphthulhu Homebrew cask

The `unbound-force/homebrew-tap` repository MUST contain a
`Casks/graphthulhu.rb` cask file that installs the graphthulhu
binary for macOS (Intel and Apple Silicon) and Linux (amd64 and
arm64). The cask MUST pin to a specific released version with
verified SHA-256 checksums for each platform archive.

#### Scenario: Install graphthulhu via tap

- **GIVEN** a user has added `unbound-force/tap` to Homebrew
- **WHEN** the user runs `brew install unbound-force/tap/graphthulhu`
- **THEN** the graphthulhu binary is installed and available on `$PATH`

#### Scenario: Platform-correct archive selected

- **GIVEN** the graphthulhu cask is installed on macOS Apple Silicon
- **WHEN** Homebrew resolves the cask
- **THEN** Homebrew downloads the `darwin_arm64` archive and its
  checksum matches the declared SHA-256

### Requirement: unbound cask declares graphthulhu dependency

Every release of the `unbound` Homebrew cask MUST declare
`depends_on cask: "graphthulhu"` so that installing `unbound`
automatically installs graphthulhu. This dependency MUST be
injected via the GoReleaser `dependencies` field so it survives
every automated release without manual patching.

#### Scenario: Single-command environment setup

- **GIVEN** a fresh macOS or Linux machine with Homebrew installed
- **WHEN** the user runs `brew install unbound-force/tap/unbound`
- **THEN** both `unbound` and `graphthulhu` are installed and the
  `graphthulhu` binary is available on `$PATH`

#### Scenario: GoReleaser dependency applied on release

- **GIVEN** a GoReleaser release is triggered for `unbound`
- **WHEN** GoReleaser generates `Casks/unbound.rb` in the tap repo
- **THEN** the generated cask contains `depends_on cask: "graphthulhu"`
  and the correct version, URL, and SHA-256 values

### Requirement: GoReleaser cask dependency configuration

The `unbound-force/unbound-force` repository MUST declare
`dependencies: [{cask: graphthulhu}]` in the `homebrew_casks`
stanza of `.goreleaser.yaml`. This ensures every automated release
produces an `unbound` cask with the graphthulhu dependency intact.

#### Scenario: Dependency present in goreleaser config

- **GIVEN** the `unbound-force/unbound-force` repository
- **WHEN** a developer inspects `.goreleaser.yaml`
- **THEN** `homebrew_casks[].dependencies` includes
  `{cask: graphthulhu}`
