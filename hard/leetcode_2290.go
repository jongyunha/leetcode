package hard

import (
	"container/heap"
	"math"
)

type GraphNodeWithObstacle struct {
	x           int
	y           int
	val         int
	weight      int
	reverseUp   bool
	reverseLeft bool
	next        []*GraphNodeWithObstacle
}

// PriorityQueue 구현
type PriorityQueue []*GraphNodeWithObstacle

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*GraphNodeWithObstacle)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 0

	root := &GraphNodeWithObstacle{
		x:    0,
		y:    0,
		val:  grid[0][0],
		next: make([]*GraphNodeWithObstacle, 0),
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, root)
	result := math.MaxInt

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*GraphNodeWithObstacle)

		if result < node.weight || node.weight > dist[node.x][node.y] {
			continue
		}

		if node.x == m-1 && node.y == n-1 {
			result = min(result, node.weight)
		}

		if node.x < m-1 && !node.reverseUp {
			newWeight := node.weight + grid[node.x+1][node.y]
			if newWeight < dist[node.x+1][node.y] {
				next := &GraphNodeWithObstacle{
					x:      node.x + 1,
					y:      node.y,
					val:    grid[node.x+1][node.y],
					weight: newWeight,
					next:   make([]*GraphNodeWithObstacle, 0),
				}
				dist[node.x+1][node.y] = newWeight
				node.next = append(node.next, next)
				heap.Push(pq, next)
			}
		}

		if node.y < n-1 && !node.reverseLeft {
			newWeight := node.weight + grid[node.x][node.y+1]
			if newWeight < dist[node.x][node.y+1] {
				next := &GraphNodeWithObstacle{
					x:      node.x,
					y:      node.y + 1,
					val:    grid[node.x][node.y+1],
					weight: newWeight,
					next:   make([]*GraphNodeWithObstacle, 0),
				}
				dist[node.x][node.y+1] = newWeight
				node.next = append(node.next, next)
				heap.Push(pq, next)
			}
		}

		if node.x > 0 {
			newWeight := node.weight + grid[node.x-1][node.y]
			if newWeight < dist[node.x-1][node.y] {
				next := &GraphNodeWithObstacle{
					x:         node.x - 1,
					y:         node.y,
					val:       grid[node.x-1][node.y],
					weight:    newWeight,
					next:      make([]*GraphNodeWithObstacle, 0),
					reverseUp: true,
				}
				dist[node.x-1][node.y] = newWeight
				node.next = append(node.next, next)
				heap.Push(pq, next)
			}
		}

		if node.y > 0 {
			newWeight := node.weight + grid[node.x][node.y-1]
			if newWeight < dist[node.x][node.y-1] {
				next := &GraphNodeWithObstacle{
					x:           node.x,
					y:           node.y - 1,
					val:         grid[node.x][node.y-1],
					weight:      newWeight,
					next:        make([]*GraphNodeWithObstacle, 0),
					reverseLeft: true,
				}
				dist[node.x][node.y-1] = newWeight
				node.next = append(node.next, next)
				heap.Push(pq, next)
			}
		}
	}

	return result
}
