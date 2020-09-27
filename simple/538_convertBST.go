package main

import "fmt"

func main() {
	var root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   13,
			Left:  nil,
			Right: nil,
		},
	}

	res := convertBST(root)
	fmt.Println(res.Val, res.Left.Val, res.Right.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(node *TreeNode)

	// 反序中序遍历：右中左
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)

			sum += node.Val
			node.Val = sum

			dfs(node.Left)
		}
	}

	dfs(root)

	return root
}

// not ok???
func convertBSTMy(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Val = preOrderBSTSum(root, root.Val)
	if root.Left != nil {
		//root.Left.Val = preOrderBSTSum(root.Left, root.Left.Val)
		convertBST(root.Left)
	}

	if root.Right != nil {
		//root.Right.Val = preOrderBSTSum(root.Right, root.Right.Val)
		convertBST(root.Right)
	}

	return root
}

func preOrderBSTSum(root *TreeNode, v int) (sum int) {
	if root == nil {
		return v
	}

	sum += v

	if root.Right != nil && root.Right.Val > v {
		sum += root.Right.Val
		sum += preOrderBSTSum(root.Right, root.Right.Val)
	}

	if root.Left != nil && root.Left.Val > v {
		sum += root.Left.Val
		sum += preOrderBSTSum(root.Left, root.Left.Val)
	}

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
