// https://www.geeksforgeeks.org/string-data-structure/?ref=shm

package main

/* EASY

- Count strings with consecutive 1’s
- Generate all binary strings from given pattern
- Add n binary strings
- Divide large number represented as string
- Program to find Smallest and Largest Word in a String
- Count number of equal pairs in a string
- Camel case of a given sentence
- Second most repeated word in a sequence
- Print all possible strings that can be made by placing spaces
- Check if all levels of two trees are anagrams or not
*/

func main() {

	CountStringsWithConsecutveOnes(3)
}

func CountStringsWithConsecutveOnes(n int) int{
	// Given a number n, count number of n length strings with consecutive 1’s in them.
	var a[n], b[n] int
	a[0], b[0] = 1,1
	for i:=0; i < n; i++ {
		a[i] = a[i - 1] + b[i - 1]
        b[i] = a[i - 1]
	}

	return (1 << n) - a[n - 1] - b[n - 1]
}
