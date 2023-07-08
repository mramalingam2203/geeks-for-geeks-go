package main

import (
	"errors"
	"fmt"
)

// Stack definition start

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

// Returns an initialized stack
func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

// Returns an new stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Returns the size of the size
func (stack *Stack) Size() int {
	return stack.top + 1
}

// Stack is full when top is equal to the last index
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

// Stack is empty when top is equal to -1
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	stack.top--
	return temp, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	return temp, nil
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

// Stack definitions over

func maxArray(array []int) int {
	max := 0
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

// next greater frequency function
func NFG(array []int) {

	max := maxArray(array)
	freq := make([]int, max+1)

	// Calculating frequency of each element
	for i := 0; i < len(array); i++ {
		freq[array[i]]++
	}

	fmt.Println(freq)

	stack := NewStack(100)
	stack.Push(0)

	for i := 1; i < len(array); i++ {

	}

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	NFG(array) //

}
