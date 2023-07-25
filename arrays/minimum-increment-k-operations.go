// https://www.geeksforgeeks.org/minimum-increment-k-operations-make-elements-equal/
/*
You are given an array of n-elements, you have to find the number of operations needed to make all elements of array equal. Where a single operation can increment an element by k. If it is not possible to make all elements equal print -1.

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{21, 33, 9, 45, 63}
	minOps(array, k)
}

func max(array []int) {
	n := len(array)
	temp := math.MaxInt64
	for i := 0; i < n; i++ {
		if array[i] > temp {
			temp = array[i]
		}
	}
	return temp
}

func minOps(arr []int, K int) int {

}
