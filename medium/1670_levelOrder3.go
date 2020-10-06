package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(levelOrder3(root))
}

func levelOrder3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	var level int

	queue = append(queue, root)
	for len(queue) > 0 {
		var levelData []int
		var tempQueue []*TreeNode

		for i, _ := range queue {
			cur := queue[i]
			if level%2 == 0 {
				levelData = append(levelData, cur.Val)
			} else {
				levelData = append([]int{cur.Val}, levelData...)
			}
			if cur.Left != nil {
				tempQueue = append(tempQueue, cur.Left)
			}
			if cur.Right != nil {
				tempQueue = append(tempQueue, cur.Right)
			}
		}

		level++
		queue = tempQueue
		res = append(res, levelData)
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
