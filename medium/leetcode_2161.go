package main

func pivotArray(nums []int, pivot int) []int {
	left := make([]int, 0)
	pivots := make([]int, 0)
	right := make([]int, 0)

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			pivots = append(pivots, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(left, pivots...), right...)
}
