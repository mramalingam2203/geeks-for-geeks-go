// https://www.geeksforgeeks.org/convert-infix-expression-to-postfix-expression/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(precedence('^'))
	fmt.Println(isOperator('*'))
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

func infixToPostfix(infix string) (string, error) {
	length = len(infix)
	infixArray := strings.Fields(infix)
	var postfix string

	for i:=0; i  <length; i++ {
		if infixArray[i] == ' ' || infixArray[i] = '\t'{
			continue
		}
}

// Function to check if the scanned character
// is an operator
func isOperator(char rune) bool {

	return char == '+' || char == '-' || char == '*' || char == '/' || char == '^'
}
