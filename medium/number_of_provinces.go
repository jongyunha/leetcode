package main

func findCircleNum(isConnected [][]int) int {
	provinces := 0
	visited := make([]bool, len(isConnected))

	var dfs func([][]int, []bool, int)
	dfs = func(isConnected [][]int, visited []bool, i int) {
		visited[i] = true
		for j := 0; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 && !visited[j] {
				dfs(isConnected, visited, j)
			}
		}
	}

	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			dfs(isConnected, visited, i)
		}
	}

	return provinces
}
