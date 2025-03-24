package main

import "container/heap"

type daysMeeting struct {
	start int
	end   int
}

type daysPriorityQueue []daysMeeting

// Len returns the length of the priority queue
func (pq daysPriorityQueue) Len() int {
	return len(pq)
}

// Less defines the ordering of elements
// We prioritize meetings with earlier end times
func (pq daysPriorityQueue) Less(i, j int) bool {
	if pq[i].start == pq[j].start {
		return pq[i].end < pq[j].end
	}
	return pq[i].start < pq[j].start
}

// Swap swaps the elements with indexes i and j
func (pq daysPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an element to the priority queue
func (pq *daysPriorityQueue) Push(x interface{}) {
	meeting := x.(daysMeeting)
	*pq = append(*pq, meeting)
}

// Pop removes and returns the minimum element (according to Less)
func (pq *daysPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	meeting := old[n-1]
	*pq = old[0 : n-1]
	return meeting
}

func (pq *daysPriorityQueue) PopFirst() daysMeeting {
	if len(*pq) == 0 {
		panic("Priority queue is empty")
	}

	// Get the first element
	first := (*pq)[0]

	// If there's only one element, clear the queue
	if len(*pq) == 1 {
		*pq = (*pq)[:0]
		return first
	}

	// Replace first element with the last element
	n := len(*pq)
	(*pq)[0] = (*pq)[n-1]
	*pq = (*pq)[:n-1]

	// Restore heap property by pushing the new first element down
	if len(*pq) > 1 {
		heap.Fix(pq, 0)
	}

	return first
}

func countDays(days int, meetings [][]int) int {
	queue := make(daysPriorityQueue, 0)
	heap.Init(&queue)

	for i := 0; i < len(meetings); i++ {
		heap.Push(&queue, daysMeeting{meetings[i][0], meetings[i][1]})
	}
	count := 0
	endTime := 0
	for queue.Len() > 0 {
		meet := queue.PopFirst()
		if meet.end < endTime {
			continue
		}

		if meet.start > endTime {
			count = count + meet.start - endTime - 1
		}
		endTime = meet.end
	}

	if days > endTime {
		count = count + days - endTime
	}
	return count
}
