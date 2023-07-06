package main

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	matrix = transposematrix(matrix)

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

func reverseColumns(mat [][]int) [][]int {
	cols := len(mat[1])
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			mat[j][i] = mat[j][cols-i]
		}
	}

	fmt.Println (mat)
}



}
