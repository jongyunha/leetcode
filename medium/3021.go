package main

func flowerGame(n int, m int) int64 {
	var result int64

	for i := 1; i <= n; i++ {
		result += int64(m / 2)
		div := m % 2
		if (div+i)%2 != 0 {
			result += int64(div)
		}
	}

	return result
}
