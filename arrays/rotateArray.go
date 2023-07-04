// https://www.geeksforgeeks.org/program-for-array-rotation-continued-reversal-algorithm/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	offset := 2
	rotateArray(arr, offset)
}

func rotateArray(array []int, d int) {
	block_1 := make([]int, d)
	block_2 := make([]int, len(array)-d)

	for i := 0; i < d; i++ {
		block_1[d-i] = array[i]
	}

	for i := 0; i < d; i++ {
		block_1[d-i] = array[i]
	}

	fmt.Println(block_1)

}
