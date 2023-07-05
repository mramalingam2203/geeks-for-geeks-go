// Prime Numbers

package main

import (
	"fmt"
)

func main() {
	fmt.Println(naive(32))
	fmt.Println(efficient(32))

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
func efficient(int n) {
	// Corner case
	if n <= 1 {
		return false
	}

	// Check from 2 to square root of n
	for i := 2; i <= int(Math.sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
		return true
	}
}
