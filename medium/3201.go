package main

// You are given an integer array nums.
// A subsequence sub of nums with length x is called valid if it satisfies:
//
// (sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2.
// Return the length of the longest valid subsequence of nums.
//
// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
//
// Example 1:
//
// Input: nums = [1,2,3,4]
//
// Output: 4
//
// Explanation:
//
// The longest valid subsequence is [1, 2, 3, 4].
//
// Example 2:
//
// Input: nums = [1,2,1,1,2,1,2]
//
// Output: 6
//
// Explanation:
//
// The longest valid subsequence is [1, 2, 1, 2, 1, 2].
//
// Example 3:
//
// Input: nums = [1,3]
//
// Output: 2
//
// Explanation:
//
// The longest valid subsequence is [1, 3].
func maximumLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	maxLength := 1

	// 각 패리티 패턴 (0: 짝수, 1: 홀수)에 대해 확인
	for targetParity := 0; targetParity <= 1; targetParity++ {
		// dp[parity] = 해당 패리티로 끝나는 부분수열의 최대 길이
		dp := [2]int{0, 0}

		for _, num := range nums {
			numParity := num % 2

			// 현재 숫자와 합했을 때 targetParity가 되는 이전 숫자의 패리티 계산
			// (prevParity + numParity) % 2 == targetParity
			// prevParity = (targetParity - numParity + 2) % 2
			prevParity := (targetParity - numParity + 2) % 2

			if dp[prevParity] > 0 {
				// 이전에 prevParity로 끝나는 부분수열이 있으면 확장
				dp[numParity] = max(dp[numParity], dp[prevParity]+1)
			} else {
				// 첫 번째 원소이거나 연결할 수 없는 경우
				dp[numParity] = max(dp[numParity], 1)
			}

			maxLength = max(maxLength, dp[numParity])
		}
	}

	return maxLength
}
