package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:  20,
			Left: nil,
			Right: &TreeNode{
				Val:   30,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	sum := 60
	fmt.Println(pathSum(root, sum))
}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var dfs func(node *TreeNode, left int)
	var tmpArr []int

	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}

		left -= node.Val
		tmpArr = append(tmpArr, node.Val)
		defer func() { tmpArr = tmpArr[:len(tmpArr)-1] }()

		if node.Left == nil && node.Right == nil && left == 0 {
			//res = append(res, tmpArr) // 【注意】会覆盖前面的 slice
			res = append(res, append([]int{}, tmpArr...))
			return
		}

		if node.Left != nil {
			dfs(node.Left, left)
		}
		if node.Right != nil {
			dfs(node.Right, left)
		}

		//tmpArr = tmpArr[:len(tmpArr)-1]
	}

	dfs(root, sum)

	return res
}

func calSum(root *TreeNode, tmpArr []int, sum int) (resArr, res []int) {
	var tmp int

	if root == nil {
		return tmpArr, res
	}

	if root.Left != nil {
		tmp += root.Val + root.Left.Val
		tmpArr = append(tmpArr, root.Val, root.Left.Val)
	}
	if root.Right != nil {
		tmp += root.Val + root.Right.Val
		tmpArr = append(tmpArr, root.Val, root.Right.Val)
	}

	if root.Left == nil && root.Right == nil {
		tmp += root.Val
		if tmp == sum {
			tmpArr = append(tmpArr, root.Val)
			res = tmpArr
		}
	}

	return tmpArr, res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
