package main

func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a][b] = true
		graph[b][a] = true
	}

	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		components := make([]int, 0)
		dfsGraph(i, graph, visited, &components)

		if isConnected(graph, components) {
			count++
		}
	}

	return count
}

func dfsGraph(node int, graph [][]bool, visited []bool, component *[]int) {
	visited[node] = true
	*component = append(*component, node)

	for neighbor := 0; neighbor < len(graph); neighbor++ {
		if graph[node][neighbor] && !visited[neighbor] {
			dfsGraph(neighbor, graph, visited, component)
		}
	}
}

func isConnected(graph [][]bool, components []int) bool {
	n := len(components)
	if n <= 1 {
		return true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !graph[components[i]][components[j]] {
				return false
			}
		}
	}

	return true
}
