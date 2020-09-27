package main

import "fmt"

func main() {
	matrix := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	fmt.Println(diagonalSort(matrix))
}

func diagonalSort(mat [][]int) [][]int {

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			for k := 0; i+k < len(mat) && j+k < len(mat[i]); k++ {
				if mat[i][j] > mat[i+k][j+k] {
					mat[i][j], mat[i+k][j+k] = mat[i+k][j+k], mat[i][j]
				}
			}
		}
	}

	return mat
}
