package main

func uniqueXorTriplets(nums []int) int {
	n := len(nums)

	if n < 3 {
		return n
	}

	result := 1
	for result <= n {
		result = result << 1
	}

	return result
}
