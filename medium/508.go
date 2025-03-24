package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	freqMap := make(map[int]int)
	maxFrequent := math.MinInt

	var findSubTreeSum func(root *TreeNode) int
	findSubTreeSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftSum := findSubTreeSum(root.Left)
		rightSum := findSubTreeSum(root.Right)

		subtreeSum := root.Val + leftSum + rightSum
		freqMap[subtreeSum]++

		if freqMap[subtreeSum] > maxFrequent {
			maxFrequent = freqMap[subtreeSum]
		}

		return subtreeSum
	}

	findSubTreeSum(root)
	result := make([]int, 0)

	for key, value := range freqMap {
		if value == maxFrequent {
			result = append(result, key)
		}
	}
	return result
}
