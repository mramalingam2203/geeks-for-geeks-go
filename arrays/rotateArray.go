// https://www.geeksforgeeks.org/program-for-array-rotation-continued-reversal-algorithm/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	offset := 2
	rotateArray(arr, offset)
}

func rotateArray(array []int, d int) {
	n := len(array)
	block_1 := make([]int, d)

	block_2 := make([]int, n-d)

	for i := 0; i <= d-1; i++ {
		block_1[i] = array[d-(i+1)]
	}

	for i := 0; i < n-d; i++ {
		fmt.Print(i, " ")
		block_2[i] = array[n-(i+1)]
	}

	//fmt.Println(block_1)
	fmt.Println(block_1)

}
