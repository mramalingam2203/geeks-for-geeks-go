// https://www.geeksforgeeks.org/find-number-of-triangles-possible/

package main

import (
	"fmt"
	"sort"
)

func main() {
	sides := []int{10, 21, 22, 100, 101, 200, 300}
	fmt.Println(byBruteForce(sides))
	fmt.Printn(bySorting(sides))

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

func bySorting(arr []int) int {
	count := 0
	sort.Ints(arr)
	var i, j, k int
	n := len(arr)
	for i = 0; i < n-2; i++ {
		k = i + 2
		for j = i + 1; j < n; j++ {
			for {
				k++
				if !(k < n && arr[i]+arr[j] > arr[k]) == true {
					break
				}
			}
			if k > j {
				count += k - j - 1
			}
		}

	}
	return count

}
