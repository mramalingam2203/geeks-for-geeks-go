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
	k := 6
	minOps(array, k)
	fmt.Println(max(array))
}

func max(array []int) int {
	n := len(array)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func minOps(arr []int, K int) int {
	return 0

}
