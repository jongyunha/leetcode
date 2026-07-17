package main

import (
	"math"
	"sort"
)

func gcdSum(nums []int) int64 {
	prefix := make([]int, len(nums))
	mxi := nums[0]
	gcd(2, 6)

	for i, num := range nums {
		mxi = int(math.Max(float64(mxi), float64(num)))
		g := gcd(mxi, num)
		prefix[i] = g
	}

	sort.Ints(prefix)
	sum := int64(0)
	for len(prefix) > 1 {
		first := prefix[0]
		prefix = prefix[1:]
		last := prefix[len(prefix)-1]
		prefix = prefix[:len(prefix)-1]
		sum += int64(gcd(first, last))
	}

	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
