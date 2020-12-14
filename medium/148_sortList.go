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
	//fmt.Println(sortList(head)) // merge sort
	fmt.Println(quickSortList(head)) // quick sort
}

// quick sort list
func quickSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	quickSort(head, nil)
	return head
}

func quickSort(head, tail *ListNode) {
	if head == tail || head.Next == tail {
		return
	}

	pivot := head.Val
	slow, fast := head, head.Next

	for fast != tail {
		if fast.Val <= pivot {
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}

		fast = fast.Next
	}

	slow.Val, head.Val = head.Val, slow.Val
	quickSort(head, slow)
	quickSort(slow.Next, tail)
}

// merge sort list
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	rightHead := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(rightHead)

	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}

		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
