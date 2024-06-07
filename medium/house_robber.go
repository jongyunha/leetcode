package main

import "math"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}

	return dp[len(dp)-1]
}
