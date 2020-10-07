package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var res int
	var mapSeq = make(map[byte]int) // {a:0,b:1,c:2}

	for i, j := 0, 0; i < len(s) && j < len(s); {
		if pos, ok := mapSeq[s[j]]; ok {
			delete(mapSeq, s[j])
			i = max(i, pos+1)
		} else {
			mapSeq[s[j]] = j
			j++
		}

		if res < j-i {
			res = j - i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
