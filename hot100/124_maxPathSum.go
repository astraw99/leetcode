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
	"math"
)

func main() {
	root := &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(maxPathSum(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt

	if root == nil {
		return maxSum
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)
		inSum := root.Val + leftSum + rightSum
		maxSum = max(maxSum, inSum)

		outSum := root.Val + max(leftSum, rightSum)
		return max(outSum, 0)
	}

	dfs(root)

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
