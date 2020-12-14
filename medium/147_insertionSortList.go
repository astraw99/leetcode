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
	fmt.Println(insertionSortList(head))
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	sorted, cur := head, head.Next
	for cur != nil {
		if sorted.Val <= cur.Val {
			sorted = sorted.Next
		} else {
			pre := dummy
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}

			// insert cur
			sorted.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
		}

		cur = sorted.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
