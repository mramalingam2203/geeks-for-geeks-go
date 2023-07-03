// https://www.geeksforgeeks.org/find-number-of-triangles-possible/

package main

import (
	"fmt"
	"sort"
)

func main() {
	sides := []int{10, 21, 22, 100, 101, 200, 300}
	fmt.Println(byBruteForce(sides))
	bySorting

}

func byBruteForce(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				// third
				if (arr[i]+arr[j] > arr[k] && arr[i]+arr[k] > arr[j] && arr[k]+arr[j] > arr[i]) == true {
					count++
				}
			}
		}
	}
	return count
}

func bySorting(a []int) {
	fmt.Println(sort.Ints(a))

}
