package main

import (
	"math"
	"sort"
)

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)

	numMap := make(map[int]bool)
	result := 0

	for _, num := range nums {
		numMap[num] = true
	}

	for i := range nums {
		curr := nums[i]
		local := 0

		for numMap[curr*curr] {
			curr = curr * curr
			local++
		}

		result = int(math.Max(float64(local), float64(result)))
	}

	if result == 0 {
		return -1
	}

	return result + 1
}
