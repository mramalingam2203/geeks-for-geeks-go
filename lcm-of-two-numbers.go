// https://www.geeksforgeeks.org/program-to-find-lcm-of-two-numbers/

package main

import (
	"fmt"
	_ "os"
)

// Function to return LCM of two numbers
func LCM(a int, b int) int {
	greater := max(a, b)
	smallest := min(a, b)

	for i := greater; ; i += greater {
		if i%smallest == 0 {
			return i
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
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(LCM(20, 80))
}
