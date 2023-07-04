// https://www.geeksforgeeks.org/program-to-print-pyramid-pattern/

package main

import (
	"fmt"
)

func main() {
	printPyramid(5)
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
