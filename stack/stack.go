// some gfg exercies using stacks

package main

type Stack struct {
	top int 
	capacity uint
	array []interface
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