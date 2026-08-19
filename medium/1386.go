package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// row -> reserved seat bitmask
	reserved := make(map[int]int)

	for _, seat := range reservedSeats {
		row := seat[0]
		col := seat[1]

		// 1번, 10번 좌석은 가족 배치에 영향을 주지 않음
		if col == 1 || col == 10 {
			continue
		}

		reserved[row] |= 1 << col
	}

	// 예약석이 없는 모든 row는 가족 2팀 가능
	answer := n * 2

	left := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
	middle := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
	right := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

	for _, mask := range reserved {
		// 이 row는 처음에 2개 가능하다고 계산했으므로 제거
		answer -= 2

		canLeft := mask&left == 0
		canMiddle := mask&middle == 0
		canRight := mask&right == 0

		// left + right는 동시에 배치 가능
		if canLeft && canRight {
			answer += 2
		} else if canLeft || canMiddle || canRight {
			answer++
		}
	}

	return answer
}
