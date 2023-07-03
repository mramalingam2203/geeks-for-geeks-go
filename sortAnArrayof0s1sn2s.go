// https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/

package main

import (
	"fmt"
	//"os"
)

func main() {
	array := []int{0, 1, 1, 2, 1, 1, 2, 0, 1, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0}
	sorted_array := make([]int, len(array))
	sorted_array = sortAnArray(array)
	fmt.Println(sorted_array)
}

func sortAnArray(a []int) []int {
	var count_0, count_1, count_2 = 0, 0, 0
	for i := range a {
		if a[i] == 0 {
			count_0++
		} else if a[i] == 1 {
			count_1++
		} else {
			count_2++
		}
	}
	fmt.Println(count_0, count_1, count_2)

	result := make([]int, count_0+count_1+count_2)

	for i := range result {
		if i < count_0 {
			result[i] = 0
		} else if i >= count_0 && i < count_1 {
			result[i] = 1
		} else if i >= count_1 && i < count_2 {
			result[i] = 2
		}
	}

	return result
}
