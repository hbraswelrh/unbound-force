## ADDED Requirements

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
injected via a GoReleaser cask template so it survives every
automated release without manual patching.

#### Scenario: Single-command environment setup

- **GIVEN** a fresh macOS or Linux machine with Homebrew installed
- **WHEN** the user runs `brew install unbound-force/tap/unbound`
- **THEN** both `unbound` and `graphthulhu` are installed and the
  `graphthulhu` binary is available on `$PATH`

#### Scenario: GoReleaser template applied on release

- **GIVEN** a GoReleaser release is triggered for `unbound`
- **WHEN** GoReleaser generates `Casks/unbound.rb` in the tap repo
- **THEN** the generated cask contains `depends_on cask: "graphthulhu"`
  and the correct version, URL, and SHA-256 values

### Requirement: GoReleaser cask template file

The `unbound-force/unbound-force` repository MUST contain a
GoReleaser cask template file (e.g., `.goreleaser-cask.rb.tmpl`)
that is referenced by the `homebrew_casks[].cask_template` field
in `.goreleaser.yaml`. The template MUST use GoReleaser's standard
template variables for version, URL, and SHA-256, and MUST include
the static `depends_on cask: "graphthulhu"` declaration.

#### Scenario: Template file present in repo

- **GIVEN** the `unbound-force/unbound-force` repository
- **WHEN** a developer inspects `.goreleaser.yaml`
- **THEN** `homebrew_casks[].cask_template` points to an existing
  template file that contains `depends_on cask: "graphthulhu"`

## MODIFIED Requirements

### Requirement: unbound GoReleaser cask configuration

Previously: the `homebrew_casks` stanza in `.goreleaser.yaml` used
the default GoReleaser-generated cask with no `depends_on` field.

Updated: the `homebrew_casks` stanza MUST include a `cask_template`
field referencing the project-local template file, so that every
automated release produces a cask with the graphthulhu dependency.

## REMOVED Requirements

None.
