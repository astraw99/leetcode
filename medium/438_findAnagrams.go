package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
	fmt.Println(findAnagrams2(s, p))
	fmt.Println(findAnagrams3(s, p))
}

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return res
	}

	l := len(p)
	for i := 0; i+l <= len(s); i++ {
		// solution 1
		/*if strAsciiSum(s[i:i+l]) == strAsciiSum(p) && isSubStr(s[i:i+l], p) {
			res = append(res, i)
		}*/

		// solution 2
		if strMultiplySum(s[i:i+l]) == strMultiplySum(p) {
			res = append(res, i)
		}
	}

	return res
}

func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	cn1 := [26]int{}
	cn2 := [26]int{}

	for i := 0; i < len(p); i++ {
		cn1[p[i]-'a']++
		cn2[s[i]-'a']++
		//fmt.Printf("%T111\n", s[i]-'a')
		//fmt.Printf("%T222\n", cn2[s[i]-'a'])
	}
	//fmt.Println("cn1", cn1)

	var res []int
	for i := 0; i <= len(s)-len(p); i++ {
		if cn1 == cn2 {
			res = append(res, i)
		}
		cn2[s[i]-'a']--

		if i+len(p) < len(s) {
			cn2[s[i+len(p)]-'a']++
		}
		//fmt.Println("cn2", cn2)
	}

	return res
}

func findAnagrams3(s string, p string) []int {
	var result []int // 用数组记录答案
	left, right := 0, 0
	// 需要的字母统计/窗口的字母统计
	needs, window := make(map[rune]int), make(map[rune]int)
	for _, v := range p {
		needs[v]++
	}
	// 已经配对的字母数
	match, ls, lp := 0, len(s), len(p)
	for right < ls {
		c1 := rune(s[right])
		if needs[c1] > 0 {
			window[c1]++
			if window[c1] == needs[c1] {
				match++
			}
		}
		//fmt.Println(needs)
		//fmt.Println(window)
		right++
		// 一旦right移动到满足配对的地方，开始调节left
		for match == len(needs) {
			// 如果 window 的大小合适
			// 就把起始索引 left 加入结果
			if right-left == lp {
				result = append(result, left)
			}
			c2 := rune(s[left])
			//fmt.Println("needs", needs)
			//fmt.Println("window", window)
			if needs[c2] > 0 {
				window[c2]--
				if window[c2] < needs[c2] {
					match--
				}
			}
			left++
		}
	}
	return result
}

func strMultiplySum(s string) int {
	sum := 1
	for i := 0; i < len(s); i++ {
		sum *= int(s[i])
	}

	return sum
}

func isSubStr(a, b string) bool {
	if len(a) == 0 || len(b) == 0 || len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if strings.IndexByte(b, a[i]) < 0 {
			return false
		}
	}

	return true
}

func strAsciiSum(s string) int {
	if len(s) == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}

	return sum
}
