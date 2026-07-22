package easy

func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	answer := make([][]int, n)
	for i := range answer {
		answer[i] = make([]int, m)
	}

	k = k % (n * m)
	for i := range grid {
		for j := range grid[i] {
			x, y := calculatePosition(i, j, k, n, m)
			answer[x][y] = grid[i][j]
		}
	}

	return answer
}

func calculatePosition(currX int, currY int, k int, n, m int) (nextX, nextY int) {
	nextX = currX
	nextY = currY
	for k > 0 {
		if nextY == m-1 && nextX == n-1 {
			nextX = 0
			nextY = 0
		} else if nextY == m-1 && nextX < n-1 {
			nextY = 0
			nextX++
		} else if nextY < m-1 {
			nextY++
		}
		k--
	}
	return nextX, nextY
}
