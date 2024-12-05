package main

import "container/heap"

type PriorityQueue []*ListNode

// Len은 heap.Interface를 구현합니다
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func sortList(head *ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	curr := head
	for curr != nil {
		heap.Push(&pq, &ListNode{Val: curr.Val})
		curr = curr.Next
	}

	if pq.Len() == 0 {
		return nil
	}

	sorted := heap.Pop(&pq).(*ListNode)
	curr = sorted
	for curr != nil && pq.Len() > 0 {
		pop := heap.Pop(&pq).(*ListNode)
		curr.Next = pop
		curr = pop
	}
	curr = nil

	return sorted
}
