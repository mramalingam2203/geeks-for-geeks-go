// https://www.geeksforgeeks.org/rearrange-array-arri/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	rearrangeArray_1(a)

}

func rearrangeArray_1(arr []int) {
	sort.Ints(arr)
	for i := range arr {
		if arr[i] != i {
			arr[i] = -1
		}
	}
	fmt.Println(arr)

}
