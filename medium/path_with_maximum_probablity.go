package main

import (
	"container/heap"
)

type Edge struct {
	node int
	prob float64
}

type MaxHeap []Edge

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make([][]Edge, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		prob := succProb[i]
		graph[a] = append(graph[a], Edge{b, prob})
		graph[b] = append(graph[b], Edge{a, prob})
	}

	maxProb := make([]float64, n)
	for i := 0; i < n; i++ {
		maxProb[i] = 0.0
	}
	maxProb[start_node] = 1.0

	maxHeap := &MaxHeap{}
	heap.Push(maxHeap, Edge{start_node, 1.0})

	for maxHeap.Len() > 0 {
		curr := heap.Pop(maxHeap).(Edge)
		currNode, currProb := curr.node, curr.prob

		if currNode == end_node {
			return currProb
		}

		for _, edge := range graph[currNode] {
			nextNode, nextProb := edge.node, edge.prob
			if currProb*nextProb > maxProb[nextNode] {
				maxProb[nextNode] = currProb * nextProb
				heap.Push(maxHeap, Edge{nextNode, maxProb[nextNode]})
			}
		}
	}

	return 0.0
}
