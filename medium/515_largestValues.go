package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	fmt.Println(largestValues(root))
}

// largestValues 您需要在二叉树的每一行中找到最大的值。
// BFS 广度优先搜索
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	list := []*TreeNode{root}

	for len(list) > 0 {
		levelMax := math.MinInt32
		size := len(list)

		for i := 0; i < size; i++ {
			curNode := list[0]
			levelMax = int(math.Max(float64(curNode.Val), float64(levelMax)))

			list = list[1:]
			if curNode.Left != nil {
				list = append(list, curNode.Left)
			}
			if curNode.Right != nil {
				list = append(list, curNode.Right)
			}
		}

		res = append(res, levelMax)
	}

	return res

	/*res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		levelMax := math.MinInt32
		for i := 0; i < size; i++ {
			curNode := stack[0]
			stack = stack[1:]
			levelMax = int(math.Max(float64(curNode.Val), float64(levelMax)))
			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
			}
			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
			}
		}
		res = append(res, levelMax)
	}

	return res*/
}

// DFS_largestValues 深度优先搜索
func largestValuesDFS(root *TreeNode) []int {
	var helper func(root *TreeNode, level int)
	var res []int

	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, math.MinInt32)
		}
		res[level] = max(res[level], root.Val)
		helper(root.Left, 1+level)
		helper(root.Right, 1+level)
	}
	helper(root, 0)

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxNode(leftNode, rightNode *TreeNode) int {

	leftVal, rightVal := 0, 0
	if leftNode == nil && rightNode == nil {
		return 0
	}

	if leftNode == nil && rightNode != nil {
		return rightNode.Val
	}

	if leftNode != nil && rightNode == nil {
		return leftNode.Val
	}

	if leftNode.Left == nil && leftNode.Right == nil {
		leftVal = leftNode.Val
	}

	if rightNode.Left == nil && rightNode.Right == nil {
		rightVal = rightNode.Val
	}

	if leftVal < rightVal {
		return rightVal
	}

	return leftVal
}
