package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 12,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(detectCycle(head))
}

// detectCycle 环路检测
// todo 没看懂？？
// 给定一个有环链表，实现一个算法返回环路的开头节点。
// 有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针，然后相遇后再从链表头开始
	if head == nil || head.Next == nil {
		return nil
	}

	// 用快慢指针，判定是否有环
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		// 慢指针走一步
		slow = slow.Next
		// 快指针走两步
		fast = fast.Next.Next

		// 快慢指针在相同节点相遇，那么判定指针有环
		if slow == fast {
			break
		}
	}

	// 链表无环，那么退出
	if fast != slow {
		return nil
	}

	// 快慢指针能相遇，有环，则头结点和相遇处依次往后移动一个节点，再次相遇处则为环的起始节点
	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func detectCycle2(head *ListNode) *ListNode {
	// 快慢指针，然后相遇后再从链表头开始
	if head == nil {
		return nil
	}
	if head.Next == head {
		return head
	}
	fir, sec := head, head
	flag := false

	for fir != nil && fir.Next != nil {
		fir = fir.Next.Next
		sec = sec.Next
		if fir == sec {
			flag = true
			break
		}
	}

	fir = head
	if flag {
		for fir != sec {
			sec = sec.Next
			fir = fir.Next
		}
		return fir
	}

	return nil
}
