// https://www.geeksforgeeks.org/the-stock-span-problem/?ref=lbp

// The STOCK-SPAN problem

package main

func main(){

}

func naiveApproach (prices []int{
	span := make([]int, len(prices))
	span[0]	= 1

	for i := range prices{
		span[i] = 1
		for j:= i-1; j >= 0 && prices[i] >= prices[j]; j--{
			span[i]++
		}
	}

}

