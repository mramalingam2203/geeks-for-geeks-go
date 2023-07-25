// https://www.geeksforgeeks.org/square-root-sqrt-decomposition-algorithm
/*
Square Root Decomposition Technique is one of the most common query optimization techniques used by competitive programmers.
This technique helps us to reduce Time Complexity by a factor of sqrt(N)
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 5, 2, 4, 6, 1, 3, 5, 7, 10}

	query(input, 3, 8)
	query(input, 1, 6)
	query(input, 8, 8)

	preprocess(input)
}

func query(arr []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum

}

// Square Root Decomposition

// divide blocks

var arr [10000]int
var block [100]int

func preprocess(input []int) {
	n := len(input)
	block_size := math.Sqrt(float64(n))
	fmt.Println(int(block_size))

	blk_idx := -1

	// building the decomposed array

	for i := 0; i < n; i++ {
		arr[i] = input[i]
		if i%int(block_size) == 0 {
			// entering next block
			// incrementing block pointer
			blk_idx++
		}
		block[blk_idx] += arr[i]
	}
	fmt.Println(block)
}
