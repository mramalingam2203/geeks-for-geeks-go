// https://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/

package main

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	low := a[0]
	high := a[1]
	key = 7
	binarySearch(a, low, high, key)
}

func binarySearch(arr []int, less, more, to_find int) {
	if less > more {
		return -1
	}
	mid := (less + more)/2
	if to_find == arr[mid] {
		return arr[mid]
	}
	if to_find > mid{
		return binarySearch(arr, mid+1, more, to_find
	}else{
		return binarySearch(arr, mid+1, more, to_find
	}

}


	}
}
