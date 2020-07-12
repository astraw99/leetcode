package main

import (
	"fmt"
	"math"
)

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}

	fmt.Println(calculateMinimumHP(dungeon))
}

// calculateMinimumHP 地下城游戏
// todo 看不懂题解？
// 编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
// 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。
// 我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
// 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
//
// 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；
// 其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}

	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 自己写的
func calculateMinimumHP2(dungeon [][]int) int {
	var res int
	rowLen := len(dungeon)
	colLen := len(dungeon[0])
	sum := 0

	i, j := 0, 0
	sum += dungeon[i][j]

	if rowLen == 1 && colLen == 1 {
		if sum >= 0 {
			return 1
		} else {
			return -sum + 1
		}
	}

	for i <= rowLen-1 {
		for j < colLen-1 {
			if i == rowLen-1 {
				j++
				sum += dungeon[i][j]
				continue
			}

			if i+1 == rowLen-1 {
				sum += dungeon[i+1][j]
				j++
				continue
			}

			if dungeon[i+1][j] <= dungeon[i][j+1] {
				sum += dungeon[i][j+1]
				j++
				continue
			}

			if dungeon[i+1][j] > dungeon[i][j+1] {
				sum += dungeon[i+1][j]
				break
			}
		}

		if i == rowLen-1 {
			break
		}

		i++
		if j == colLen-1 {
			sum += dungeon[i][j]
		}
	}

	switch {
	case sum <= 0:
		res = -sum + 1
	case sum > 0:
		res = sum
	}

	return res
}
