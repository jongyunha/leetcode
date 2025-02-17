package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= i && dp[i-1] <= i+nums[i] {
			dp[i] = i + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-2] >= len(nums)-1
}
