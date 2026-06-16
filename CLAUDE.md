# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the CLI
go run cmd/main.go

# Run with options
go run cmd/main.go --year 2024 --calendar common
go run cmd/main.go 0 50 99.9 100

# Run tests
go test -v ./pkg/cli

# Tidy dependencies
go mod tidy
```

## Architecture

This is a Go CLI tool (module name: `uptime-calculator`) with two components:

- `cmd/main.go` — entry point; parses flags (`--year`, `--calendar`, `--debug`) and positional args (custom uptime percentiles), then delegates to `pkg/cli`
- `pkg/cli/cli.go` — all calculation logic: `Run()`, `calculate_uptime()`, `validateArgs()`, `isLeapYear()`

### Key design decisions

**Calendar modes** — `--calendar` selects year length: `gregorian` (365.2425 days, default), `tropical` (365.2422 days), or `common` (365 days). Leap years override `common` to 366 days; gregorian/tropical are fixed constants unaffected by leap years.

**Calculation flow** — `Run()` converts year length to total seconds, then `calculate_uptime()` computes `(100 - uptime%) * total_seconds / 100` and decomposes it into days/hours/minutes/seconds using `math.Mod` for the remainders. The `Downtime` struct stores all fields as `float64`; integer display happens at print time.

**Default percentiles** — `{99.0, 99.9, 99.99, 99.999, 99.9999, 99.99999}` (2–7 nines). Positional args override these; invalid args fall back to defaults with an error message printed.

**Test note** — `cli_test.go` tests `calculate_uptime()` directly with raw float comparisons against exact expected values. New test cases must provide the precise float64 values that the function returns, as there is no tolerance/delta in assertions.
