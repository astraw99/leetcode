package main

import "fmt"

func main() {
	n := 120
	fmt.Println(cuttingRope2(n))
}

func cuttingRope2(n int) int {
	var res = 1

	if n <= 3 {
		return n - 1
	}

	if n%3 == 1 {
		res *= 4
		n -= 4
	}

	if n%3 == 2 {
		res *= 2
		n -= 2
	}

	for i := 0; i < n/3; i++ {
		res %= 1e9 + 7
		res *= 3
	}

	return res % (1e9 + 7)
}
