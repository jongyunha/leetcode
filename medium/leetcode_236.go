package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ancestor *TreeNode
	var findCommonAncestor func(root, p, q *TreeNode)
	findCommonAncestor = func(root, p, q *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			findCommonAncestor(root.Left, q, p)
		}

		if root.Right != nil {
			findCommonAncestor(root.Right, q, p)
		}

		if ancestor != nil {
			return
		}

		checker := make(map[int]*TreeNode)
		checker[root.Val] = root

		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]
			checker[pop.Val] = pop

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}

		if _, ok := checker[p.Val]; !ok {
			return
		}

		if _, ok := checker[q.Val]; !ok {
			return
		}

		ancestor = root
	}

	findCommonAncestor(root, p, q)
	return ancestor
}
