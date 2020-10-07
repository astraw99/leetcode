package main

import (
	"fmt"
)

func main() {
	A := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	B := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(isSubStructure(A, B))    // dfs
	fmt.Println(isSubStructureBfs(A, B)) // bfs
}

func isSubStructureBfs(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var res bool
	var queue []*TreeNode

	queue = append(queue, A)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Val == B.Val {
			if dfs(cur, B) {
				return true
			}
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return res
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var res bool
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	if A.Val == B.Val {
		res = dfs(A, B)
	}
	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)
	}

	return res
}

func dfs(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
