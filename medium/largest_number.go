package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})

	if nums[0] == 0 {
		return "0"
	}

	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}

	return sb.String()
}

func compare(a, b int) bool {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	return aStr+bStr > bStr+aStr
}
