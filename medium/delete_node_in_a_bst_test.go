package main

import "testing"

func TestDeleteNodeInABst(t *testing.T) {
	tests := []struct {
		root *TreeNode
		key  int
		want *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			key: 3,
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			root: &TreeNode{Val: 0},
			key:  0,
			want: nil,
		},
	}
	for _, test := range tests {
		if got := deleteNode(test.root, test.key); !isSameTree(got, test.want) {
			t.Errorf("deleteNode(%v, %v) = %v; want %v\n", test.root, test.key, got, test.want)
		}
	}
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
