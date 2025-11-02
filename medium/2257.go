package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	prison := make([][]int, m)
	for i := 0; i < len(prison); i++ {
		prison[i] = make([]int, n)
	}

	// 벽을 2로 표시
	for i := 0; i < len(walls); i++ {
		x, y := walls[i][0], walls[i][1]
		prison[x][y] = 2
	}

	// 가드를 3으로 표시 (1과 구분하기 위해)
	for i := 0; i < len(guards); i++ {
		x, y := guards[i][0], guards[i][1]
		prison[x][y] = 3
	}

	// 가드가 볼 수 있는 영역을 1로 표시
	for i := 0; i < len(guards); i++ {
		x, y := guards[i][0], guards[i][1]
		fillGuard(x, y, &prison, m, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if prison[i][j] == 0 {
				count++
			}
		}
	}

	return count
}

func fillGuard(currX, currY int, prison *[][]int, m, n int) {
	// 아래 방향
	for i := currX + 1; i < m; i++ {
		if (*prison)[i][currY] == 2 || (*prison)[i][currY] == 3 {
			break
		}
		(*prison)[i][currY] = 1
	}

	// 위 방향
	for i := currX - 1; i >= 0; i-- {
		if (*prison)[i][currY] == 2 || (*prison)[i][currY] == 3 {
			break
		}
		(*prison)[i][currY] = 1
	}

	// 오른쪽 방향
	for i := currY + 1; i < n; i++ {
		if (*prison)[currX][i] == 2 || (*prison)[currX][i] == 3 {
			break
		}
		(*prison)[currX][i] = 1
	}

	// 왼쪽 방향
	for i := currY - 1; i >= 0; i-- {
		if (*prison)[currX][i] == 2 || (*prison)[currX][i] == 3 {
			break
		}
		(*prison)[currX][i] = 1
	}
}
