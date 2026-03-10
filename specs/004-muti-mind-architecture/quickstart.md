# Muti-Mind Quickstart

## Installation

Muti-Mind operates as an OpenCode agent backed by a local CLI tool and
requires the `graphthulhu` MCP server to be active.

### Single-command install (recommended)

```bash
brew install unbound-force/tap/unbound
```

This installs both the `unbound` CLI and `graphthulhu` automatically.
No separate download step needed.

### Manual install

If you manage tools outside Homebrew, install each binary separately:

```bash
# Install unbound
brew install unbound-force/tap/unbound

# Install graphthulhu (if not pulled automatically)
brew install unbound-force/tap/graphthulhu
```

### Prerequisites

1. Ensure GitHub CLI (`gh`) is authenticated if you intend to sync
   with GitHub Issues:
   ```bash
   gh auth login
   ```
2. The Muti-Mind agent is located at `.opencode/agents/muti-mind-po.md`
   and is activated automatically by OpenCode when working in a repo
   scaffolded with `unbound init`.

## Initialization

Initialize a new backlog in your project:

```bash
/muti-mind.init
```
This creates the `.muti-mind/backlog/` directory.

## Managing the Backlog

Add a new item:
```bash
/muti-mind.backlog-add --type story --title "Implement user login" --priority P2
```

View the backlog:
```bash
/muti-mind.backlog-list
```

## AI Prioritization

To let Muti-Mind analyze dependencies via the Knowledge Graph and re-score the backlog:

```bash
/muti-mind.prioritize
```

## GitHub Sync

Push your local backlog to GitHub Issues:

```bash
/muti-mind.sync-push
```

Pull new issues from GitHub into your local backlog:

```bash
/muti-mind.sync-pull
```