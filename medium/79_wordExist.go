package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

type pair struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	var check func(i, j, k int) bool // i,j 位置，k 个字符 word[k]
	check = func(i, j, k int) bool {

		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()

		directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			newI, newJ := i+dir.x, j+dir.y
			if 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}

		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}

	return false
}
