package main

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var way int
	var dfs func(root *TreeNode, carry, targetSum int)

	dfs = func(root *TreeNode, carry, targetSum int) {
		if carry == targetSum {
			way++
		}

		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left, carry+root.Left.Val, targetSum)
		}
		if root.Right != nil {
			dfs(root.Right, carry+root.Right.Val, targetSum)
		}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Left != nil {
			queue = append(queue, curr.Left)
			nodes = append(nodes, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
			nodes = append(nodes, curr.Right)
		}
	}

	for _, node := range nodes {
		dfs(node, node.Val, targetSum)
	}

	return way
}
