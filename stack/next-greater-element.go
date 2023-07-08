// https://www.geeksforgeeks.org/next-greater-element/?ref=lbp

package main


func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printNGE(array []int)

}

func printNGE(arr []int) {
	stack := new(Stack)
	stack.push(arr[0])
	for  i := 1; i < n; i++ {
		if stack.isEmpty() {
            stack.push(arr[i])
            continue
        }


	}
