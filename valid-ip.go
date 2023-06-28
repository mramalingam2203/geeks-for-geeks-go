// https://practice.geeksforgeeks.org/problems/validate-an-ip-address-1587115621

package main

import (
	"fmt"
	_ "os"
	"strconv"
	"strings"
)

func main() {
	ip := "255.255.a.1"
	fmt.Println(isValid(ip))
}

func isValid(s string) bool {
	delimiter := "."
	result := strings.Split(s, delimiter)
	var n int
	if len(result) > 4 {
		return false
	} else {
		for i := 0; i < len(result); i++ {
			n, _ = strconv.Atoi(result[i])
			fmt.Println(n)
			if !(n >= 0 && n <= 255) {
				return false
			}
		}
	}
	return true

}
