package main

type DAGEdge struct {
	curr              int
	directions        map[int][]*DAGEdge
	reverseDirections map[int][]*DAGEdge
}

func findChampion(n int, edges [][]int) int {
	if n > 1 && len(edges) == 0 {
		return -1
	}
	if n == 1 && len(edges) == 0 {
		return 0
	}

	edgeMap := make(map[int]*DAGEdge)

	for _, edge := range edges {
		origin := edge[0]
		destination := edge[1]
		if _, ok := edgeMap[origin]; !ok {
			newEdge := &DAGEdge{
				curr:              origin,
				directions:        make(map[int][]*DAGEdge),
				reverseDirections: make(map[int][]*DAGEdge),
			}
			edgeMap[origin] = newEdge
		}

		if _, ok := edgeMap[destination]; !ok {
			newEdge := &DAGEdge{
				curr:              destination,
				directions:        make(map[int][]*DAGEdge),
				reverseDirections: make(map[int][]*DAGEdge),
			}
			edgeMap[destination] = newEdge
		}

		originEdge := edgeMap[origin]
		destinationEdge := edgeMap[destination]

		if originEdge.directions[destinationEdge.curr] == nil {
			originEdge.directions[destinationEdge.curr] = make([]*DAGEdge, 0)
		}

		if destinationEdge.reverseDirections[originEdge.curr] == nil {
			destinationEdge.reverseDirections[originEdge.curr] = make([]*DAGEdge, 0)
		}

		originEdge.directions[destinationEdge.curr] = append(originEdge.directions[destinationEdge.curr], destinationEdge)
		destinationEdge.reverseDirections[originEdge.curr] = append(destinationEdge.reverseDirections[originEdge.curr], originEdge)
	}

	if len(edgeMap) < n {
		return -1
	}

	championCount := 0
	champion := -1

	for _, edge := range edgeMap {
		if len(edge.reverseDirections) != 0 {
			continue
		}
		championCount++
		if championCount > 1 {
			return -1
		}
		champion = edge.curr
	}

	return champion
}
