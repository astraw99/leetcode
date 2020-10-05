package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 20,
			Next: &ListNode{
				Val:  30,
				Next: nil,
			},
		},
	}
	fmt.Println(swapPairsRecursive(head))
	fmt.Println(swapPairs(head))
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	preNode := dummy
	for head != nil && head.Next != nil {
		firstNode := head
		secondNode := head.Next

		preNode.Next = secondNode

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		preNode = firstNode
		head = firstNode.Next
	}

	return dummy.Next
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairsRecursive(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
