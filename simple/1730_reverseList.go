package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 20,
		},
	}
	fmt.Println(reverseList(head))
	fmt.Println(reverseList2(head))
}

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	for head != nil {
		temp = &ListNode{
			Val:  head.Val,
			Next: temp,
		}
		head = head.Next
	}

	return temp
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

type ListNode struct {
	Val  int
	Next *ListNode
}
