// https://www.geeksforgeeks.org/majority-element

package main

import "fmt"

func main() {
	array := []int{3, 4, 3, 2, 4, 4, 4, 4}
	findTheMajorityElement_naive(array)
	findTheMajorityElement_mooresvoting(array)
	findTheMajorityElement_hashing(array)
}

// Naive approach
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

// Binary Search Tree
func findTheMajorityElement_BST(arr []int) {}

// Moore's Voting Algoprithm
func findTheMajorityElement_mooresvoting(arr []int) {
	n := len(arr)
	maj_index := 0
	count := 1
	for i := 1; i < n; i++ {
		if arr[i] == arr[maj_index] {
			count++
		} else {
			count--
		}
		if count == 0 {
			maj_index = i
			count = 1
		}
	}

	count = 0
	isMajor := false
	for i := 0; i < n; i++ {
		if arr[i] == arr[maj_index] {
			count++
			if count > n/2 {
				isMajor = true
				break
			}
		}
	}

	if isMajor == true {
		fmt.Println("Majority element exist: ", arr[maj_index])
	} else {
		fmt.Println("No majority element")
	}
}

func findTheMajorityElement_hashing(arr []int) {
	hash := make(map[int]int)
	n := len(arr)
	count := 0

	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < n; j++ {
			if arr[j] == arr[i] {
				count++
				hash[arr[j]] = count
			} else {
				hash[arr[j]] = count
			}
		}
	}
	fmt.Println(hash)
}

func findTheMajorityElement_sorting(arr []int) {}

func findTheMajorityElement_recursion(arr []int) {

	
}
