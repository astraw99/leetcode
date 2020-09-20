package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	//nums := []int{1, 2, 3, 4}
	fmt.Println(sortedArrayToBST(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST [LeetCode #108]
// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
func sortedArrayToBST(nums []int) *TreeNode {
	return createBst(nums, 0, len(nums)-1)
}

// createBst 构建左右平衡子树
func createBst(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (right + left) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = createBst(nums, left, mid-1)
	root.Right = createBst(nums, mid+1, right)

	return root
}
