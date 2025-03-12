package easy

func maximumCount(nums []int) int {
	positive := 0
	negative := 0

	for _, num := range nums {
		if num < 0 {
			negative++
		} else if num > 0 {
			positive++
		}
	}

	if positive > negative {
		return positive
	}

	return negative
}
