package main

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sortedByEnd := make([][]int, n)
	copy(sortedByEnd, events)

	sort.Slice(sortedByEnd, func(i, j int) bool {
		return sortedByEnd[i][1] < sortedByEnd[j][1]
	})

	maxValueTillEnd := make([]int, n)
	maxValueTillEnd[0] = sortedByEnd[0][2]

	for i := 1; i < n; i++ {
		maxValueTillEnd[i] = max(maxValueTillEnd[i-1], sortedByEnd[i][2])
	}

	result := 0
	for _, event := range events {
		startTime := event[0]
		currentValue := event[2]

		idx := sort.Search(n, func(i int) bool {
			return sortedByEnd[i][1] >= startTime
		})

		if idx > 0 {
			currentValue += maxValueTillEnd[idx-1]
		}

		result = max(result, currentValue)
	}

	return result
}
