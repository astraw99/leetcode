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

	fmt.Println(reverseBetween(head, 1, 2))
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	preM := dummy
	for i := 1; i < m; i++ {
		preM = preM.Next
	}

	cur := preM.Next
	preN := &ListNode{} // nil node
	for i := m; i <= n; i++ {
		next := cur.Next
		cur.Next = preN
		preN = cur
		cur = next
	}

	preM.Next.Next = cur
	preM.Next = preN

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
