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
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversalStack(root))
}

// root left right
func preorderTraversal(root *TreeNode) []int {
	var res []int

	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

// root left right
func preorderTraversalStack(root *TreeNode) []int {
	var res []int
	var stack = &Stack{}

	if root != nil {
		stack.push(root)
	}

	for !stack.isEmpty() {
		cur := stack.pop()
		if cur != nil {
			if cur.Right != nil {
				stack.push(cur.Right)
			}
			if cur.Left != nil {
				stack.push(cur.Left)
			}
			stack.push(cur)
			stack.push(nil)
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
