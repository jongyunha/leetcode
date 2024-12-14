package main

import "math"

func continuousSubarrays(nums []int) int64 {
	n := len(nums)
	var result int64
	left := 0

	window := make(map[int]int)

	for right := 0; right < n; right++ {
		window[nums[right]]++

		for getMax(window)-getMin(window) > 2 {
			window[nums[left]]--
			if window[nums[left]] == 0 {
				delete(window, nums[left])
			}
			left++
		}

		result += int64(right - left + 1)
	}

	return result
}

func getMax(m map[int]int) int {
	max := math.MinInt32
	for k := range m {
		if k > max {
			max = k
		}
	}
	return max
}

func getMin(m map[int]int) int {
	min := math.MaxInt32
	for k := range m {
		if k < min {
			min = k
		}
	}
	return min
}
