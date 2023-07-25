// https://www.geeksforgeeks.org/sum-absolute-differences-pairs-given-array

package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 7, 9, 11, 14}
	sumOfPairs(array)

}

func mag(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func sumOfPairs(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if i != j {
				sum += mag(arr[i] - arr[j])
			}
		}
	}
	fmt.Println(sum)
	return sum
}
