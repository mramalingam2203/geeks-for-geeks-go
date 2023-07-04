// https://www.geeksforgeeks.org/mathematical-algorithms/?ref=shm

package main

import "fmt"

func main() {
	a := []int{250, 780, 390, 427, 963}
	fmt.Println(gcd(23, 2))
	fmt.Println(lcm_of_array_elements(a))
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
