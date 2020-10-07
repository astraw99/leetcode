package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12258
	fmt.Println(translateNum(num))
}

// 0:a,1:b,z:25
func translateNum(num int) int {
	//numStr := fmt.Sprintf("%d", num)
	numStr := strconv.Itoa(num)
	l := len(numStr)

	dp := make([]int, l+1) // 动态规划，dp[i] 表示以第i个数字结尾时，可转化为字符串的总数
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= l; i++ {
		dp[i] = dp[i-1]
		if numStr[i-2:i] >= "10" && numStr[i-2:i] <= "25" {
			dp[i] += dp[i-2]
		}
	}

	return dp[l]
}
