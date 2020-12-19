package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(pathSum(root, 10))
}

func pathSum(root *TreeNode, sum int) int {
	var res int
	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, target int) int
	dfs = func(root *TreeNode, target int) (result int) {
		if root == nil {
			return
		}
		target -= root.Val
		if target == 0 {
			result++
		}

		if root.Left != nil {
			result += dfs(root.Left, target)
		}
		if root.Right != nil {
			result += dfs(root.Right, target)
		}

		return result
	}

	res += dfs(root, sum)
	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
