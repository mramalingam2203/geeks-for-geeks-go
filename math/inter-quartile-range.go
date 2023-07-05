// https://www.geeksforgeeks.org/interquartile-range-iqr/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 3, 4, 5, 5, 6, 7, 11}
	iqr(array)
}

func iqr(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr)
	//Q2 := arr[len(arr)/2]
	Q1 := arr[len(arr)/4]
	Q3 := arr[3*len(arr)/4]

	fmt.Println(Q3 - Q1)

}
