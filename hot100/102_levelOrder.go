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
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrderDFS(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS (Breadth First Search)
func levelOrderBFS(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		var levelVal []int
		levelQ := q
		var tmpQ []*TreeNode
		for _, v := range levelQ {
			levelVal = append(levelVal, v.Val)
			if v.Left != nil {
				tmpQ = append(tmpQ, v.Left)
			}
			if v.Right != nil {
				tmpQ = append(tmpQ, v.Right)
			}
		}

		res = append(res, levelVal)
		q = tmpQ
	}

	return res
}

// DFS (Depth First Search)
func levelOrderDFS(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	dfs := func(root *TreeNode, level int) {}
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
