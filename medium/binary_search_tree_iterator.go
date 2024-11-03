package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	queue []int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	queue := make([]int, 0)

	var inOrderTraversal func(root *TreeNode)

	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraversal(root.Left)
		queue = append(queue, root.Val)
		inOrderTraversal(root.Right)
	}

	inOrderTraversal(root)

	return BSTIterator{
		queue: queue,
	}
}

func (this *BSTIterator) Next() int {
	pop := this.queue[0]
	this.queue = this.queue[1:]

	return pop
}

func (this *BSTIterator) HasNext() bool {
	return len(this.queue) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
