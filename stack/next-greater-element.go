// https://www.geeksforgeeks.org/next-greater-element/?ref=lbp

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printNGE(array)

}

func printNGE(arr []int) {
	stack := new(Stack)
	stack.push(arr[0])
	for i := 1; i < len(arr); i++ {
		if stack.isEmpty() {
			stack.push(arr[i])
			continue
		}
		fmt.Println(stack.top.data)

		for {
			if stack.top.data-arr[i] == true {
				fmt.Println(stack.top.data, " --> ", arr[i])
			}
		}

		stack.push(arr[i])

	}
}
