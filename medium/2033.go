package main

import (
	"sort"
)

func minOperations2(grid [][]int, x int) int {
	n, m := len(grid), len(grid[0])

	flatten := make([]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			flatten = append(flatten, grid[i][j])
		}
	}

	sort.Ints(flatten)

	mid := flatten[len(flatten)/2]

	count := 0
	for i := 0; i < len(flatten); i++ {
		value := flatten[i]
		if value == mid {
			continue
		}

		if value < mid {
			if (mid-value)%x != 0 {
				return -1
			}
			count += (mid - value) / x
		} else {
			if (value-mid)%x != 0 {
				return -1
			}
			count += (value - mid) / x
		}
	}

	return count
}
