/*
 * Copyright (C) 2023 https://github.com/astraw99. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License"); you may not use this
 * file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  3,
		Next: nil,
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			Val: carry % 10,
		}
		cur = cur.Next
		carry /= 10
	}

	return dummy.Next
}

// self wrong!!
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// "807"
	resStr := fmt.Sprint(traverseListNode(l1) + traverseListNode(l2))
	for i, _ := range resStr {
		nodeVal, _ := strconv.Atoi(string(resStr[i]))
		res = &ListNode{
			Val:  nodeVal,
			Next: res,
		}
	}

	return res
}

func traverseListNode(list *ListNode) int {
	if list == nil {
		return 0
	}

	tmpStr := ""
	for list != nil {
		tmpStr = fmt.Sprint(list.Val) + tmpStr
		list = list.Next
	}

	res, err := strconv.Atoi(tmpStr)
	if err != nil {
		return 0
	}

	return res
}
