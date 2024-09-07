package main

func chalkReplacer(chalk []int, k int) int {
	total := 0
	for _, c := range chalk {
		total += c
	}

	k %= total
	for i, c := range chalk {
		if k < c {
			return i
		}

		k -= c
	}

	return -1
}
