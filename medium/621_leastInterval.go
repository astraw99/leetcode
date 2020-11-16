package main

import "fmt"

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}

// 1. 计算每个任务出现的次数
// 2. 找出出现次数最多的任务，假设出现次数为 x
// 3. 计算至少需要的时间 (x - 1) * (n + 1)，记为 min_time
// 4. 计算出现次数为 x 的任务总数 count，计算最终结果为 min_time + count
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	maxCnt, maxTimes := 0, 0

	var m = make(map[byte]int, 26)
	for _, v := range tasks {
		m[v]++
		if maxCnt < m[v] {
			maxCnt = m[v]
			maxTimes = 1
		} else if maxCnt == m[v] {
			maxTimes++
		}
	}

	minTime := (maxCnt - 1) * (n + 1)
	totalTime := minTime + maxTimes
	if totalTime < len(tasks) {
		return len(tasks)
	}

	return totalTime
}
