// https://www.geeksforgeeks.org/the-stock-span-problem/?ref=lbp

// The STOCK-SPAN problem

package main

import "fmt"

func main() {
	stock_prices := []int{80, 60, 70, 60, 75, 85}
	fmt.Println(naiveApproach(stock_prices))
	fmt.Println(stackApproach(stock_prices))

}

func naiveApproach(prices []int) []int {
	span := make([]int, len(prices))
	span[0] = 1

	for i := range prices {
		span[i] = 1
		for j := i - 1; j >= 0 && prices[i] >= prices[j]; j-- {
			span[i]++
		}
	}
	return span
}

func stackApproach(prices []int) []int {
	stack := NewStack(100)
	span := make([]int, len(prices))

	stack.Push(0)
	for i := 1; i < len(prices); i++ {
		if stack.IsEmpty() && prices[stack.top] <= prices[i] == false {
			stack.Pop()
		}
		if stack.IsEmpty() == true {
			span[i] = i + 1
		} else {
			span[i] = i - stack.top
		}
		stack.Push(i)
	}
	return span
}
