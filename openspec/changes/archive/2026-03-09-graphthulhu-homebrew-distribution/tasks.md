## 1. Fix OpenSpec Schema

- [x] 1.1 Add `version: 1` to `openspec/schemas/unbound-force/schema.yaml`
- [x] 1.2 Verify schema validates: `openspec schema validate unbound-force`

## 2. graphthulhu Homebrew Cask (unbound-force/homebrew-tap)

- [x] 2.1 Clone or checkout `unbound-force/homebrew-tap` locally
- [x] 2.2 Create `Casks/graphthulhu.rb` with v0.4.0 checksums for all
      four platforms (darwin_amd64, darwin_arm64, linux_amd64, linux_arm64)
- [x] 2.3 Run `brew audit --cask Casks/graphthulhu.rb` to lint the cask
- [x] 2.4 Commit and push `Casks/graphthulhu.rb` to `homebrew-tap` main

## 3. GoReleaser Cask Template (unbound-force/unbound-force)

- [x] 3.1 Add `dependencies: [{cask: graphthulhu}]` to `homebrew_casks`
      stanza in `.goreleaser.yaml` (GoReleaser v2 native field — no
      template file needed)
- [x] 3.2 Add `hooks.post.install` xattr stanza to remove macOS
      quarantine bit on unsigned binary
- [x] 3.3 Run `goreleaser release --snapshot --clean` to verify the
      config renders correctly without errors
- [x] 3.4 Inspect the generated cask in `dist/` to confirm
      `depends_on cask: ["graphthulhu"]` is present

## 4. Verification

- [x] 4.1 Run `brew install --cask unbound-force/tap/graphthulhu` on a
      clean machine (or CI) to confirm the cask installs successfully
- [x] 4.2 Confirm `graphthulhu --version` works after install
- [ ] 4.3 Verify `brew install unbound-force/tap/unbound` pulls
      graphthulhu as a dependency (post next `unbound` release)
- [x] 4.4 Confirm constitution alignment: Composability First PASS —
      single install command produces a fully functional environment

## 5. Upstream Contribution (best-effort)

- [x] 5.1 Fork `skridlevsky/graphthulhu` and create a branch
      `feat/homebrew-tap`
- [x] 5.2 Add `Casks/graphthulhu.rb` (same content as task 2.2) to
      the fork
- [x] 5.3 Open a PR to `skridlevsky/graphthulhu` with a description
      explaining the tap pattern and offering to transfer cask ownership
      once they have their own tap

## 6. Documentation

- [x] 6.1 Update `specs/004-muti-mind-architecture/quickstart.md`
      to reflect the single-command install
- [x] 6.2 Update `AGENTS.md` Recent Changes section with this change
