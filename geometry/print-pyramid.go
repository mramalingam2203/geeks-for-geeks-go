// https://www.geeksforgeeks.org/program-to-print-pyramid-pattern/

package main

import (
	"fmt"
)

func main() {
	printPyramid(5)
	printHourglassPattern(5)
}

func printPyramid(n int) {
	for i := 1; i < n; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}

	// For printing the lower part of pyramid
	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(" * ")
		}
		fmt.Println()
	}

}

func printHourglassPattern(n int) {

	// for loop for printing
	// upper half
	for i := 1; i <= n; i++ {

		// printing i spaces at
		// the beginning of each row
		for k := 1; k < i; k++ {
			fmt.Print(" ")
		}

		// printing i to rows value
		// at the end of each row
		for j := i; j <= n; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

	// for loop for printing lower half
	for i := n - 1; i >= 1; i-- {

		// printing i spaces at the
		// beginning of each row
		for k := 1; k < i; k++ {
			fmt.Print(" ")
		}

		// printing i to rows value
		// at the end of each row
		for j := i; j <= n; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}

}
