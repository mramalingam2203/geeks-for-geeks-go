// https://www.geeksforgeeks.org/find-the-first-missing-number/

package main

import "fmt"

func main() {

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 10}
	m := 11

	fmt.Println(findFirstMissingNumber(array, m))

}

func findFirstMissingNumber(arr []int, m int) int {
	for i := 0; i < m-1; i++ {
		if binarySearch(arr, i) == -1 {
			return i
		}
	}

	return 0

}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Target not found
}
