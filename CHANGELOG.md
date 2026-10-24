# Changelog

All notable changes to this fork are listed here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the project uses
[Semantic Versioning](https://semver.org/).

Earlier history is in [upstream's releases](https://github.com/criteo/command-launcher/releases).

## [Unreleased]

### Added

- Minimal harness: `AGENTS.md` for agent routing and `Makefile` with
  `make check` for full verification (`build` + `test` + `integration`).
- `make install` target that builds the binary and copies it to
  `$CDT_HOME`, with an error when the variable is not set. It shows the
  currently installed version as a hint, prompts for a version (defaulting
  to `branch-dev`), and passes it to `build.sh` so it is straightforward
  to add a fork suffix (`fork`, `1.16.2-fork`, etc.) without changing
  `build.sh`.

### Fixed

- The update command compares versions numerically instead of lexicographically,
  and strips pre-release suffixes before comparison. This fixes three problems:
  - No more downgrade prompts: a locally built or side-loaded binary newer
    than the published release no longer sees an "update" to the older one.
  - Forked builds with a `-suffix` (e.g. `1.16.2-fork`, `1.16.2-alpha.1`)
    are silent while their base version matches the remote, and only prompt
    when a newer stable release ships.
  - Developer builds with bare branch names (e.g. `main-dev`) are covered
    by a string-comparison fallback when versions cannot be parsed as dotted
    numeric triples.
