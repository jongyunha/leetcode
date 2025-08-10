package main

func reorderedPowerOf2(n int) bool {
	countN := make(map[int]int)
	temp := n
	for temp > 0 {
		digit := temp % 10
		countN[digit]++
		temp /= 10
	}

	power := 1
	for power <= 1000000000 {
		countPower := make(map[int]int)
		temp := power
		for temp > 0 {
			digit := temp % 10
			countPower[digit]++
			temp /= 10
		}

		if mapsEqual(countN, countPower) {
			return true
		}

		power *= 2
	}

	return false
}

func mapsEqual(m1, m2 map[int]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
