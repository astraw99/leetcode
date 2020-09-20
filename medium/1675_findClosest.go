package main

import (
	"fmt"
	"math"
)

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"

	fmt.Println(findClosest(words, word1, word2))
}

// findClosest [LeetCode #1675]
// 有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)
// 如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
func findClosest(words []string, word1 string, word2 string) int {
	var temp1 []int
	var temp2 []int

	for i, v := range words {
		if v == word1 {
			temp1 = append(temp1, i)
		}

		if v == word2 {
			temp2 = append(temp2, i)
		}
	}

	//fmt.Println(temp1, temp2)

	distance := math.MaxInt32
	for _, v := range temp1 {
		for _, val := range temp2 {
			min := int(math.Abs(float64(val - v)))
			//fmt.Println(min)
			if min < distance {
				distance = min
			}
			if distance == 1 {
				return distance
			}
		}
	}

	return distance
}
