package main

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	n := len(nums)
	count := 0
	for i := 2; i < n; i++ {
		left := 0
		right := i - 1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}
	return count
}
