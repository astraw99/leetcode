package main

import "fmt"

func main() {
	fmt.Println(fizzBuzz(5))
}

func fizzBuzz(n int) []string {
	var res []string
	if n < 1 {
		return res
	}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprint(i))
		}
	}

	return res
}
