## ADDED Requirements

### Requirement: Snoopy regression test

The scaffold engine test suite MUST include a
`TestScaffoldOutput_NoSnoopyLegacyReferences` (or
equivalent) regression test that verifies no stale
"pinkman" strings remain in scaffold asset output.

#### Scenario: Stale Pinkman reference detected

- **GIVEN** the scaffold engine has been updated with
  renamed assets
- **WHEN** the test scans all embedded asset content
  for the pattern `pinkman` (case-insensitive)
- **THEN** zero matches are found, confirming the
  rename is complete

## MODIFIED Requirements

### Requirement: Agent file identity

The OSS Scout agent file MUST use "Snoopy" as the
hero identity name.

Previously: The agent file used "Pinkman" as the hero
identity name in the `title`, role description, persona
section, and all self-references.

#### Scenario: Agent file loaded by OpenCode

- **GIVEN** the file `.opencode/agents/snoopy.md`
  exists
- **WHEN** OpenCode loads the agent roster
- **THEN** the agent is registered with
  `subagent_type: "snoopy"` and all identity references
  use "Snoopy"

### Requirement: Scaffold asset identity

The scaffold engine MUST embed and deploy
`opencode/agents/snoopy.md` (not `pinkman.md`).

Previously: The scaffold engine embedded and deployed
`opencode/agents/pinkman.md`.

#### Scenario: uf init deploys Snoopy agent

- **GIVEN** a fresh project directory
- **WHEN** `uf init` is run
- **THEN** `.opencode/agents/snoopy.md` is created
  and `.opencode/agents/pinkman.md` is NOT created

### Requirement: Slash command routing

The `/scout` command file MUST route to
`subagent_type: "snoopy"`.

Previously: The `/scout` command routed to
`subagent_type: "pinkman"`.

#### Scenario: User invokes /scout

- **GIVEN** the `/scout` command is registered
- **WHEN** the user runs `/scout discover go cli`
- **THEN** the command dispatches to the `snoopy`
  subagent

### Requirement: Dewey learning tags

The OSS Scout agent MUST tag stored learnings with
`snoopy-{mode}` patterns (e.g., `snoopy-discover`,
`snoopy-trend`, `snoopy-audit`, `snoopy-report`).

Previously: Learnings were tagged with `pinkman-{mode}`
patterns.

#### Scenario: Scouting report stored in Dewey

- **GIVEN** a scouting operation completes
- **WHEN** the agent stores a learning in Dewey
- **THEN** the learning tag uses the `snoopy-` prefix

### Requirement: Data directory paths

The agent MUST write scouting reports to
`.uf/snoopy/reports/` and the `.gitignore` MUST
exclude `.uf/snoopy/` from version control.

Previously: Reports were written to
`.uf/pinkman/reports/` with `.uf/pinkman/` in
`.gitignore`.

#### Scenario: Report file written

- **GIVEN** the agent produces a scouting report
- **WHEN** the report is written to disk
- **THEN** the path is `.uf/snoopy/reports/{filename}`

### Requirement: Artifact envelope producer

The agent MUST use `producer: snoopy` in artifact
envelope metadata.

Previously: `producer: pinkman`.

#### Scenario: Artifact envelope produced

- **GIVEN** the agent produces an artifact envelope
- **WHEN** the envelope is serialized
- **THEN** the `producer` field value is `snoopy`

### Requirement: Documentation references

`AGENTS.md` and `unbound-force.md` MUST reference
"Snoopy" as the OSS Scout identity.

Previously: Both files referenced "Pinkman".

#### Scenario: AGENTS.md Utility Agents table

- **GIVEN** `AGENTS.md` contains the Utility Agents
  table
- **WHEN** the table is read
- **THEN** the OSS Scout row references `snoopy.md`
  and "Snoopy"

### Requirement: Go test assertions

Test assertions referencing `pinkman.md` in
`expectedAssetPaths`, `isToolOwned` test cases, and
file count comments MUST be updated to reference
`snoopy.md`.

Previously: Tests referenced `pinkman.md`.

#### Scenario: Scaffold test validates asset list

- **GIVEN** the scaffold test runs
- **WHEN** `expectedAssetPaths` is checked
- **THEN** it contains `"opencode/agents/snoopy.md"`
  and does not contain `"opencode/agents/pinkman.md"`

## REMOVED Requirements

None. All existing OSS Scout capabilities are preserved
under the new identity.
