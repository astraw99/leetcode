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

import "fmt"

func main() {
	headA, headB := &ListNode{}, &ListNode{}
	fmt.Println(getIntersectionNode(headA, headB))
	fmt.Println(getIntersectionNode2(headA, headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}

// hash method: also right.
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	var tmp = make(map[*ListNode]bool)
	for headA != nil {
		tmp[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if tmp[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
