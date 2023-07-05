// Euler Totient function

package main

import "fmt"

func main() {
	fmt.Println(naive(10))
}

func gcd(a, b) {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)

}

func naive(n int) int {
	count := 0
	for i := range n {
		if gcd(1, i) == 1 {
			count++
		}
	}
}
