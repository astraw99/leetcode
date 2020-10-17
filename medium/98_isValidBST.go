package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// self code wrong!!
func isValidBST_(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}
	}

	return dfsValidBST(root, root.Val, true) && dfsValidBST(root, root.Val, false)
}

func dfsValidBST(root *TreeNode, rootVal int, isLeft bool) bool {
	if root == nil {
		return true
	}

	if isLeft {
		if root.Left != nil {
			if root.Val <= root.Left.Val || rootVal <= root.Left.Val {
				return false
			}
		}

		if root.Right != nil {
			if root.Val >= root.Right.Val || rootVal <= root.Right.Val {
				return false
			}
		}
	} else {
		if root.Left != nil {
			if root.Val <= root.Left.Val || rootVal >= root.Left.Val {
				return false
			}
		}

		if root.Right != nil {
			if root.Val >= root.Right.Val || rootVal >= root.Right.Val {
				return false
			}
		}
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
