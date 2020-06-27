package main

import "fmt"

func main() {
	// [1, 2, 3, 3, 2, 1]
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(removeDuplicateNodes(head))
}

// removeDuplicateNodes 去除重复链表节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mapTemp := make(map[int]int)
	curNode := head
	for curNode != nil && curNode.Next != nil {
		fmt.Println("curVal", curNode.Val)
		mapTemp[curNode.Val] = 1

		if _, ok := mapTemp[curNode.Next.Val]; ok {
			curNode.Next = curNode.Next.Next
		} else {
			mapTemp[curNode.Next.Val] = 1
			curNode = curNode.Next
		}

		fmt.Println("cur", curNode)
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
