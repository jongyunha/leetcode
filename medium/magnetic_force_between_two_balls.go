package main

import (
	"slices"
)

func maxDistance(position []int, m int) int {
	slices.Sort(position)
	left, right := 1, position[len(position)-1]-position[0]
	ans := -1
	for left <= right {
		mid := left + (right-left)/2

		lastPosition := position[0]
		balls := 1
		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= mid {
				lastPosition = position[i]
				balls++
			}
		}

		if balls >= m {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}

	return ans
}
