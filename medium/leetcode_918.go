package main

// Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.
//
// A circular array means the end of the array connects to the beginning of the array. Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].
//
// A subarray may only include each element of the fixed buffer nums at most once. Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.
//
// Example 1:
//
// Input: nums = [1,-2,3,-2]
// Output: 3
// Explanation: Subarray [3] has maximum sum 3.
// Example 2:
//
// Input: nums = [5,-3,5]
// Output: 10
// Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
// Example 3:
//
// Input: nums = [-3,-2,-3]
// Output: -2
// Explanation: Subarray [-2] has maximum sum -2.
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	kadane := func(nums []int, findMax bool) int {
		curr := nums[0]
		result := nums[0]

		for i := 1; i < len(nums); i++ {
			if findMax {
				curr = max(nums[i], curr+nums[i])
				result = max(curr, result)
			} else {
				curr = min(nums[i], curr+nums[i])
				result = min(curr, result)
			}
		}
		return result
	}

	findMax := kadane(nums, true)
	if findMax < 0 {
		return findMax
	}

	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	findMin := kadane(nums, false)
	return max(findMax, totalSum-findMin)
}
