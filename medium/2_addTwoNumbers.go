package main

import (
	"fmt"
	"strconv"
)

func main() {
	l1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	carry := 0

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

	return dummy.Next
}

// self wrong!!
func addTwoNumbers_(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var stack1, stack2 string
	for l1 != nil {
		stack1 = fmt.Sprint(l1.Val) + stack1
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = fmt.Sprint(l2.Val) + stack2
		l2 = l2.Next
	}

	stack1Int, _ := strconv.Atoi(stack1)
	stack2Int, _ := strconv.Atoi(stack2)
	sum := stack1Int + stack2Int
	sumStr := fmt.Sprint(sum)
	for i := len(sumStr) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  int(sumStr[i] - '0'),
			Next: nil,
		}
		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
