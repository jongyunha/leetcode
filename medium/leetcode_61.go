package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	curr := head

	nodes := make([]*ListNode, 0)
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	n := len(nodes)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return nodes[0]
	}

	branch := nodes[len(nodes)-(k%n+1)]

	first := branch.Next
	if first == nil {
		return head
	}

	branch.Next = nil

	last := nodes[len(nodes)-1]
	last.Next = head

	return first
}
