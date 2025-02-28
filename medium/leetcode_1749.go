package main

func maxAbsoluteSum(nums []int) int {
	maxSum, minSum := 0, 0
	currentMax, currentMin := 0, 0

	for _, num := range nums {
		currentMax = max(num, currentMax+num)
		currentMin = min(num, currentMin+num)
		maxSum = max(maxSum, currentMax)
		minSum = min(minSum, currentMin)
	}

	return max(maxSum, -minSum)
}
