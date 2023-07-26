// experiments with slices and testing

/*
- Sum of Elements: Write a function that takes a slice of integers as input and returns the sum of all elements in the slice.

- Find Maximum and Minimum: Create a function that takes a slice of integers and returns the maximum and minimum values in the slice.

- Reverse a Slice: Write a function that takes a slice of any data type and returns a new slice with the elements in reverse order.

- Filter Even Numbers: Create a function that takes a slice of integers and returns a new slice containing only the even numbers from the original slice.

- Merge Slices: Write a function that takes two slices of integers and merges them into a single sorted slice.

- Rotate Slice: Implement a function that takes a slice of any data type and rotates it to the left by a given number of positions.

- Remove Duplicates: Create a function that takes a slice of integers and removes any duplicates, returning a new slice with unique elements.

- Chunk a Slice: Write a function that takes a slice of any data type and a chunk size, and returns a slice of slices where each inner slice contains the given number of elements.

- Count Occurrences: Implement a function that takes a slice of strings and a target string. It should return the number of occurrences of the target string in the slice.

- Palindrome Check: Create a function that checks if a given slice of bytes is a palindrome (reads the same backward as forward).

*/

package main

import "fmt"

func main() {
	createSlices()
}

func createSlices() {
	// int_slice := make([]int, 3)
	// float_slice := make([]float32, 3)
	// bool_slice := make([]bool, 3)
	rows := 3
	cols := 3
	// any_slice := make([]interface{}, 3)
	any_2DSlice := make([][]interface{}, rows)
	// Initialize each row with values

	for i := 0; i < rows; i++ {
		any_2DSlice[i] = make([]interface{}, cols)
		for j := 0; j < cols; j++ {
			any_2DSlice[i][j] = (i*cols + j + 1) // Assigning values based on the position
		}
	}

	fmt.Println(any_2DSlice)
}
