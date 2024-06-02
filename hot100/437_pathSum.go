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
		Val: 10,
	}
	targetSum := 10
	fmt.Println(pathSum(root, targetSum))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	if root == nil {
		return res
	}

	res = rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func rootSum(root *TreeNode, targetSum int) int {
	var res int
	if root == nil {
		return res
	}
	if root.Val == targetSum {
		res++
	}

	if root.Left != nil {
		res += rootSum(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		res += rootSum(root.Right, targetSum-root.Val)
	}

	return res
}
