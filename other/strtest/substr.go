/*
 * 无重复字符的最长子串
 * 无重复字符的最长子串
 */
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	str := "abcabcbb" // abcabcbb  dvdf  pwwkew
	fmt.Println(LengthOfLongestSubstring(str))
	fmt.Println(LengthOfLongestSubstring2(str))
}

// LengthOfLongestSubstring 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if len(s) <= 1 {
		return strLen
	}

	byteSlice := []byte(s)

	maxStr := []byte{byteSlice[0]}
	// 通过滑动起止位置实现
	start := 0
	end := 1
	maxLen := 1
	for i := 1; i < strLen; i++ {
		if byteSlice[i] != byteSlice[i-1] {
			temp1 := string(byteSlice[start:end])
			temp2 := string(byteSlice[i])
			fmt.Println(temp1, temp2, start, end)

			// abcabcbb
			if strings.IndexByte(string(byteSlice[start:end]), byteSlice[i]) < 0 {
				end++
			} else {
				start = strings.LastIndexByte(string(byteSlice[:end]), byteSlice[i]) + 1
				end = i + 1
			}
		}

		if byteSlice[i] == byteSlice[i-1] {
			start = i
			end = i + 1
		}

		if end-start > maxLen {
			maxLen = end - start
			maxStr = byteSlice[start:end]
		}
	}

	fmt.Println(string(maxStr))

	return maxLen
}

// LengthOfLongestSubstring2 快速 map 写法
func LengthOfLongestSubstring2(s string) int {

	n := len(s)
	ans := 0

	m := make(map[string]int)
	start := 0
	end := 0
	for ; end < n; end++ {
		ch := string(s[end])
		val, flag := m[ch]
		if flag {
			start = int(math.Max(float64(val), float64(start)))
		}
		ans = int(math.Max(float64(ans), float64(end-start+1)))
		m[ch] = end + 1
	}
	return ans
}
