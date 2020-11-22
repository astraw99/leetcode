package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(levelOrder(root))
	fmt.Println(levelOrder_(root))
}

// dfs
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	dfs := func(root *TreeNode, level int) {}
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{}) // new []int
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

// bfs
func levelOrder_(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelQ := queue
		queue = make([]*TreeNode, 0)
		var levelData []int
		for _, cur := range levelQ {
			levelData = append(levelData, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, levelData)
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
