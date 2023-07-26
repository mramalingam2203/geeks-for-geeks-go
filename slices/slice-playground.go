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

package sliceplayground

func main() {
	createSlices()
}

func createSlices() {
	int_slice := make([]int{1, 2, 3}
	float_slice := make([]float{1.0,2,0, -3.0, -4.0}
	bool_slice := make([]bool{true, false, false, true, false}
	int_int_slice := make([][]int) {1,2,3,4}, {1,2,3}, {0, 1, 2, 3, 4,5}}
	}