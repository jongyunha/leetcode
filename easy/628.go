package easy

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	if nums[1] < 0 {
		a1 := nums[n-3] * nums[n-2] * nums[n-1]
		a2 := nums[1] * nums[0] * nums[n-1]

		return int(math.Max(float64(a1), float64(a2)))
	}
	return nums[n-3] * nums[n-2] * nums[n-1]
}
