// https://www.geeksforgeeks.org/square-root-sqrt-decomposition-algorithm
/*
Square Root Decomposition Technique is one of the most common query optimization techniques used by competitive programmers.
This technique helps us to reduce Time Complexity by a factor of sqrt(N)
*/

package main

func main() {
	input := []int{1, 5, 2, 4, 6, 1, 3, 5, 7, 10}
	l := 3
	r := 8
	query(l.r)
}

func query(arr []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum

}
