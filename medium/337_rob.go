package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(rob(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	notRob, Rob := 0, 0
	Rob += root.Val
	if root.Left != nil {
		Rob += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		Rob += rob(root.Right.Left) + rob(root.Right.Right)
	}

	notRob = rob(root.Left) + rob(root.Right)

	return max(notRob, Rob)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
