package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) < 2 {
		return 0
	}

	times := make([]time.Time, 0)

	for _, t := range timePoints {
		s := strings.Split(t, ":")
		hour, minute := s[0], s[1]
		hourInt, _ := strconv.Atoi(hour)
		minuteInt, _ := strconv.Atoi(minute)
		date := time.Date(2000, 1, 1, hourInt, minuteInt, 0, 0, time.UTC)
		times = append(times, date)
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].After(times[j])
	})

	minDiff := math.MaxInt

	for i := range times {
		t1 := times[i]
		t2 := times[(i+1)%len(times)]

		subMin := t1.Sub(t2).Minutes()
		if subMin < 0 {
			subMin += 60 * 24
		}

		minDiff = int(min(subMin, float64(minDiff)))
	}

	return minDiff
}
