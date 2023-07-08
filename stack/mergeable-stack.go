// https://www.geeksforgeeks.org/create-mergable-stack/?ref=lbp

// some gfg exercies using stacks

package main

import "errors"

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

// Returns a new stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Stack is full when top is equal to the last index
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

// Stack is empty when top is equal to -1
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

// Returns the size of the size
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
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


func (stack1, stack2  *Stack){
	


}
