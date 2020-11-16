package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	s := "3[a2[c]]" // accaccacc
	fmt.Println(decodeString(s))
	fmt.Println(decodeString_(s))
}

func decodeString(s string) string {
	type pair struct {
		mul  int    //当前 [] 中字符串的倍数
		last string //是上个 [ 到当前 [ 的字符串
	}
	stack := list.New()
	var multi int = 0
	var res string = ""
	for _, v := range s {
		if v >= '0' && v <= '9' {
			multi = multi*10 + int(v-'0') //统计下一个 [ 前的倍数
		} else if v == '[' {
			stack.PushBack(pair{multi, res}) //把当前 [ 前所有的字符串组成的序列和当前[倍数压入栈
			multi, res = 0, ""
		} else if v == ']' {
			temp := stack.Back().Value.(pair)
			stack.Remove(stack.Back())
			res = temp.last + strings.Repeat(res, temp.mul) // 字符串拼接
		} else {
			res += string(v) //遇到字母就直接拼接到当前res中直到遇到 ] 时就进行 更新字符串序列
		}
	}

	return res
}

// self code - right :)
func decodeString_(s string) string {
	var res string
	if len(s) == 0 {
		return ""
	}

	type Pair struct {
		Times   int
		LastStr string
	}

	var stack []Pair
	var times int
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			times = times*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stack = append(stack, Pair{Times: times, LastStr: res})
			res = ""
			times = 0
		} else if s[i] == ']' {
			if len(stack) > 0 {
				lastStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res = lastStack.LastStr + strings.Repeat(res, lastStack.Times)
			}
		} else {
			res += string(s[i])
		}
	}

	return res
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
