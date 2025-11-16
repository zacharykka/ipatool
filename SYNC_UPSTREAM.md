# Upstream Sync Guide

This repository is a personal fork of `majd/ipatool` with additional features
(such as `IPATOOL_PROFILE`-based multi-account support). This document
describes how to keep your fork in sync with upstream while preserving your
custom changes.

## 1. One-time upstream remote setup

Run these commands once in your local clone:

```bash
cd /Users/shan/Developer/github/ipatool

# Add upstream remote pointing to the original project
git remote add upstream https://github.com/majd/ipatool.git

# Verify remotes
git remote -v
```

You should see at least:

- `origin` → your fork (e.g. `github.com:zacharykka/ipatool.git`)
- `upstream` → `https://github.com/majd/ipatool.git`

## 2. Regularly syncing with upstream

When upstream has new changes, update your local `main` as follows:

```bash
cd /Users/shan/Developer/github/ipatool

# Fetch latest upstream branches
git fetch upstream

# Switch to local main branch
git checkout main
```

### Option A: Rebase onto upstream (recommended)

Keeps history linear and makes your fork look like "upstream + your patches":

```bash
# Rebase local main on top of upstream/main
git rebase upstream/main

# If conflicts appear, resolve them, then
# git add <file>
# git rebase --continue

# After a successful rebase, update your fork
git push origin main --force-with-lease
```

### Option B: Merge upstream/main into main

If you prefer to avoid rewriting history:

```bash
# Merge upstream changes
git merge upstream/main

# Resolve any conflicts, then
# git add <file>
# git commit

# Push merged main to your fork
git push origin main
```

## 3. Files that may conflict

Custom multi-account support is intentionally isolated, but a few files may
conflict when upstream changes them:

- `cmd/common.go`
  - Custom logic wraps the base keychain with multi-account support:
    ```go
    base := keychain.New(keychain.Args{Keyring: ring})
    return multiaccount.NewProfileKeychain(base)
    ```
  - When resolving conflicts, keep upstream changes to `newKeychain` but ensure
    the final function still returns a wrapped keychain as above.
- `README.md`
  - This fork adds a **"Multi-account usage (fork only)"** section.
  - When upstream updates the README, keep their changes and reinsert this
    section in an appropriate place (e.g. near the Usage/Compiling sections).
- `pkg/multiaccount/*`
  - These files only exist in the fork. Upstream does not touch them, so any
    conflicts here are unlikely. If they do appear, prefer your forked
    implementation.

## 4. Typical workflow summary

1. `git fetch upstream`
2. `git checkout main`
3. `git rebase upstream/main` (or `git merge upstream/main`)
4. Resolve conflicts, paying special attention to:
   - `cmd/common.go`
   - `README.md`
   - `pkg/multiaccount/*`
5. Run tests as needed, e.g. `go test ./...`
6. `git push origin main` (or `git push origin main --force-with-lease` when
   using rebase)

This keeps your fork up-to-date with upstream while preserving the
multi-account feature and any other customizations.
