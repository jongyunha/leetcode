package main

func minReorder(n int, connections [][]int) int {
	// build a graph
	graph := make(map[int]map[int]bool)
	for _, conn := range connections {
		if _, ok := graph[conn[0]]; !ok {
			graph[conn[0]] = make(map[int]bool)
		}
		graph[conn[0]][conn[1]] = true
		if _, ok := graph[conn[1]]; !ok {
			graph[conn[1]] = make(map[int]bool)
		}
		graph[conn[1]][conn[0]] = false
	}

	// dfs
	return dfs(graph, 0, -1)
}

func dfs(graph map[int]map[int]bool, node, parent int) int {
	count := 0
	for k := range graph[node] {
		if k == parent {
			continue
		}
		if graph[node][k] {
			count++
		}
		count += dfs(graph, k, node)
	}
	return count
}
