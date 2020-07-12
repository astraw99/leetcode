package main

import (
	"fmt"
	"strings"
)

func main() {
	dict := []string{"potimzz"}
	sentence := "potimzzpotimzz"

	fmt.Println(respace(dict, sentence))
}

// respace 恢复空格 todo 还没提交成功
// 哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子
// "I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。
// 在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。
// 假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。
func respace(dictionary []string, sentence string) int {
	var res []string
	sentTemp := []string{sentence}

	for _, v := range dictionary {
		for i, subSent := range sentTemp {
			if idx := strings.Index(subSent, v); idx != -1 {
				res = append(res, v)
				sentTemp = append(sentTemp[:i], sentTemp[i+1:]...)

				if len(subSent[:idx]) != 0 {
					sentTemp = append(sentTemp, subSent[:idx])
				}
				if len(subSent[idx+len(v):]) != 0 {
					sentTemp = append(sentTemp, subSent[idx+len(v):])
				}
			}
		}

	}

	fmt.Println(res, sentTemp)

	return len(res) + len(sentTemp)
}
