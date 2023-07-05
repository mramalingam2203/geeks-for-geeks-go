// https://www.geeksforgeeks.org/check-if-a-given-number-can-be-represented-in-given-a-no-of-digits-in-any-base/

package main

import (
	"fmt"
)

func main() {
}

func checkBase(m, n int) bool {
	for i := 2; i <= 32; i++ {
		if m < i && n == 1 {
			return true
		}

	}

}
