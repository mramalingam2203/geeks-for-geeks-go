// Prime Numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(naive(31))
	fmt.Println(efficient(20))
	sundaram(203)

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

func sundaram(n int) {
	// Since we want primes smaller than n, we reduce n to half
	nNew := (n - 1) / 2

	// This array is used to separate numbers of the form i+j+2ij
	// from others where  1 <= i <= j
	marked := make([]bool, nNew+1)

	// Main logic of Sundaram.  Mark all numbers of the
	// form i + j + 2ij as true where 1 <= i <= j
	for i := 1; i <= nNew; i++ {
		for j := i; i+j+2*i*j <= nNew; j++ {
			marked[i+j+2*i*j] = true
		}
	}

	// Since 2 is a prime number
	if n > 2 {
		fmt.Print("2", " ")
	}

	// Print other primes. Remaining primes are of the form
	// 2*i + 1 such that marked[i] is false.
	for i := 1; i <= nNew; i++ {
		if marked[i] == false {
			fmt.Print(2*i+1, " ")
		}
	}
}
