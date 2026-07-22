package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maximumIndex := 0
	for i := range nums {
		if nums[i] > nums[maximumIndex] {
			maximumIndex = i
		}
	}

	root := &TreeNode{
		Val: nums[maximumIndex],
	}
	left := nums[:maximumIndex]
	right := nums[maximumIndex+1:]
	root.Left = dfsConstructMaximumBinaryTree(root, left)
	root.Right = dfsConstructMaximumBinaryTree(root, right)
	return root
}

func dfsConstructMaximumBinaryTree(parent *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maximumIndex := 0
	for i := range nums {
		if nums[i] > nums[maximumIndex] {
			maximumIndex = i
		}
	}

	left := nums[:maximumIndex]
	right := nums[maximumIndex+1:]
	curr := &TreeNode{
		Val: nums[maximumIndex],
	}

	curr.Left = dfsConstructMaximumBinaryTree(curr, left)
	curr.Right = dfsConstructMaximumBinaryTree(curr, right)
	return curr
}
