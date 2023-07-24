// https://www.geeksforgeeks.org/modify-array-to-another-given-array-by-replacing-with-the-sum-of-the-array/

package main

import "fmt"

func main() {
	array := []int{1, 1, 1, 2}

	fmt.Println(allOnes(array))
}

func largest(array []int) int {
	temp := 0
	var idx int
	for i := range array {
		if array[i] > temp {
			temp = array[i]
			idx = i
		}
	}
	return idx
}

func sumArray(array []int, largest int) int {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	sum -= largest
	return sum
}

func allOnes(array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] != 1 {
			return false
		}
		continue
	}
	return true
}

func modifyArrayToAnother(target []int) bool {
	var idx int
	for allOnes(target) == false {
		idx = largest(target)
		target[idx] = target[idx] - sumArray(target, target[idx])
		fmt.Println(idx)
	}

	return false
}
