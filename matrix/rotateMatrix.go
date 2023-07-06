package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	transposematrix(matrix)
}

func transposematrix(mat [][]int) {
	result := make([]int, 0, len(mat)*len(mat))

	for i := 0; i < len(mat); i++ {
		fmt.Println()
		for j := 0; j < len(mat); j++ {
			result = append(result, mat[i][j])
		}
	}

	fmt.Println(result)
}
