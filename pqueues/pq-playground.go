/* Priority Queues playground*/

package main

import (
	"fmt"
)

func main() {

}

type MaxHeap struct {
	heap []int
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
