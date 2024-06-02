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
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(maxDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var res int

	if root == nil {
		return res
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
