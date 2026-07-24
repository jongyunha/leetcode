package main

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return countLessOrEqual(nums, upper) -
		countLessOrEqual(nums, lower-1)
}

func countLessOrEqual(nums []int, target int) int64 {
	left := 0
	right := len(nums) - 1
	var count int64
	for left < right {
		sum := nums[left] + nums[right]
		if sum <= target {
			count += int64(right - left)
			left++
		} else {
			right--
		}
	}
	return count
}
