// https://www.geeksforgeeks.org/smallest-difference-triplet-from-three-arrays/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

}

func simpleMethod(arr1 []int, arr2 []int, arr3 []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	sort.Ints(arr3)
	n := len(arr1)
	var sum, max, min int
	var i, j, k, res_min, res_max, res_mid int
	const diff int = math.MaxInt64
	for i < n && j < n && k < n {
		sum = arr1[i] + arr2[j] + arr3[k]
		// maximum number
		max = maximum(arr1[i], arr2[j], arr3[k])
		// Find minimum and increment its index.
		min = minimum(arr1[i], arr2[j], arr3[k])
		if min == arr1[i] {
			i++
		} else if min == arr2[j] {
			j++
		} else {
			k++
		}

		if diff > (max - min) {
			diff = max - min
			res_max = max
			res_mid = sum - (max + min)
			res_min = min
		}
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maximum(a int, b int, c int) int {
	return max(max(a, b), c)
}

func minimum(a int, b int, c int) {
	return min(min(a, b), c)
}
