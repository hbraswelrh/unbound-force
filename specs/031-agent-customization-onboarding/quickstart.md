# Quickstart: Agent Customization Onboarding

## Prerequisites

- Unbound Force CLI installed (`brew install
  unbound-force/tap/unbound-force`)
- OpenCode configured with at least one AI model
- Run `uf init` to scaffold agent files (deploys the
  onboarding agent and `/onboard` command)

## Getting Started

### 1. Run the Onboarding Interview

Invoke the onboarding agent:

```
/onboard
```

The agent conducts a structured interview across three
dimensions:

1. **Inspiration**: What motivates you? What does
   "great" look like for your project?
2. **Interests**: What domain, technologies, and
   architectural patterns matter most?
3. **Objectives**: What concrete goals do you want
   to achieve, and by when?

The interview takes approximately 5-10 minutes. You can
answer in freeform text -- the agent asks follow-up
questions if your answers are vague.

### 2. Review Your Profile

After the interview, the agent synthesizes your inputs
and presents a summary. You can:

- **Confirm** to finalize the profile
- **Edit** any section before saving
- **Restart** a specific dimension

Your profile is saved to `.uf/onboarding/profile.md`.

### 3. Use Heroes with Your Profile

Once your profile exists, all heroes automatically
use it to contextualize their behavior:

| What You Said | How Heroes Adapt |
|--------------|-----------------|
| "I care about security" | The Divisor's Adversary persona gives extra attention to security findings |
| "Ship by Q3" | Muti-Mind prioritizes deadline-relevant backlog items higher |
| "90% test coverage" | Gaze targets the coverage threshold you specified |
| "Go and CLI tools" | Cobalt-Crush applies Go convention pack patterns |
| "Track velocity" | Mx F focuses dashboard on velocity metrics |

No additional configuration needed. If you don't run
`/onboard`, heroes operate with their default behavior
unchanged.

### 4. Update Your Profile Later

Priorities change. Re-run `/onboard` anytime:

```
/onboard
```

The agent detects your existing profile and offers to
update specific sections rather than restarting. Your
previous profile is saved to
`.uf/onboarding/history/` with a timestamp.

## What Gets Created

```text
.uf/onboarding/
├── profile.md              # Your active user profile
└── history/                # Previous profile versions
    └── 2026-04-21T14-30-00.md
```

## Committing Your Profile

Your profile is user configuration, not runtime data.
Commit it to version control:

```bash
git add .uf/onboarding/profile.md
git commit -m "chore: add user profile for agent customization"
```

This ensures your profile persists across branches and
is available in CI environments where heroes may run.

## Troubleshooting

**Q: Heroes aren't using my profile.**
Check that `.uf/onboarding/profile.md` exists and has
`status: complete` in its frontmatter. Draft profiles
may be partially used.

**Q: I want to start over completely.**
Delete `.uf/onboarding/profile.md` and run `/onboard`
again. Your old profile is preserved in the history
directory.

**Q: My objective maps to the wrong hero.**
Edit the "Hero Mapping" table in your profile manually,
or re-run `/onboard` and describe your objective
differently.
