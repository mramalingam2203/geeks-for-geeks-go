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
	for i := 0; i < len(arr); i++ {
		if arr[i+1]%2 == 0&arr1[i+2]%2 != 0 {
			tmp = arr[i+2]
			arr[i+1] = arr[i+2]
			arr[i+2] = tmp
		}
	}
	fmt.Println(arr)

}
