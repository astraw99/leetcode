package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Println(permutation(s))
}

func permutation(S string) []string {
	var r []string
	if len(S) == 0 {
		r = append(r, "")
		return r
	}

	for i := 0; i < len(S); i++ {
		if !inArray(S[i], []byte(S[:i])) {
			var r0 []string
			if i == 0 {
				r0 = permutation(S[1:])
			} else if i == len(S)-1 {
				r0 = permutation(S[:len(S)-1])
			} else {
				r0 = permutation(S[:i] + S[i+1:])
			}

			for j, _ := range r0 {
				str := string(S[i]) + r0[j]
				r = append(r, str)
			}
		}
	}

	return r
}

func inArray(b byte, bytes []byte) bool {
	for i, _ := range bytes {
		if bytes[i] == b {
			return true
		}
	}

	return false
}

// self code, not right !!!
func Permutation(S string) []string {
	var res []string

	for i, _ := range S {
		var sub []byte
		sub = append(sub, S[0:i]...)
		sub = append(sub, S[i+1:]...)
		res = append(res, subPermutation(string(sub), S[i])...)
	}

	var ret []string
	var distinct = make(map[string]int)
	for _, str := range res {
		if _, ok := distinct[str]; ok {
			continue
		}
		ret = append(ret, str)
		distinct[str] += 1
	}

	return ret
}

func subPermutation(sub string, b byte) []string {
	var res []string

	res = append(res, sub+string(b))

	for i, _ := range sub {
		var bytes []byte
		bytes = append(bytes, sub[0:i]...)
		bytes = append(bytes, b)
		bytes = append(bytes, sub[i:]...)
		str := string(bytes)
		res = append(res, str)
	}

	return res
}
