package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{0, 0, 0, 1, 1}}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			v := dfs(grid, visited, i, j)
			if v > max {
				max = v
			}
		}
	}

	return max
}

func dfs(grid [][]int, visited [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != 1 || visited[i][j] == 1 {
		return 0
	}

	visited[i][j] = 1

	return 1 + dfs(grid, visited, i-1, j) + dfs(grid, visited, i+1, j) + dfs(grid, visited, i, j-1) + dfs(grid, visited, i, j+1)
}
