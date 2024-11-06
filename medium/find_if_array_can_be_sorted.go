package main

import (
	"math"
	"math/bits"
)

func canSortArray(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		minIndex := -1
		minValue := math.MaxInt
		for j := i; j < len(nums); j++ {
			n2 := nums[j]

			if n1 > n2 && isSwappable(n1, n2) {
				minIndex = j
				minValue = int(math.Min(float64(minValue), float64(n2)))
			}
		}
		isValid := true
		for j := i + 1; j < minIndex; j++ {
			if !isSwappable(nums[j-1], nums[j]) {
				isValid = false
			}
		}

		if isValid && minIndex != -1 {
			nums[i] = minValue
			nums[minIndex] = n1
		}
	}

	return sortConfirm(nums)
}

func isSwappable(n1, n2 int) bool {
	return bits.OnesCount(uint(n1)) == bits.OnesCount(uint(n2))
}

func sortConfirm(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
