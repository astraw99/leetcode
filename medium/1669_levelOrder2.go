package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   22,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	fmt.Println(levelOrder2(root))
	fmt.Println(levelOrder2_(root))
}

func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var levelData []int
		var tempQueue []*TreeNode

		for i, _ := range queue {
			cur := queue[i]
			levelData = append(levelData, cur.Val)
			if cur.Left != nil {
				tempQueue = append(tempQueue, cur.Left)
			}
			if cur.Right != nil {
				tempQueue = append(tempQueue, cur.Right)
			}
		}

		queue = tempQueue
		res = append(res, levelData)
	}

	return res
}

// other's code
var res [][]int

func levelOrder2_(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res = make([][]int, 1)
	res[0] = []int{root.Val}

	if root.Left != nil {
		subFunc(root.Left, 1)
	}
	if root.Right != nil {
		subFunc(root.Right, 1)
	}

	return res
}

func subFunc(root *TreeNode, depth int) {
	fmt.Println(res)
	if depth >= len(res) {
		res = append(res, []int{root.Val})
	} else {
		res[depth] = append(res[depth], root.Val)
	}
	if root.Left != nil {
		subFunc(root.Left, depth+1)
	}
	if root.Right != nil {
		subFunc(root.Right, depth+1)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
