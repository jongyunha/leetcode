package main

import "strconv"

func canPartition(s string, target int) bool {
	if len(s) == 0 {
		return target == 0
	}
	for i := 1; i <= len(s); i++ {
		num, _ := strconv.Atoi(s[:i])
		if num <= target && canPartition(s[i:], target-num) {
			return true
		}
	}
	return false
}

func punishmentNumber(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		square := i * i
		if canPartition(strconv.Itoa(square), i) {
			sum += square
		}
	}
	return sum
}
