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
	"sort"
)

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	fmt.Println(sortList(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var res *ListNode

	if head == nil {
		return res
	}

	var tmpSlice []int
	for head != nil {
		tmpSlice = append(tmpSlice, head.Val)
		head = head.Next
	}

	sort.Ints(tmpSlice)

	for i := len(tmpSlice) - 1; i >= 0; i-- {
		res = &ListNode{
			Val:  tmpSlice[i],
			Next: res,
		}
	}

	return res
}
