// https://www.geeksforgeeks.org/recamans-sequence/

package main

import (
	"fmt"
)

func main() {
	reccaman(17)

}

func reccaman(n int) {
	arr := make([]int, n)
	// First term of the sequence is always 0
	arr[0] = 0

	for i := 1; i < len(arr); i++ {
		curr := arr[i-1] - i
		for j := 0; j < i; j++ {
			if arr[j] == curr || curr < 0 {
				curr = arr[i-1] + i
				break
			}
		}
		arr[i] = curr
		fmt.Print(arr[i], " ")
	}

}
