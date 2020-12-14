package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val:  20,
			Next: nil,
		},
	}
	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}

		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
