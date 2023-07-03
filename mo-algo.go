// https://www.geeksforgeeks.org/mos-algorithm-query-square-root-decomposition-set-1-introduction/

package main

import (
	"fmt"
)

func printQuerySums(a []int, q [][]int) {
	//len_a := len(a)
	len_q := len(q)

	for i := 0; i < len_q; i++ {
		sum := 0
		for j := a[q[i][0]]; j <= q[i][1]; j++ {
			fmt.Print(a[j], " ")
			sum += a[j]
		}
		fmt.Println(sum)
	}

}

func main() {
	array := []int{1, 1, 2, 1, 3, 4, 5, 2, 8}
	queries := [][]int{{0, 4}, {1, 3}, {2, 4}}

	printQuerySums(array, queries)
}
