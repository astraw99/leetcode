package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}

	var backtrack func(s string, path []string, pos int)
	backtrack = func(s string, path []string, pos int) {
		if pos == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := pos; i < len(s); i++ {
			if isPalindrome(s[pos : i+1]) {
				path = append(path, s[pos:i+1])
				backtrack(s, path, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(s, []string{}, 0)

	return res
}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-1] {
			return false
		}
		l--
	}

	return true
}
