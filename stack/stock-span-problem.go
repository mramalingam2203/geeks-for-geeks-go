// https://www.geeksforgeeks.org/the-stock-span-problem/?ref=lbp

// The STOCK-SPAN problem

package main

import "fmt"

func main() {
	stock_prices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(naiveApproach(stock_prices))
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
