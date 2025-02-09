package cli

import (
	"fmt"
	"math"
	"strconv"
)

const (
	HoursInDay      = 24.0
	MinutesInHour   = 60.0
	SecondsInMinute = 60.0
)

type Downtime struct {
	Days    float64
	Hours   float64
	Minutes float64
	Seconds float64
}

func validateArgs(args []string) ([]float64, error) {

	var percentiles []float64

	for _, arg := range args {
		uptime, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid parameter type: %s", arg)
		}
		if uptime < 0 || uptime > 100 {
			return nil, fmt.Errorf("invalid parameter value: %s", arg)
		}
		percentiles = append(percentiles, uptime)
	}

	return percentiles, nil
}

func Run(year int, debug bool, args []string) {

	var default_uptimes = []float64{99.0, 99.9, 99.99, 99.999, 99.9999, 99.99999}

	uptimes, err := validateArgs(args)

	if len(uptimes) == 0 {
		uptimes = default_uptimes
	}

	if err != nil {
		fmt.Printf("Error: %s\n\tContinuing with the known defaults %v\n\n", err, default_uptimes)
	}

	var number_of_days float64

	if isLeapYear(year) {
		number_of_days = 366.0
	} else {
		number_of_days = 365.0
	}

	var total_seconds_in_a_year = number_of_days * HoursInDay * MinutesInHour * SecondsInMinute

	fmt.Printf("Calculated allowed downtime for uptime requirement in year: %d (%f days):\n", year, number_of_days)
	for _, uptime := range uptimes {
		downtime := calculate_uptime(uptime, total_seconds_in_a_year)
		if !debug {
			fmt.Printf("\t%f%% is: %d days %d hours %d minutes %d seconds\n",
				uptime,
				int(downtime.Days),
				int(downtime.Hours),
				int(downtime.Minutes),
				int(downtime.Seconds))
		} else {
			fmt.Printf("\t%f%% is: %f days %f hours %f minutes %f seconds\n",
				uptime,
				downtime.Days,
				downtime.Hours,
				downtime.Minutes,
				downtime.Seconds)
		}
	}
}

func calculate_uptime(uptime float64, total_seconds_in_a_year float64) Downtime {

	var downtime = 100.0 - uptime
	var calculated_total_downtime_in_seconds = total_seconds_in_a_year * downtime / 100.0
	var days_of_downtime = calculated_total_downtime_in_seconds / (HoursInDay * MinutesInHour * SecondsInMinute)

	var remaining_seconds = math.Mod(calculated_total_downtime_in_seconds, (HoursInDay * MinutesInHour * SecondsInMinute))
	var hours_of_downtime = remaining_seconds / (MinutesInHour * SecondsInMinute)
	remaining_seconds = math.Mod(remaining_seconds, (MinutesInHour * SecondsInMinute))

	var minutes_of_downtime = remaining_seconds / SecondsInMinute
	remaining_seconds = math.Mod(remaining_seconds, SecondsInMinute)

	return Downtime{
		Days:    days_of_downtime,
		Hours:   hours_of_downtime,
		Minutes: minutes_of_downtime,
		Seconds: remaining_seconds,
	}
}

func isLeapYear(year int) bool {
	if year == 0 || year%4 > 0 {
		return false
	} else {
		if year%100 == 0 && year%400 > 0 {
			return false
		} else {
			return true
		}
	}
}
