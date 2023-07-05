// https://www.geeksforgeeks.org/print-first-k-digits-1n-n-positive-integer/

package main

import (
	"fmt"
)

func main() {
	printFirstkDigits(21, 4)
}

func printFirstkDigits(n, k int) {
	rem := 1
	for i := 0; i < k; i++ {
		fmt.Print(10 * rem / n)
		rem = (10 * rem) % n
	}
	fmt.Println()
}
