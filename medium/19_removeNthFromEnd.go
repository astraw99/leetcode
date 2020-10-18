package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val:  10,
		Next: nil,
	}
	n := 1
	fmt.Println(removeNthFromEnd(head, n))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	cur := dummy
	i := 0
	for head != nil {
		if i >= n {
			cur = cur.Next
		}

		head = head.Next
		i++
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
