package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   20,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	sum += sumOfLeftLeaves(root.Left)
	sum += sumOfLeftLeaves(root.Right)

	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
