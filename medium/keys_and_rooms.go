package main

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := make([]bool, len(rooms))
	visited[0] = true

	var dfs func(level int, rooms [][]int)

	dfs = func(level int, rooms [][]int) {
		keys := rooms[level]
		for _, key := range keys {
			if visited[key] {
				continue
			}

			visited[key] = true
			dfs(key, rooms)
		}
	}

	dfs(0, rooms)

	for i := range visited {
		if !visited[i] {
			return false
		}
	}

	return true
}
