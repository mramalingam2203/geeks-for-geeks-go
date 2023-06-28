// https://practice.geeksforgeeks.org/problems/power-of-2-1587115620/1?page=3&sortBy=submissions

package main

import (
	"fmt"
	"math"
	"os"
)

func isPowerofTwo(n uint64) bool {
	if n < 0 || n > 1e18 {
		fmt.Println("Numbers outliers")
		os.Exit(0)
	}

	if !isIntegral(math.Log2(float64(n))) {
		return false
	}
	return true
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func main() {
	fmt.Println(isPowerofTwo(5))
}
