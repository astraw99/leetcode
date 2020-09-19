package main

import "fmt"

func main() {
	var root = &TreeNode{
		Val: 100,
		Left: &TreeNode{
			Val:  200,
			Left: nil,
			Right: &TreeNode{
				Val:   300,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}

	fmt.Println(*invertTree(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
