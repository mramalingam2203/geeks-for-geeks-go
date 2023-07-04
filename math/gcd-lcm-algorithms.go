// https://www.geeksforgeeks.org/mathematical-algorithms/?ref=shm

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 9, 10, 12, 15}
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
