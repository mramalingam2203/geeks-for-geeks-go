// https://www.geeksforgeeks.org/reorder-a-array-according-to-given-indexes/

package main

func main() {
	a := []int {1, 2, 3, 4}
	seq := []int{}{0,3,2, 1}
	output := reorder(a, seq)

}

func reorder(a []int, index []int) []int {
	result := make([]int, 0, len(a))
	return result

}

