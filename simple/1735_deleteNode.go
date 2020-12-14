package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	fmt.Println(deleteNode(head, 5))
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			return dummy.Next
		}

		pre = head
		head = head.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
