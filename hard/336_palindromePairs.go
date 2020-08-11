package main

import "fmt"

func main() {
	words := []string{"a", "abc", "aba", ""}
	fmt.Println(palindromePairs(words))
}

func palindromePairs(words []string) [][]int {
	wordsRev := []string{}
	indices := map[string]int{}

	n := len(words)
	for _, word := range words {
		wordsRev = append(wordsRev, reverse(word))
	}
	for i := 0; i < n; i++ {
		indices[wordsRev[i]] = i
	}

	var ret [][]int
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		if m == 0 {
			continue
		}
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1, indices)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1, indices)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}

func palindromePairs2(words []string) [][]int {
	var res [][]int

	for i := 0; i < len(words); i++ {
		for idx, _ := range words {
			if i != idx {
				// corner case
				if (len(words[i]) + len(words[idx])) == 1 {
					res = append(res, []int{i, idx})
					continue
				}

				// filter odds
				if (len(words[i])+len(words[idx]))%2 != 0 {
					continue
				}

				if isPalindrome2(words[i] + words[idx]) {
					res = append(res, []int{i, idx})
				}
			}
		}
	}

	return res
}

func isPalindrome2(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
