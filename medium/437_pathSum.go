package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   -2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	sum := 8
	fmt.Println(pathSum(root, sum))
	fmt.Println(pathSum2(root, sum))
}

// 用于存放路径和
var sumCountMap map[int]int

func pathSum(root *TreeNode, sum int) int {
	// 初始化用于记录路径和map
	sumCountMap = make(map[int]int)
	// 路径为0的1条
	sumCountMap[0] = 1
	// 递归
	return dfs(root, sum, 0)
}

// 大致懂了，细节有点晕~~
// dfs + 回溯
func dfs(node *TreeNode, sum, nowSum int) int {
	// 结束条件
	if node == nil {
		return 0
	}
	var res int
	// 来到这里之后的路径和
	nowSum += node.Val
	// 当前值减去目标值，map里有多少个这样的值就表示有多少个节点到这里是目标路径
	res += sumCountMap[nowSum-sum]
	// 当前值存起来
	sumCountMap[nowSum]++
	//fmt.Println(sumCountMap, res)
	// 左右递归
	res += dfs(node.Left, sum, nowSum)
	res += dfs(node.Right, sum, nowSum)
	// 回溯，取消当前值
	sumCountMap[nowSum]--

	return res
}

func pathSum2(root *TreeNode, sum int) int {
	var res int
	//var resPath [][]int
	if root == nil {
		return res
	}

	var dfs = func(root *TreeNode, path []int, target int) {}
	dfs = func(root *TreeNode, path []int, target int) {
		target -= root.Val
		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if target == 0 {
			res++
			//resPath = append(resPath, append([]int{}, path...))
			// no return: 路径下面可重复
		}

		if root.Left != nil {
			dfs(root.Left, path, target)
		}

		if root.Right != nil {
			dfs(root.Right, path, target)
		}
	}

	dfs(root, []int{}, sum)
	res += pathSum2(root.Left, sum)
	res += pathSum2(root.Right, sum)

	//fmt.Println(resPath)

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
