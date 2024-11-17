package hard

import (
	"math"
)

// Given an integer array nums and an integer k, return the length of the shortest non-empty subarray of nums with a sum of at least k. If there is no such subarray, return -1.
//
// A subarray is a contiguous part of an array.
//
// Example 1:
//
// Input: nums = [1], k = 1
// Output: 1
// Example 2:
//
// Input: nums = [1,2], k = 4
// Output: -1
// Example 3:
//
// Input: nums = [2,-1,2], k = 3
// Output: 3
// O(n^2) is time limit exceeded
// Approach: Prefix sum, sliding window, monotonic queue, binary search
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	result := math.MaxInt
	deque := make([]int, 0)
	prefixSum := make([]int, n+1)
	for i := 0; i < len(prefixSum)-1; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for i := 0; i <= n; i++ {
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= k {
			result = min(result, i-deque[0])
			deque = deque[1:]
		}
		for len(deque) > 0 && prefixSum[i] < prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	if result == math.MaxInt {
		return -1
	}

	return result
}
