package main

import (
	"container/heap"
	"math"
)

type MaxValueHeap []int

func (h MaxValueHeap) Len() int { return len(h) }
func (h MaxValueHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxValueHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxValueHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxValueHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	heap.Fix(h, 0)
	return x
}

func maximumSum(nums []int) int {
	sumMap := make(map[int]*MaxValueHeap)

	for _, num := range nums {
		digit := 0
		temp := num

		for temp > 0 {
			digit += temp % 10
			temp /= 10
		}

		if _, exists := sumMap[digit]; !exists {
			h := &MaxValueHeap{}
			sumMap[digit] = h
		}

		heap.Push(sumMap[digit], num)
	}

	result := math.MinInt

	for _, h := range sumMap {
		if h.Len() < 2 {
			continue
		}

		n1 := h.Pop().(int)
		n2 := h.Pop().(int)

		if result < n1+n2 {
			result = n1 + n2
		}
	}

	if result == math.MinInt {
		return -1
	}

	return result
}
