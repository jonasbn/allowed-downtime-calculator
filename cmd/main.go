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
	calendar := flag.String("calendar", "gregorian", "specify definition of year length: gregorian, tropical or common")
	debug := flag.Bool("debug", false, "generate debug output")
	flag.Parse()

	args := flag.Args()

	// Execute the CLI logic
	cli.Run(*year, *calendar, *debug, args)
}
