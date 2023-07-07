// https://www.geeksforgeeks.org/rearrange-positive-and-negative-numbers-publish/

package main

import "fmt"

func main() {
	a := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}
	rearrange(a)
}

func rearrange(a []int) {
	var tmp int
	for i := 0; i < len(a)-2; i += 2 {
		if a[i] >= 0 && a[i+1] < 0 {
			tmp = a[i]
			a[i] = a[i+1]
			a[i+1] = tmp
		}
	}
	fmt.Println(a)

}
