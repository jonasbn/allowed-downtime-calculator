package cli

import (
	"fmt"
	"testing"
)

func TestCalculateUptime(t *testing.T) {

	tests := []struct {
		uptime              float64
		totalSecondsInAYear float64
		expectedDowntime    Downtime
	}{
		{
			uptime:              99.0,
			totalSecondsInAYear: 365 * HoursInDay * MinutesInHour * SecondsInMinute,
			expectedDowntime: Downtime{
				Days:    3.65,
				Hours:   15.6,
				Minutes: 36,
				Seconds: 0,
			},
		},
		{
			uptime:              99.9,
			totalSecondsInAYear: 365 * HoursInDay * MinutesInHour * SecondsInMinute,
			expectedDowntime: Downtime{
				Days:    0.36499999999997923,
				Hours:   8.759999999999502,
				Minutes: 45.59999999997011,
				Seconds: 35.99999999820648,
			},
		},
		{
			uptime:              99.99,
			totalSecondsInAYear: 366 * HoursInDay * MinutesInHour * SecondsInMinute,
			expectedDowntime: Downtime{
				Days:    0.03660000000001872,
				Hours:   0.8784000000004494,
				Minutes: 52.704000000026966,
				Seconds: 42.24000000161777,
			},
		},
		{
			uptime:              99.999,
			totalSecondsInAYear: 365 * HoursInDay * MinutesInHour * SecondsInMinute,
			expectedDowntime: Downtime{
				Days:    0.0036500000000174284,
				Hours:   0.08760000000041827,
				Minutes: 5.256000000025097,
				Seconds: 15.360000001505796,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("uptime: %f", tt.uptime), func(t *testing.T) {
			got := calculate_uptime(tt.uptime, tt.totalSecondsInAYear)
			if got != tt.expectedDowntime {
				t.Errorf("calculate_uptime() = %v, want %v", got, tt.expectedDowntime)
			}
		})
	}
}
