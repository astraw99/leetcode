package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "bbbacddceeb"
	B := "ceebbbbacdd"

	fmt.Println(rotateString(A, B))
	fmt.Println(rotateString2(A, B))
}

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}

// rotateString 旋转字符串
// 给定两个字符串, A 和 B。
// A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。
func rotateString2(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(A) == 0 && len(B) == 0 {
		return true
	}

	var posArr []int
	needle := B[0]
	for i := 0; i < len(A); i++ {
		if A[i] == needle {
			posArr = append(posArr, i)
		}
	}

	for _, pos := range posArr {
		if B == A[pos:]+A[:pos] {
			return true
		}
	}

	return false
}
