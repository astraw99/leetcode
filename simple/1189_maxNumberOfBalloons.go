package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) int {
	var sum int
	target := "balloon"

	if len(text) < len(target) {
		return sum
	}

	for _, v := range target {
		if pos := strings.IndexRune(text, v); pos == -1 {
			return sum
		} else {
			text = text[:pos] + text[pos+1:]
		}
	}

	return 1 + maxNumberOfBalloons(text)
}
