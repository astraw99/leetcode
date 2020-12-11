package main

import "fmt"

func main() {
	head := &ListNode{
		Val:  10,
		Next: nil,
	}
	fmt.Println(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

type ListNode struct {
	Val  int
	Next *ListNode
}
