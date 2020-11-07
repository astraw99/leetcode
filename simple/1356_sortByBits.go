package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []int{1111, 7644, 1107, 6978, 8742, 1, 7403, 7694, 9193, 4401, 377, 8641, 5311, 624, 3554, 6631}
	fmt.Println(sortByBits(arr))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		//binI, _ := strconv.Atoi(fmt.Sprintf("%b", arr[i]))
		//binJ, _ := strconv.Atoi(fmt.Sprintf("%b", arr[j]))
		//fmt.Println(binI, binJ)
		cntI := strings.Count(fmt.Sprintf("%b", arr[i]), "1")
		cntJ := strings.Count(fmt.Sprintf("%b", arr[j]), "1")
		return cntI < cntJ || cntI == cntJ && arr[i] < arr[j]
	})

	return arr
}
