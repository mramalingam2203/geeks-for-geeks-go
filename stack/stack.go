// some gfg exercies using stacks

package main

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

func main() {

}
