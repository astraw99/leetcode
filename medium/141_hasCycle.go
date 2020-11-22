package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  20,
			Next: nil,
		},
	}
	fmt.Println(hasCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
