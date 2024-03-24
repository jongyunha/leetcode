package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		newQueue := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}

	return result
}
