package main

import (
	"fmt"
	"sort"
)

func main() {

	array1 := [][]int{{1, 2, 3, 20},
		{5, 6, 20, 25},
		{1, 3, 5, 6},
		{6, 7, 8, 15}}

	rotateMatrix(array1)
	findUniqueElements(array1)

	array2 := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}

	rotateMatrixByNinetyDegrees(array2)
	sortMatrix(array2)

}

func rotateMatrix(mat [][]int) {
	mat = transposematrix(mat)
	mat = reverseColumns(mat)
	mat = transposematrix(mat)
	mat = reverseColumns(mat)
	fmt.Println(mat)

}
func transposematrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[1])

	// Create a new matrix with dimensions swapped
	transposed := make([][]int, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]int, rows)
	}

	// Fill the transposed matrix with values from the original matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = mat[i][j]
		}
	}

	return transposed

}

func reverseColumns(arr [][]int) [][]int {
	rows := len(arr)
	if rows == 0 {
		return arr
	}

	cols := len(arr[0])
	if cols == 0 {
		return arr
	}

	reversed := make([][]int, rows)
	for i := 0; i < rows; i++ {
		reversed[i] = make([]int, cols)
	}

	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			reversed[i][j] = arr[i][cols-j-1]
		}
	}
	return reversed
}

func findUniqueElements(mat [][]int) []int {

	rows := len(mat)
	cols := len(mat[1])

	// find maximum number in matrix
	maximum := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maximum < mat[i][j] {
				maximum = mat[i][j]
			}
		}
	}

	// Take 1-D array of (maximum + 1) size
	b := make([]int, maximum+1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			y := mat[i][j]
			b[y]++
		}
	}

	for i := 1; i <= maximum; i++ {
		if b[i] == 1 {
			fmt.Print(i, " ")
		}
	}

	return b

}

func rotateMatrixByNinetyDegrees(matrix [][]int) {
	matrix = reverseColumns(matrix)
	matrix = transposematrix(matrix)

	fmt.Println()
	fmt.Print(matrix)

}

func sortMatrix(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	matrix_1D := make([]int, rows*cols)
	for i := 0; i < rows*cols; i++ {
		matrix_1D[i] = matrix[i/rows]
	}

	sort.Ints(matrix_1D)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = matrix_1D[i/rows]
		}
	}
	fmt.Println(matrix)
}
