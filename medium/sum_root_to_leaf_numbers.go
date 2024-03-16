package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return sum
	}

	return sumNumbersHelper(root.Left, sum) + sumNumbersHelper(root.Right, sum)
}
