// https://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 7
	fmt.Println(binarySearch(a, key))
}

func binarySearch(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1 // target not found
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid // target found
	} else if arr[mid] > target {
		return binarySearchRecursive(arr, target, left, mid-1) // search left half
	} else {
		return binarySearchRecursive(arr, target, mid+1, right) // search right half
	}
}

func binarySearch_iter(arr []int, target int) int {
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
