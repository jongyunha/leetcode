package main

import "math"

func SlicePop[T any](s []T, i int) ([]T, T) {
	elem := s[i]
	s = append(s[:i], s[i+1:]...)
	return s, elem
}

func findMaxIndex(s []int) int {
	var maxIndex int
	maxValue := math.MinInt
	for i := 0; i < len(s); i++ {
		if (s)[i] > maxValue {
			maxValue = (s)[i]
			maxIndex = i
		}
	}

	return maxIndex
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	treeQueue := make([]*TreeNode, 0)
	treeQueue = append(treeQueue, root)
	valueQueue := make([]int, 0)

	for len(treeQueue) != 0 {
		localTreeQueue := make([]*TreeNode, 0)
		var sum int
		for len(treeQueue) != 0 {
			var pop *TreeNode
			treeQueue, pop = SlicePop[*TreeNode](treeQueue, len(treeQueue)-1)
			sum += pop.Val
			if pop.Left != nil {
				localTreeQueue = append(localTreeQueue, pop.Left)
			}
			if pop.Right != nil {
				localTreeQueue = append(localTreeQueue, pop.Right)
			}
		}
		valueQueue = append(valueQueue, sum)
		treeQueue = localTreeQueue
	}

	return findMaxIndex(valueQueue) + 1
}
