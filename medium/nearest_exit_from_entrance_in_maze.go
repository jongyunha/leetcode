package main

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	visited[entrance[0]][entrance[1]] = true
	queue := make([][]int, 0)
	queue = append(queue, entrance)

	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if (curr[0] == 0 || curr[0] == m-1 || curr[1] == 0 || curr[1] == n-1) && (curr[0] != entrance[0] || curr[1] != entrance[1]) {
				return steps
			}

			for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				x, y := curr[0]+dir[0], curr[1]+dir[1]
				if x < 0 || x >= m || y < 0 || y >= n || maze[x][y] == '+' || visited[x][y] {
					continue
				}

				visited[x][y] = true
				queue = append(queue, []int{x, y})
			}
		}
		steps++
	}

	return -1
}
