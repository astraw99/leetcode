package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversalStack(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}

func inorderTraversalStack(root *TreeNode) []int {
	var res []int

	stack := &Stack{}
	if root != nil {
		stack.push(root)
	}

	for !stack.isEmpty() {
		cur := stack.pop()
		if cur != nil {
			if cur.Right != nil {
				stack.push(cur.Right)
			}
			stack.push(cur)
			stack.push(nil)
			if cur.Left != nil {
				stack.push(cur.Left)
			}
		} else {
			res = append(res, stack.pop().Val)
		}
	}

	return res
}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) pop() *TreeNode {
	if s.isEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v
}

func (s *Stack) push(v *TreeNode) {
	s.data = append(s.data, v)
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}
