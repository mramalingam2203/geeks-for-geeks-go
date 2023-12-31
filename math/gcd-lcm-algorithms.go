// https://www.geeksforgeeks.org/mathematical-algorithms/?ref=shm

package main

import (
	"fmt"
	"math"
)

func main() {
	//a := []int{2, 4, 6, 8, 16}
	a := []int{1, 2, 3}

	fmt.Println(gcd(23, 2))
	fmt.Println(lcm_of_array_elements(a))
	fmt.Println(Euclidean(a))
	fmt.Println(getGCD(12, 36))
	fmt.Println(lcm_gcd_distributive(2, 3, 9))
	fmt.Println(gcd_fp(3.0, 4.5))
	//fmt.Println(gcdOfArray(a))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
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

		if a > b {
			a = a % b
		} else {
			b = b % a
		}
		if a > 0 && b > 0 == false {
			break
		}
	}

	if a == 0 {
		return b
	}
	return a
}

func gcdOfArray(arr []int) int {
	gcd := 2
	for i := range arr {
		gcd = getGCD(gcd, arr[i])
	}
	return gcd
}

func SteinsAlgorithm(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	return 0
}

// Given three integers x, y, z, the task is to compute the value of GCD(LCM(x,y), LCM(x,z)).

func lcm_gcd_distributive(x, y, z int) int {
	return gcd(lcm(x, y), lcm(x, z))
}

func gcd_fp(a, b float64) float64 {
	if a < b {
		return gcd_fp(b, a)
	}
	if math.Abs(b) < 1e-3 {
		return a
	} else {
		return gcd_fp(b, a-math.Floor(a/b)*b)
	}

}
