package main

import (
	"math"
	"sort"
)

func maximumGap(nums []int) int {
	sort.Ints(nums)

	gap := 0

	for i := 0; i < len(nums)-1; i++ {
		n1, n2 := nums[i], nums[i+1]
		temp := math.Abs(float64(n1 - n2))
		gap = int(math.Max(float64(gap), temp))
	}

	return gap
}
