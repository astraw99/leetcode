package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	B := []int{2, 5, 6}
	merge(A, 3, B, 3)
	fmt.Println(A)
}

func merge(A []int, m int, B []int, n int) {
	copy(A[m:], B)
	sort.Ints(A)
}

// self wrong!!
func mergeSelf(A []int, m int, B []int, n int) {
	var res []int

	for len(A) > 0 || len(B) > 0 {
		if len(A) > 0 && len(B) > 0 {
			if A[0] <= B[0] {
				res = append(res, A[0])
				A = A[1:]
			} else {
				res = append(res, B[0])
				B = B[1:]
			}
		} else if len(A) > 0 {
			res = append(res, A...)
			A = nil
		} else if len(B) > 0 {
			res = append(res, B...)
			B = nil
		}
	}

	fmt.Println(res)

	A = res
}
