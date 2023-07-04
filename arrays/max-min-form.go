// https://www.geeksforgeeks.org/rearrange-array-maximum-minimum-form/

package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	rearrange(a)

}

func rearrange(arr []int) {
	n := len(arr)

	// Auxiliary array to hold modified array
	temp := make([]int, n)
	// Indexes of smallest and largest elements
	// from remaining array.
	small := 0
	large := n - 1

	// To indicate whether we need to copy remaining
	// largest or remaining smallest at next position

	var flag bool = true

	// Store result in temp[]

	for i := 0; i < n; i++ {
		if flag == true {
			large--
			temp[i] = arr[large]
			fmt.Print(temp[i])
		} else {
			small++
			temp[i] = arr[small]
			fmt.Print(temp[i])
		}

		flag = !flag
	}

	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}

	fmt.Println(arr)
}
