// https://www.geeksforgeeks.org/majority-element

package main

import "fmt"

func main() {
	array := []int{3, 4, 3, 2, 4, 4, 4, 4}
}

func findTheMajorityElement_naive(arr []int) {
	n := len(arr)
	count, maxCount := 0, 0
	var index int
	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			index = i
		}
	}
	if maxCount > n/2 {
		fmt.Println(arr[index])
	} else {
		fmt.Println("no majority element found")
	}

}

func findTheMajorityElement_BST(arr []int)
func findTheMajorityElement_mooresvoting(arr []int)
func findTheMajorityElement_hashing(arr []int)
func findTheMajorityElement_sorting(arr []int)
func findTheMajorityElement_recursion(arr []int)
