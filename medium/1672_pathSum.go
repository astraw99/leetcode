package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	sum := 6
	fmt.Println(pathSum(root, sum))
}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var path []int
	var dfs func(node *TreeNode, target int)
	dfs = func(node *TreeNode, target int) {
		target -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if node.Left == nil && node.Right == nil && target == 0 {
			res = append(res, append([]int{}, path...))
		}
		if node.Left != nil {
			dfs(node.Left, target)
		}
		if node.Right != nil {
			dfs(node.Right, target)
		}
	}

	dfs(root, sum)

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
