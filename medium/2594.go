package main

import (
	"math"
	"slices"
)

func repairCars(ranks []int, cars int) int64 {
	minRepairTime := 0
	maxRepairTime := slices.Max(ranks) * cars * cars

	result := int64(maxRepairTime)
	for minRepairTime < maxRepairTime {
		midTime := (minRepairTime + maxRepairTime) / 2

		repairs := 0
		for _, rank := range ranks {
			if repairs >= cars {
				break
			}
			repairs += int(math.Sqrt(float64(midTime / rank)))
		}

		if repairs >= cars {
			maxRepairTime = midTime
			if result > int64(midTime) {
				result = int64(midTime)
			}
		} else {
			minRepairTime = midTime + 1
		}
		midTime = (minRepairTime + maxRepairTime) / 2
	}

	return result
}
