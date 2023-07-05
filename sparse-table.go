// https://www.geeksforgeeks.org/sparse-table/

package main

import "fmt"

func main() {
	arr := []int{7, 2, 3, 0, 5, 10, 3, 12, 18}
	query := [][]int{{0, 4}, {4, 7}, {7, 8}}

	range_minimum_query(arr, query)

}

func range_minimum_query(array []int, queries [][]int) {
	for i := range queries {
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			
			temp := array[j]
			
			if array[j+1] > max {
			

		}
		fmt.Println()
	}

}
