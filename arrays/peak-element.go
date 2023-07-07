// find peak element

// https://www.geeksforgeeks.org/find-a-peak-in-a-given-array/

package main

import "fmt"

func main() {
	a := []int{10, 20, 15, 2, 23, 90, 67}
	fmt.Println(peakElement(a))
	fmt.Println(findPeak(a))
}

func peakElement(a []int) []int {
	result := make([]int, 0)
	for i := 1; i < len(a)-1; i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			result = append(result, a[i])
		}
	}
	return result

}

// A wrapper over recursive function findPeakUtil()
func findPeak(arr []int) int {
	return findPeakUtil(arr, 0, len(arr)-1, len(arr))
}

func findPeakUtil(arr []int, low, high, n int) int {

	mid := low + (high-low)/2

	if mid == 0 || arr[mid-1] <= arr[mid] && mid == n-1 || arr[mid+1] <= arr[mid] {
		return mid
	} else if mid > 0 && arr[mid-1] > arr[mid] {
		return findPeakUtil(arr, low, (mid - 1), n)
	} else {
		return findPeakUtil(arr, (mid + 1), high, n)
	}
}


