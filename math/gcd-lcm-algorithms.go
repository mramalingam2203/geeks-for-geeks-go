// https://www.geeksforgeeks.org/mathematical-algorithms/?ref=shm

package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 16}
	fmt.Println(gcd(23, 2))
	fmt.Println(lcm_of_array_elements(a))
	fmt.Println(Euclidean(a))
	fmt.Println(gcd_array_1(a))
	fmt.Println(getGCD(12, 36))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm_of_array_elements(arr []int) uint64 {
	result := uint64(arr[0])
	for i := range arr {
		result = ((uint64(arr[i])) * result) / (uint64(gcd(arr[i], int(result))))
	}

	return result
}

func Euclidean(arr []int) int {
	lcm := arr[0]
	var num1, num2, gcd_val int
	for i := range arr {
		num1 = lcm
		num2 = arr[i]
		gcd_val = gcd(num1, num2)
		lcm = (lcm * arr[i]) / gcd_val
	}
	return lcm
}

func gcd_array_iter(arr []int) int {
	result := 0

	for i := 0; i < len(arr); i++ {
		result = gcd(arr[i], result)
		if result == 1 {
			return 1
		}
	}

	return result
}

// iterative implementation
func getGCD(a, b int) int {
	for {
		if a > 0 && b > 0 {
			if a > b {
				a = a % b
			} else {
				b = b % a
			}
		}
		if a == 0 {
			return b
		}
	}
	return a

}
