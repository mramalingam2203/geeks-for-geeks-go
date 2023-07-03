// https://www.geeksforgeeks.org/smallest-difference-triplet-from-three-arrays/

package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {

}

func simpleMethod(arr1 []int, arr2 []int, arr3 []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	sort.Ints(arr3)
	var sum, max, min int
	var i, j, k, res_min, res_max, res_mid
	const diff int = math.MaxInt64)
	for ; i < n && j < n && k < n;{

		sum = arr1[i] + arr2[j] + arr3[k];

        // maximum number
        max = maximum(arr1[i], arr2[j], arr3[k]);
 
        // Find minimum and increment its index.
        min = minimum(arr1[i], arr2[j], arr3[k]);

	}

}

func max (a int, b int){
	if a < b {
		return b
	}
	return a
}


func min (a int, b int){
	if a > b {
		return b
	}
	return a
}

func maximum(a int, b int, c int){
	return max(max(a, b), c)
}


func minimum(a int, b int, c int){
	return min(min(a, b), c)
}
