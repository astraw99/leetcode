package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := uint32(1024)
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	return strings.Count(strconv.FormatUint(uint64(num), 2), "1")
}
