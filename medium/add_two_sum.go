package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var current *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var sum int

		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		sum += carry

		if result == nil {
			result = &ListNode{sum % 10, nil}
			current = result
		} else {
			current.Next = &ListNode{sum % 10, nil}
			current = current.Next
		}

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{carry, nil}
	}

	return result
}

func main() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	l2 := &ListNode{9, &ListNode{9, nil}}

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}