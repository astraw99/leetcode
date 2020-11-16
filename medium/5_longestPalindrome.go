package main

import "fmt"

func main() {
	s := "ababa"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	var res string
	if len(s) == 0 {
		return res
	}

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// dp[i][j]: s 的第 i 到 j 个字母组成的串是否为回文串
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && len(res) < l+1 {
				res = s[i : i+l+1]
			}
		}
	}

	return res
}

// self wrong!!
func longestPalindrome_(s string) string {
	var res string
	if len(s) == 0 {
		return res
	}

	flag := false
	for i, j := 0, len(s); i < len(s) && j >= 0; {
		if isPalindrome(s[i:j]) {
			return s[i:j]
		} else {
			if flag {
				i++
			} else {
				j--
			}
			flag = !flag
		}

		if len(res) < j-i+1 {
			res = s[i : j+1]
		}
	}

	return res
}

func isPalindrome(s string) bool {
	// fmt.Println(s)
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}
