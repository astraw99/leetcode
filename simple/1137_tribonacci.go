package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}

	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}

	return a
}
