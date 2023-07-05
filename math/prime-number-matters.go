// Prime Numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(naive(31))
	fmt.Println(efficient(20))
	segmented_sieve(23, 93)
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

func sundaram(n int) []int {
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

	result := []int{2}

	// Print other primes. Remaining primes are of the form
	// 2*i + 1 such that marked[i] is false.
	for i := 1; i <= nNew; i++ {
		if marked[i] == false {
			result = append(result, 2*i+1)
		}
	}

	return result
}

func filterSlice(slice, filter []int) []int {
	filteredSlice := make([]int, 0)
	filterSet := make(map[int]bool)

	// Create a set of filter values for efficient lookup
	for _, val := range filter {
		filterSet[val] = true
	}

	// Iterate over the original slice and add elements to the filtered slice if they are in the filter set
	for _, val := range slice {
		if !filterSet[val] {
			filteredSlice = append(filteredSlice, val)
		}
	}

	return filteredSlice
}

func segmented_sieve(low, high int) {
	arr_1 := sundaram(low)
	arr_2 := sundaram(high)
	result := filterSlice(arr_2, arr_1)
	fmt.Println(result)

}
