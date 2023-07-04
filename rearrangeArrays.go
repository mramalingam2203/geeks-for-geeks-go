// https://www.geeksforgeeks.org/rearrange-array-such-that-even-positioned-are-greater-than-odd/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.Ints(a)
	rearrangeArray(a)

}

func rearrangeArray(arr []int) {
	var tmp int
	for i := 1; i < len(arr)-1; i += 2 {
		if arr[i] < arr[i+1] {
			tmp = arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}
	fmt.Println(arr)

}
