package main

import (
	"fmt"
)

func main() {
	obj := Constructor([]string{"pear", "apple"})
	fmt.Println(obj.F("a", "e"))
}

type WordFilter struct {
	words []string
}

func Constructor(words []string) WordFilter {
	return WordFilter{
		words: words,
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	var res = -1
	length := len(this.words)
	for i := length - 1; i >= 0; i-- {
		v := this.words[i]
		fmt.Println(v[:len(prefix)], v[len(v)-len(suffix):])
		if v[:len(prefix)] == prefix && v[len(v)-len(suffix):] == suffix {
			return i
		}
		/*if strings.Index(v, prefix) == 0 && strings.LastIndex(v, suffix)+len(suffix) == len(v) {
			return i
		}*/
	}

	/*for i, v := range this.words {
		if strings.Index(v, prefix) == 0 && strings.LastIndex(v, suffix)+len(suffix) == len(v) {
			res = i
		}
	}*/

	return res
}
