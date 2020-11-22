package main

import "fmt"

func main() {
	numCourses := 10
	prerequisites := [][]int{
		{2, 1},
		{5, 4},
	}

	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses) // [first][]{indegs}
		indegs = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indegs[info[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indegs[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		current := q[0]
		result = append(result, current)
		q = q[1:]

		for _, v := range edges[current] {
			indegs[v]--
			if indegs[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return len(result) == numCourses
}

func _canFinish(numCourses int, prerequisites [][]int) bool {
	learned := make(map[int]int)
	sum := 0

	for _, row := range prerequisites {
		length := len(row)
		for i := length - 1; i > 0; i-- {
			if _, ok := learned[row[i]]; ok {
				if sum <= numCourses {
					return true
				}
				return false
			}

			learned[row[i]] = row[i]
		}

		sum += length
	}

	if sum <= numCourses {
		return true
	}

	return false
}
