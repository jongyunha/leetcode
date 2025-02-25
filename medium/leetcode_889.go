package main

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if len(pre) == 1 {
		return root
	}

	left := &TreeNode{Val: pre[1]}
	var L int
	for i, curr := range post {
		if curr == left.Val {
			L = i + 1
		}
	}

	root.Left = constructFromPrePost(pre[1:L+1], post[:L])
	root.Right = constructFromPrePost(pre[1+L:], post[L:len(post)-1])

	return root
}
