package main

func rob2(nums []int) int {
	n := len(nums)

	// 특수 케이스 처리
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// 첫 번째 집을 제외한 경우와 마지막 집을 제외한 경우 중 최댓값 선택
	return max(robLinear(nums[1:]), robLinear(nums[:n-1]))
}

func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// dp[i]는 0~i 집까지 고려했을 때 최대 금액
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		// i번째 집을 털거나 털지 않거나
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}
