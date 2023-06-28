// https://practice.geeksforgeeks.org/problems/power-of-2-1587115620/1?page=3&sortBy=submissions

package main

import (
	"fmt"
	"math"
	"os"
)

func isPowerofTwo_log(n uint64) bool {
	if n < 0 || n > 1e18 {
		fmt.Println("Numbers outliers")
		os.Exit(0)
	}

	if !isIntegral(math.Log2(float64(n))) {
		return false
	}
	return true
}

func isPowerofTwo_multiply(n uint64) bool {
	if n == 0 {
		return false
	}
	for {
		if n != 1 {
			n = n / 2
			if n%2 != 0 && n != 1 {
				return false
			}
		}
	}

	return true
}

func powerOfTwo_bitwise(n int) bool {
	return n && (!(n && (n - 1)))
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func main() {
	fmt.Println(isPowerofTwo_multiply(5))
}
