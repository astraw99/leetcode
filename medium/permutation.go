package main

import "fmt"

func main() {
	s := "qqe"
	fmt.Println(permutation(s))
}

func permutation(S string) []string {
	var r []string
	if len(S) == 0 {
		r = append(r, "")
		return r
	}
	for i := 0; i < len(S); i++ {
		sl := []byte(S)
		if !inArray(sl[:i], S[i]) {
			var r0 []string
			if i == 0 {
				r0 = permutation(S[1:])
			} else if i == len(S)-1 {
				r0 = permutation(S[:len(S)-1])
			} else {
				r0 = permutation(S[:i] + S[i+1:])
			}
			fmt.Println(r0, string([]byte{S[i]}))
			for j, _ := range r0 {
				s := string([]byte{S[i]}) + r0[j]
				r = append(r, s)
			}
		}
	}
	return r
}

func inArray(s []byte, ta byte) bool {
	for i, _ := range s {
		if s[i] == ta {
			return true
		}
	}
	return false
}
