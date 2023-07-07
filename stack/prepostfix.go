// https://www.geeksforgeeks.org/convert-infix-expression-to-postfix-expression/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(precedence('^'))
}

func precedence(char rune) int {
	switch {
	case char == '+' || char == '-':
		return 1
	case char == '*' || char == '/':
		return 2
	case char == '^':
		return 3
	}
	return -1

}
