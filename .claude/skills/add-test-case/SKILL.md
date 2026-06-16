---
name: add-test-case
description: Add a test case to cli_test.go using exact float64 values from calculate_uptime()
---

Add a test case to pkg/cli/cli_test.go.

IMPORTANT constraints from the codebase:
- Tests use exact float64 comparison (no delta/tolerance)
- Expected values must be copied from actual calculate_uptime() output, not hand-calculated
- Run the function with debug output first, capture exact output, then use those values

Steps:
1. Run: go run cmd/main.go <uptime_percent> to see current output
2. Call calculate_uptime() with the target inputs and capture exact float64 results
3. Add the struct literal with those exact values to the tests slice in cli_test.go
4. Run: go test -v ./pkg/cli to confirm the new test passes
