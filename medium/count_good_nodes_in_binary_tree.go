package main

import "math"

func goodNodes(root *TreeNode) int {
	var count int
	return countGoodNodes(root, math.MinInt, &count)
}

func countGoodNodes(root *TreeNode, maxValue int, count *int) int {
	if root == nil {
		return *count
	}
	if maxValue < root.Val {
		maxValue = root.Val
	}
	countGoodNodes(root.Left, maxValue, count)
	countGoodNodes(root.Right, maxValue, count)

	if root.Val >= maxValue {
		*count++
	}
	return *count
}
