package main

import "sort"

func minSetSize(arr []int) int {
	length := len(arr)
	counter := make(map[int]int)
	for _, ele := range arr {
		counter[ele]++
	}

	counterSet := make([]int, 0, len(counter))
	for _, value := range counter {
		counterSet = append(counterSet, value)
	}

	sort.Ints(counterSet)
	answer := 0
	sum := 0

	for sum < length/2 {
		idx := findFloor(counterSet, length/2-sum)
		sum += counterSet[idx]
		answer++
		counterSet = append(counterSet[:idx], counterSet[idx+1:]...)
	}

	return answer
}

func findFloor(arr []int, target int) int {
	if arr[len(arr)-1] <= target {
		return len(arr) - 1
	}

	left := 0
	right := len(arr) - 1
	var mid int

	for left < right {
		mid = (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
