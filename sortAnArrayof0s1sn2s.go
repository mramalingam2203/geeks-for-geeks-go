// https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/

package main

import (
	"fmt"
	"os"
)

func main() {
	array := []int{0, 1, 1, 2, 1, 1, 2, 0, 1, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0}
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

}
