// https://www.geeksforgeeks.org/merge-two-sorted-arrays-o1-extra-space/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array_1 := []int{1, 5, 9, 10, 15, 20}
	array_2 := []int{2, 3, 8, 13}

	merge_sorted(array_1, array_2)

}

func merge_sorted(arr1, arr2 []int) {
	merged := make([]int, len(arr1)+len(arr2))
	for i := 0; i < len(arr1); i++ {
		merged[i] = arr1[i]
	}
	for i := len(arr1); i < len(merged); i++ {
		merged[i] = arr2[i-len(arr1)]
	}
	sort.Ints(merged)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = merged[i]
	}

	for i := 0; i < len(arr2); i++ {
		arr2[i] = merged[len(arr1)+i]
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
}
