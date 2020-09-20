package main

import "fmt"

func main() {
	s := "abcde"
	p := "a*d"

	fmt.Println(isMatch(s, p))
}

// isMatch [LeetCode #44]
// todo 本题是动态规划，没看懂？？ 请求指导 -- 20.07.05
// 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符串（包括空字符串）。
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
