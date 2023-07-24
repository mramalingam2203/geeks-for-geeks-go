// Print All Distinct Elements of a given integer array

// https://www.geeksforgeeks.org/print-distinct-elements-given-integer-array

package main

import (
	"fmt"
	"sort"
	_ "sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 6, 7, 5, 5, 6, 8, 9, 10, 10, 10, 11, -23, 23, 34}
	printDistinctElements(arr)
}

func printDistinctElements(array []int) {
	sort.Ints(array)
	fmt.Println(array[0])
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			fmt.Println(array[i])
		}
	}

}
