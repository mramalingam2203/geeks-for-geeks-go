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

	block_2 := make([]int, len(array)-d)

	for i := 0; i <= d-1; i++ {
		block_1[i] = array[d-(i+1)]
		//fmt.Print(i, " ")
	}

	for i := d; i < n; i++ {
		fmt.Print(i, " ")
		block_2[i] = array[n-(i+1)]
		//block_2[len(array)-d-i] = array[len(array)-d-(i+1)]
	}

	//fmt.Println(block_1)
	//fmt.Println(block_2)

}
