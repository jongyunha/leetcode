package main

import "testing"

func Test_linked_list_in_binary_tree(t *testing.T) {
	// head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
	head1 := &ListNode{Val: 4}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 8}

	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 2}
	root1.Right.Left.Left = &TreeNode{Val: 5}
	root1.Right.Left.Right = &TreeNode{Val: 8}

	testcase := []struct {
		head *ListNode ``
		root *TreeNode
		want bool
	}{
		{head1, root1, true},
	}

	for i, v := range testcase {
		got := isSubPath(v.head, v.root)
		if got != v.want {
			t.Errorf("Test case #%v: isSubPath(%v, %v) = %v, want %v", i, v.head, v.root, got, v.want)
		}
	}
}
