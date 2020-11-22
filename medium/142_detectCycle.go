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
	fmt.Println(detectCycle(head))
	fmt.Println(detectCycle2(head))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}

		m[head] = struct{}{}
		head = head.Next
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	isCycle := false
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			isCycle = true
			break
		}
	}

	if isCycle {
		slow = head
		for slow != fast {
			slow = slow.Next
			if fast != nil {
				fast = fast.Next
			}
		}
		return slow
	}

	return nil
}
