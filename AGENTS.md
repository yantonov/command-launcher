# AGENTS.md

## Project
Command Launcher — a CLI launcher that manages and auto-updates CLI apps from remote registries.
Stack: Go 1.23+, Cobra + Viper, no database.

## Commands
- Build with custom args: `./build.sh [version] [app-name] [app-long-name]`
- Run one integration test: `./test/integration.sh <test-name>`
- **Full verification: `make check`** (also lists all targets: `make` or `make help`)

## Definition of Done
A change is done only when `make check` is green. That means built, unit-tested,
and integration-tested against the binary just built.

## Hard constraints
1. Keep existing integration tests passing. If a test fails, the change is broken.
2. Match the conventions of the package you are editing: logrus for logging,
   testify/assert for tests, Cobra `RunE` for commands. Read two call sites in the
   same package before writing one.

## Work rules
- Priority order: correctness → performance → style.
- Binary metadata comes from `-ldflags`: `main.version`, `main.appName`,
  `main.appLongName`, `main.buildNum`.

## Details
- `README.md` — project overview, features, release process
- `CHANGELOG.md` — fork-specific changes
- `Makefile` — all build/check targets
