---
description: Run the onboarding interview to create or update your user profile.
---

# Command: /onboard

## User Input

```text
$ARGUMENTS
```

## Description

Run the onboarding interview to create or update your
user profile at `.uf/onboarding/profile.md`. The profile
captures your inspiration, interests, and objectives,
then maps them to existing Unbound Force heroes so the
entire swarm works in alignment with your goals.

## Execution

1. **Delegate to the onboarding agent** using the Task
   tool with `subagent_type: "onboarding"`.

2. **Construct the prompt** for the agent:
   - If `$ARGUMENTS` is not empty, pass it as context:
     "The user invoked /onboard with the following
     input: $ARGUMENTS. Use this as initial context
     for the interview."
   - If `$ARGUMENTS` is empty, use: "Conduct the
     onboarding interview. Check if a profile already
     exists at .uf/onboarding/profile.md — if so,
     offer to update it rather than starting from
     scratch."

3. **Expected outcome**: The agent produces or updates
   a user profile at `.uf/onboarding/profile.md` with
   YAML frontmatter (version, user, status, timestamps)
   and four Markdown sections (Inspiration, Interests,
   Objectives, Hero Mapping).

## Guardrails

- This command MUST NOT modify any existing hero agent
  files.
- This command MUST NOT create new agents that duplicate
  existing hero capabilities.
- The onboarding agent routes users to existing heroes
  for any objective that matches an existing capability.
