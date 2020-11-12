package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	var res int
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(grid, i, j) >= 1 {
				res++
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' || grid[i][j] == '*' {
		return 0
	}

	grid[i][j] = '*'

	return 1 + dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
}
