package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aab"
	fmt.Println(permutation(s))
}

func permutation(s string) []string {
	var res []string
	if len(s) == 0 {
		return []string{""}
	}

	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[:i], s[i]) < 0 {
			var r0 []string
			if i == 0 {
				r0 = permutation(s[1:])
			} else if i == len(s)-1 {
				r0 = permutation(s[:len(s)-1])
			} else {
				r0 = permutation(s[:i] + s[i+1:])
			}

			for j, _ := range r0 {
				res = append(res, string(s[i])+r0[j])
			}
		}
	}

	return res
}
