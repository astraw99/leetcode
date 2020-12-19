package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   10,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  40,
			Left: nil,
		},
	}

	fmt.Println(isBalanced(root))
	fmt.Println(isBalanced2(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(treeDepth(root.Left)-treeDepth(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(treeDepth(root.Left), treeDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := height(root.Left)
	r := height(root.Right)
	if l > r+1 || l+1 < r {
		return false
	}
	return isBalanced2(root.Left) && isBalanced2(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	fmt.Println(root.Val)

	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
