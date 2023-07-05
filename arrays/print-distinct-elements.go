// Print All Distinct Elements of a given integer array

// https://www.geeksforgeeks.org/print-distinct-elements-given-integer-array/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2, 3, 4, 4, 3, 5, 2, 2, 1, 2, 3, 4, 5}
	//distinct_elements(array)
	distinct_elements_by_sort(array)

}

func distinct_elements(arr []int) {
	var i, j int
	for i = 0; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[i] == arr[j] {
				break
			}
		}
		if i == j {
			fmt.Print(arr[i], " ")
		}

	}

}

func distinct_elements_by_sort(arr []int) {
	sort.Ints(arr)
	var i int
	for i = range arr {
		for {
			if i < len(arr)-1 && arr[i] == arr[i+1] {
				i++
				break
			}
			fmt.Print(arr[i], " ")
			break
		}

	}
}
