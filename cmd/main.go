package main

import (
	"flag"
	"time"

	"uptime-calculator/pkg/cli"
)

func main() {
	// resolve current year
	current_year := time.Now().Year()

	// Define command-line flags
	year := flag.Int("year", current_year, "specify a year or it defaults to the current year")
	flag.Parse()

	// Execute the CLI logic
	cli.Run(*year)
}
