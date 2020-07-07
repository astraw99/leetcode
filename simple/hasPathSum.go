package main

import "fmt"

func main() {
	root := &TreeNode{10, nil, nil}
	fmt.Println(hasPathSum(root, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum [LeetCode #112]
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
// 说明: 叶子节点是指没有子节点的节点。
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	sum -= root.Val
	if hasPathSum(root.Left, sum) {
		return true
	} else {
		return hasPathSum(root.Right, sum)
	}
}
