package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ac", "bb"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//sum1, sum2 := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
		//sum1 += int(s[i] - 'a')
		//sum2 += int(t[i] - 'a')
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
