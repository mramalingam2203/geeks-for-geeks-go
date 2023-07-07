// https://www.geeksforgeeks.org/convert-infix-expression-to-postfix-expression/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(precedence('/'))
}

func precedence(char rune) int {
	switch operator := '+'; operator {
	case '+':
	case '-':
		return 1
	case '*':
	case '/':
		return 2
	case '^':
		return 3
	default:
		return -1
	}

	return 0
}
