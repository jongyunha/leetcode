package main

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]bool)
	for _, ban := range banned {
		bannedMap[ban] = true
	}

	sum := 0
	count := 0
	for i := 1; i <= n; i++ {
		if i > maxSum {
			break
		}

		if _, ok := bannedMap[i]; ok {
			continue
		}

		sum += i
		if sum > maxSum {
			break
		}
		count++
	}

	return count
}
