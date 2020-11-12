package main

import (
	"fmt"
)

func main() {
	s := "applepenapple"
	wordDic := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDic))
}

// wordBreak 判断字符串是否可由单词组成
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}

	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
