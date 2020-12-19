package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l2 := &ListNode{
		Val:  3,
		Next: nil,
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head

	for l1 != nil || l2 != nil || carry > 0 {
		cur.Next = new(ListNode)
		cur = cur.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Val = carry % 10
		carry /= 10
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
