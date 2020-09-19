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
	fmt.Println(postorderTraversal(root))
	fmt.Println(postorderTraversalStack(root))
}

// left right root
func postorderTraversal(root *TreeNode) []int {
	var res []int

	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}

	return res
}

// left right root
func postorderTraversalStack(root *TreeNode) []int {
	var res []int
	var stack = &Stack{}

	if root != nil {
		stack.push(root)
	}

	for !stack.isEmpty() {
		cur := stack.pop()
		if cur != nil {
			stack.push(cur)
			stack.push(nil)

			if cur.Right != nil {
				stack.push(cur.Right)
			}

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

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) push(v *TreeNode) {
	s.data = append(s.data, v)
}

func (s *Stack) pop() *TreeNode {
	if s.isEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
