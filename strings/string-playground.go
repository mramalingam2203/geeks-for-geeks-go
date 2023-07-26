// https://www.geeksforgeeks.org/string-data-structure/?ref=shm

package main

import (
	"fmt"
	"strings"
	_ "unicode"
)

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
	//fmt.Println(CountStringsWithConsecutveOnes_naive(5))
	//fmt.Println(CountStringsWithConsecutveOnes_bitwise(5))
	//array := []int{1, 0}
	// fmt.Println(getCombinations(array, 2))
	// str := "010??10?101"
	//PrintAllBinaryStrings(str, 0)
	str := "geeksforgeeks"
	//findLargestSmallestWord(str)
	//countPairs(str)
	//printAllPossibleSubstrings(str)
	secondMostRepeatedWordInASentence(str)
	//camelCase(str)
	firstNonRepeatingChar(str)
	kthNonRepeatingChar(str, 3)
	//lexicographicallyNext(str)
	palindromeSubstrings(str, 0, 10)
	minimumRotationsRequired(str)

}

func generateCombinations(arr []int, start, k int, combination []int, result *[][]int) {
	if k == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}

	for i := start; i <= len(arr)-k; i++ {
		combination = append(combination, arr[i])
		generateCombinations(arr, i+1, k-1, combination, result)
		combination = combination[:len(combination)-1]
	}
}

func getCombinations(arr []int, k int) [][]int {
	var result [][]int
	var combination []int

	generateCombinations(arr, 0, k, combination, &result)
	return result
}

func CountStringsWithConsecutveOnes_naive(n int) int {

	return 0

}

func CountStringsWithConsecutveOnes_bitwise(n int) int {
	// Given a number n, count number of n length strings with consecutive 1’s in them.
	a := make([]int, n)
	b := make([]int, n)

	a[0], b[0] = 1, 1
	for i := 1; i < n; i++ {
		a[i] = a[i-1] + b[i-1]
		b[i] = a[i-1]
	}

	return (1 << n) - a[n-1] - b[n-1]
}

// https://www.geeksforgeeks.org/generate-all-binary-strings-from-given-pattern/
/*
func PrintAllBinaryStrings(string s, index int) {
	if index == len(s)-1 {
		fmt.Println()

	}

	if s[index] == '?' {
		s[index] = '0'
		PrintAllBinaryStrings(s, index+1)

		s[index] = '1'
		PrintAllBinaryStrings(s, index+1)
	} else {
		PrintAllBinaryStrings(str, index+1)
	}
}
*/
func addBinaryString(s []string) {

}

func divideLargeNumber(s string, divisor int) {

}

func findLargestSmallestWord(s string) {
	words := strings.Fields(s)
	fmt.Println(len(words))
	var min, max int
	idx_min, idx_max := 0, 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) < min {
			min = len(words[i])
			idx_min = i
			fmt.Println(idx_min)
		}
		if len(words[i]) > max {
			max = len(words[i])
			idx_max = i
			fmt.Println(idx_max, idx_max)
		}
	}
	fmt.Println(words[idx_min], words[idx_max])
}

// Equal pars of a string
func countPairs(s string) int {

	wordCount := make(map[string]int)
	words := strings.Fields(s)

	// Iterate through each word and update the map
	for _, word := range words {
		// Convert the word to lowercase to make the counting case-insensitive
		word = strings.ToLower(word)
		wordCount[word]++
	}

	fmt.Println(wordCount)

	// Stores the answer
	ans := 0

	return ans

}

// https://www.geeksforgeeks.org/program-print-substrings-given-string/

func printAllPossibleSubstrings(s string) {

	for ln := 1; ln <= len(s); ln++ {
		for i := 0; i <= len(s)-ln; i++ {
			j := i + ln - 1
			for k := i; k <= j; k++ {
				fmt.Print(string(s[k]))
			}
			fmt.Println()
		}
	}
}

func secondMostRepeatedWordInASentence(s string) {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	// Iterate through each word and update the map
	for _, word := range words {
		// Convert the word to lowercase to make the counting case-insensitive
		word = strings.ToLower(word)
		wordCount[word]++
	}

	fmt.Println(wordCount)
}

// https://www.geeksforgeeks.org/camel-case-given-sentence/

// func camelCase(s string) {
// 	words := strings.Fields(s)
// 	for i := 0; i < len(words); i++ {
// 		words[i][0] = unicode.ToUpper(rune(words[i][0]))
// 	}
// }

// https://www.geeksforgeeks.org/given-a-string-find-its-first-non-repeating-character/
func firstNonRepeatingChar(s string) {
	wordCount := make(map[rune]int)

	// Iterate through each word and update the map
	for _, character := range s {
		// Convert the word to lowercase to make the counting case-insensitive
		wordCount[character]++
	}

	var k rune
	// Iterate through each word and update the map
	for _, character := range s {
		if wordCount[character] == 1 {
			k = character
			break
		}
	}

	fmt.Println(string(k))

}

// https://www.geeksforgeeks.org/given-a-string-find-its-first-non-repeating-character/
func kthNonRepeatingChar(s string, repeat int) {
	wordCount := make(map[rune]int)

	// Iterate through each word and update the map
	for _, character := range s {
		// Convert the word to lowercase to make the counting case-insensitive
		wordCount[character]++
	}

	var k rune
	count := 0
	// Iterate through each word and update the map
	for _, character := range s {
		if wordCount[character] == 1 && count != repeat {
			k = character
			count++
		}
	}
	fmt.Println(string(k))

}

// https://www.geeksforgeeks.org/lexicographically-next-string/

// func lexicographicallyNext(s string) {
// 	r := []rune(s)

// }

// https://www.geeksforgeeks.org/palindrome-substring-queries/
/*
Given a string and several queries on the substrings of the given input
string to check whether the substring is a palindrome or not.
*/

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func palindromeSubstrings(s string, l, r int) {
	fmt.Println(isPalindrome(s[l : r+1]))

}

// https://www.geeksforgeeks.org/minimum-rotations-required-get-string/

func minimumRotationsRequired(s string) {
	//count := 0
	temp := s + s
	fmt.Println(temp)
	fmt.Println(temp[2 : len(s)-1])

}
