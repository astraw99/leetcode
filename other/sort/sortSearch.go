package main

import (
	"fmt"
	"sort"
)

func main() {
	GuessingGame()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")

	fmt.Println(111)

	answer := sort.Search(100, func(i int) bool {

		fmt.Println(222)
		fmt.Printf("s : %v\n", s)

		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})

	fmt.Println(333)

	fmt.Printf("Your number is %d.\n", answer)
}
