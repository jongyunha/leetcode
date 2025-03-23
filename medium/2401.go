package main

func longestNiceSubarray(nums []int) int {
	n := len(nums)
	maxLength := 1

	for i := 0; i < n; i++ {
		current := 0
		j := i

		for j < n && (current&nums[j]) == 0 {
			current |= nums[j]
			j++
		}

		maxLength = max(maxLength, j-i)
	}

	return maxLength
}
