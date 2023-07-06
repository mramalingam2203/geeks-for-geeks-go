package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	transposematrix(matrix)
}

func transposematrix(mat [][]int) {
	//trans := [][]int{{}}
	for i := 0; i < len(mat[i]); i++ {
		fmt.Println()
		for j := 0; j < len(mat[i]); j++ {
			fmt.Print(mat[i][j], " ")
		}
	}
}
