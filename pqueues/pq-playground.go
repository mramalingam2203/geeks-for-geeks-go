/* Priority Queues playground*/

package main

import (
	"fmt"
)

func main() {
	pq := NewMaxHeap()

	pq.Push(5)
	pq.Push(10)
	pq.Push(3)
	pq.Push(8)

	fmt.Println("Max element:", pq.Peek()) // Output: Max element: 10

	fmt.Println("Popped:", pq.Pop()) // Output: Popped: 10
	fmt.Println("Popped:", pq.Pop()) // Output: Popped: 8
	fmt.Println("Length now:", pq.Len())

	/* find the k-th largest element in an array */
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	k := 3
	kthLargest := findKthLargest(nums, k)
	fmt.Printf("The %dth largest element is: %d\n", k, kthLargest)

}

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Push(val int) {
	h.heap = append(h.heap, val)
	h.heapifyUp(len(h.heap) - 1)
}

func (h *MaxHeap) Len() int {
	return len(h.heap)

}

func (h *MaxHeap) Pop() int {
	if len(h.heap) == 0 {
		return -1 // Handle empty heap case
	}

	max := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)

	return max
}

func (h *MaxHeap) Peek() int {
	if len(h.heap) == 0 {
		return -1 // Handle empty heap case
	}
	return h.heap[0]
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[parentIndex] >= h.heap[index] {
			break
		}
		h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
		index = parentIndex
	}

}

func (h *MaxHeap) heapifyDown(index int) {
	for {

		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex < len(h.heap) && h.heap[leftChildIndex] > h.heap[largestIndex] {
			largestIndex = leftChildIndex
		}

		if rightChildIndex < len(h.heap) && h.heap[rightChildIndex] > h.heap[largestIndex] {
			largestIndex = rightChildIndex
		}

		if largestIndex == index {
			break
		}

		h.heap[index], h.heap[largestIndex] = h.heap[largestIndex], h.heap[index]
		index = largestIndex
	}
}

/*
Kth Largest Element in an Array: Given an array of integers, find the kth largest element. You can solve this problem using a min-heap. Keep the heap's size limited to k, and as you traverse the array, insert elements into the heap. If the heap size exceeds k, remove the smallest element. At the end, the heap's top element will be the kth largest.
*/

func findKthLargest(nums []int, k int) int {
	pq := NewMaxHeap()
	for _, num := range nums {
		pq.Push(num)
		if pq.Len() > k {
			pq.Pop()
		}
	}

	return pq.heap[0]
}

/*
Given k sorted linked lists, merge them into a single sorted linked list. This problem can be solved using a min-heap. You can insert the first element from each list into the heap and then keep popping the smallest element and insert the next element from the corresponding list until all lists are exhausted.
*/
func mergeKSortedArrays(arrays [][]int) []int {
	pq := NewMaxHeap()

	// Initialize the heap with the last element of each array
	for _, arr := range arrays {
		if len(arr) > 0 {
			pq.Push(arr[len(arr)-1])
		}
	}

	merged := []int{}
	for pq.Len() > 0 {

	}

}
