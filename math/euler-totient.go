// Euler Totient function

package main

import "fmt"

func main() {
	fmt.Println(naive(10))
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)

}

func naive(n int) int {
	for i := 2; i < n; i++ {
		count := 0
		if gcd(i, n) == 1 {
			count++
		}
		fmt.Println("Phi(", i, "):= ", count)
	}
	return count
}
