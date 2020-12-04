package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}

	var letterMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack = func(digits string, path string, start int) {}
	backtrack = func(digits string, path string, start int) {
		if start == len(digits) {
			res = append(res, path)
			return
		}

		letters := letterMap[digits[start]]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, path+string(letters[i]), start+1)
		}
	}

	backtrack(digits, "", 0)

	return res
}
