package main

import (
	"fmt"
)

func main() {
	J := "aAc"
	S := "aAABbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	var res int
	var mapS = make(map[rune]int)

	for i := 0; i < len(S); i++ {
		// S[i] 默认为 ascii 编码 byte(uint8)
		fmt.Printf("%T %v\n", S[i], S[i])
		// 默认字符(char) 为 unicode 编码 rune(int32)
		fmt.Printf("%T\n", 'a')
	}

	for _, v := range S {
		// range 会隐式转为 unicode 编码字符 rune(int32)
		if count, ok := mapS[v]; ok {
			mapS[v] = count + 1
		} else {
			mapS[v] = 1
		}
	}

	for _, v := range J {
		res += mapS[v]
	}

	return res
}
