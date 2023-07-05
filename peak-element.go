// find peak element

// https://www.geeksforgeeks.org/find-a-peak-in-a-given-array/

package main

import "fmt"

func main() {
	a := []int{5, 10, 20, 15}
	fmt.Println(peakElement(a))
}

func peakElement(a []int) int {
	for i := 1; i < len(a)-1; i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			return a[i]
		}
	}
	return -1

}
