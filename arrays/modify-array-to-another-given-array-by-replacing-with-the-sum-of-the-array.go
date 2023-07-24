// https://www.geeksforgeeks.org/modify-array-to-another-given-array-by-replacing-with-the-sum-of-the-array/

package main

import "fmt"

func main() {
	array := []int{9, 3, 5}
	modifyArrayToAnother(array)
}

func largest(array []int) int {
	temp := array[0]
	for i := range array {
		if array[0] > temp {
			temp = array[i]
		}
	}
	return temp
}

func sumArray(array []int, largest int) int {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	sum -= largest
	return sum
}

func modifyArrayToAnother(target []int) bool {
	for i := range target {
		fmt.Println(largest(target))
	}
	return false
}
