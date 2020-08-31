package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "Let's take LeetCode contest"
	s := "you like me"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	var str string
	var sep = " "

	if len(s) == 0 {
		return str
	}

	splitArr := strings.Split(s, sep)

	var reversedArr []string
	for _, v := range splitArr {
		reversedArr = append(reversedArr, reverseStr(v))
	}

	return strings.Join(reversedArr, sep)
}

func reverseStr(s string) string {
	if len(s) <= 1 {
		return s
	}

	byteArr := []byte(s)

	i, j := 0, len(s)-1
	for i < len(s)/2 {
		byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
		i++
		j--
	}

	return string(byteArr)
}
