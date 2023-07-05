// Prime Numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(naive(31))
	fmt.Println(efficient(20))

}

func naive(n int) bool {

	if n == 1 || n == 2 {
		return true
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// is prime or not
func efficient(n int) bool {
	// Corner case
	if n <= 1 {
		return false
	}

	// Check from 2 to square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
