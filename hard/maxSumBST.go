package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  9,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(maxSumBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

// maxSumBST [LeetCode #1373]
// 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
// 二叉搜索树的定义如下：
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	validBST(root)
	return max
}

// 没看懂？？ 请求指导 -- 20.06.27
func validBST(curNode *TreeNode) (int, int, bool, int) {
	ret := true
	sum := curNode.Val
	left, right := curNode.Val, curNode.Val

	if curNode.Left != nil {
		left1, right1, ok, sum1 := validBST(curNode.Left)
		if !ok || left1 >= curNode.Val || right1 >= curNode.Val {
			ret = false
		}
		left = left1
		sum += sum1
	}
	if curNode.Right != nil {
		left1, right1, ok, sum1 := validBST(curNode.Right)
		if !ok || right1 <= curNode.Val || left1 <= curNode.Val {
			ret = false
		}
		right = right1
		sum += sum1
	}

	if ret {
		max = Max(max, sum)
	}

	return left, right, ret, sum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 自己写的
func maxSumBST2(root *TreeNode) int {
	resSum := 0
	leftSum := calLeftSum(root.Left)
	rightSum := calRightSum(root.Right)

	fmt.Println(leftSum, rightSum)
	if leftSum > rightSum {
		resSum = leftSum
	} else {
		resSum = rightSum
	}

	if isBST(root) {
		resSum += root.Val
	}

	return resSum
}

// calLeftSum 左子树求和
func calLeftSum(curNode *TreeNode) (sum int) {
	for curNode != nil {
		// 叶节点
		if isBST(curNode) && curNode.Left == nil && curNode.Right == nil {
			sum += curNode.Val
			return sum
		}

		if !isBST(curNode) {
			sum = 0
			curNode = curNode.Left
		}

		if isBST(curNode) {
			sum += curNode.Val
			curNode = curNode.Left
		}
	}

	return sum
}

// calRightSum 右子树求和
func calRightSum(curNode *TreeNode) (sum int) {
	for curNode != nil {
		// 叶节点
		if isBST(curNode) && curNode.Left == nil && curNode.Right == nil {
			sum += curNode.Val
			return sum
		}

		if !isBST(curNode) {
			sum = 0
			curNode = curNode.Right
		}

		if isBST(curNode) {
			sum += curNode.Val
			curNode = curNode.Right
		}
	}

	return sum
}

// isBST 是否满足 BST
func isBST(curNode *TreeNode) bool {
	if curNode.Left != nil && curNode.Left.Val >= curNode.Val {
		return false
	}

	if curNode.Right != nil && curNode.Right.Val <= curNode.Val {
		return false
	}

	return true
}
