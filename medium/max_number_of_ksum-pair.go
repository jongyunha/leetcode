package main

import "sort"

func maxOperations(nums []int, k int) int {
	left, right := 0, len(nums)-1
	var operations int
	sort.Ints(nums)

	for left < right {
		lValue := nums[left]
		rValue := nums[right]
		if lValue == -1 {
			left++
			continue
		}

		if rValue == -1 {
			right--
			continue
		}

		sumValue := lValue + rValue

		if sumValue == k && left != right {
			operations++
			nums[left] = -1
			nums[right] = -1
		} else if sumValue <= k {
			left++
		} else {
			right--
		}
	}

	return operations
}
