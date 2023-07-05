// https://www.geeksforgeeks.org/sort-array-wave-form-2/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{10, 90, 49, 2, 1, 5, 23}
	wave_sort(array)
}

/*
- Sort the array.
- Traverse the array from index 0 to N-1, and increase the value of the index by 2.
- While traversing the array swap arr[i] with arr[i+1].
- Print the final array.
*/

func wave_sort(a []int) {
	var tmp int
	sort.Ints(a)
	for i := 0; i < len(a)-1; i += 2 {
		tmp = a[i]
		a[i] = a[i+1]
		a[i+1] = tmp
	}
	fmt.Print(a)

}
