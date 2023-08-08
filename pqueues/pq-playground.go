/* Priority Queues playground*/

package main

import "fmt"

func main() {
	h := NewMaxHeap()
	fmt.Println(h)
}

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
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
