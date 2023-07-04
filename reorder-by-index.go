// https://www.geeksforgeeks.org/reorder-a-array-according-to-given-indexes/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	seq := []int{0, 3, 2, 1}
	output := reorder(a, seq)
	fmt.Println(output)

}

func reorder(a []int, index []int) []int {
	result := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		result = append(result, a[index[i]])
	}
	return result
}
