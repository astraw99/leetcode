package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "abcabc"
	s := "aaacb"
	//s := "ababbbc"
	fmt.Println(numberOfSubstrings(s))
}

// numberOfSubstrings
// 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
// 请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
func numberOfSubstrings(s string) int {
	start, count := 0, 0
	end := start + 3
	length := len(s)
	temp := s[start:end]

	for start >= 0 && end <= length && len(temp) >= 3 {
		if strings.IndexByte(temp, 'a') > -1 && strings.IndexByte(temp, 'b') > -1 && strings.IndexByte(temp, 'c') > -1 {
			count += length - end + 1
			start++
			if end < start+3 {
				end = start + 3
			}
		} else {
			end++
		}

		if end > length {
			break
		}
		temp = s[start:end]
	}

	return count
}

// numberOfSubstrings 高手写法二
func numberOfSubstrings2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	var left, right, res int
	count := make([]int, 3)
	right = -1
	for {
		if count[0]*count[1]*count[2] == 0 {
			right++
			if right < n {
				fmt.Println(s[right], s[right]-'a', 'b')
				fmt.Printf("%T", 'b')       // int32 == rune
				fmt.Printf("%T", byte('b')) // uint8 == byte
				count[s[right]-'a']++
			} else {
				break
			}
		} else {
			res += n - right
			count[s[left]-'a']--
			left++
		}
	}

	return res
}
