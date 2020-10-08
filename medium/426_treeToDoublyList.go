package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(treeToDoublyList(root))
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var pre *TreeNode
	inorderDfs(root, &pre)

	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}

	head.Left = tail
	tail.Right = head

	return head
}

func inorderDfs(root *TreeNode, pre **TreeNode) {
	if root != nil {
		inorderDfs(root.Left, pre)

		if *pre != nil {
			(*pre).Right = root
			root.Left = *pre
		}

		*pre = root

		inorderDfs(root.Right, pre)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
