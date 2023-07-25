// https://www.geeksforgeeks.org/sum-absolute-differences-pairs-given-array



package main


func main() {
	array := []int{1, 2, 3, 4}
	sumOfPairs(array)

}


func mag(a int){
	if a < 0{
		return -1*a
	}
	return a
}

func sumOfPairs{arr []int) int {
	for i := 0; i < len(arr); i++{
		for j := i; j < len(arr); j++{
			fmt.Println(arr[i], arr[j])
		}
	}
	return 0
}