package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   30,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   60,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  40,
			Left: nil,
			Right: &TreeNode{
				Val:   50,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := height(root.Left)
	r := height(root.Right)
	if l > r+1 || l+1 < r {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
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
