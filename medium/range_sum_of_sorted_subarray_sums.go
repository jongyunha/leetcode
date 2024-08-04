package main

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	if n == 0 {
		return 0
	}

	var buildPrefixSum func(nums []int) []int

	buildPrefixSum = func(nums []int) []int {
		length := len(nums)
		prefixSum := make([]int, length)

		for i := 0; i < length; i++ {
			if i == 0 {
				prefixSum[i] = nums[i]
			} else {
				prefixSum[i] = nums[i] + prefixSum[i-1]
			}
		}

		return prefixSum
	}

	totalPrefixSum := make([]int, 0)

	for i := 0; i < n; i++ {
		build := buildPrefixSum(nums[i:])
		totalPrefixSum = append(totalPrefixSum, build...)
	}

	sort.Ints(totalPrefixSum)

	var sum int64

	for i := left - 1; i < right; i++ {
		sum += int64(totalPrefixSum[i])
	}

	return int(sum % 1000000007)
}
