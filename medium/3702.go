package main

func longestSubsequence(nums []int) int {
	xor := 0
	hasNonZero := false
	for _, num := range nums {
		xor ^= num
		if num != 0 {
			hasNonZero = true
		}
	}
	if !hasNonZero {
		return 0
	}
	if xor != 0 {
		return len(nums)
	}
	return len(nums) - 1
}
