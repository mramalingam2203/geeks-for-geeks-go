// https://www.geeksforgeeks.org/write-a-c-function-to-print-the-middle-of-the-linked-list/?ref=lbp

package main

import "fmt"

type ListNode struct { // defines a ListNode in a linked list
	data interface{} // the datum
	next *ListNode   // pointer to the next ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) insertAtBeginning(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
	return
}

func (ll *LinkedList) insertAtEnd(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
	return
}

// Insert adds an item at position i
func (ll *LinkedList) insert(position int, data interface{}) error {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	newNode := &ListNode{data, nil}

	var prev, current *ListNode
	prev = nil
	current = ll.head

	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}

	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}

	ll.size++
	return nil
}

func (ll *LinkedList) deleteFront() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *LinkedList) deleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteLast: List is empty")
	}
	var prev *ListNode
	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.data, nil
}

// delete removes an element at position i
func (ll *LinkedList) delete(position int) (interface{}, error) {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	// Complete this method
	var prev, current *ListNode
	prev = nil
	current = ll.head

	pos := 0

	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}

		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}

func (ll *LinkedList) display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// with extra size field
func (ll *LinkedList) length() int {
	return ll.size
}

// length returns the linked list size
func (ll *LinkedList) length2() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func main() {
	demo()

}
func demo() {
	fmt.Println("Demoing the capabilites")
	ll := LinkedList{}

	fmt.Printf("insertAtBeginning: A\n")
	ll.insertAtBeginning("A")
	fmt.Printf("insertAtBeginning: B\n")
	ll.insertAtBeginning("B")
	fmt.Printf("insertAtEnd: C\n")
	ll.insertAtEnd("C")
	fmt.Printf("insert: D\n")
	ll.insert(2, "D")
	fmt.Printf("length: %d\n", ll.length())

	err := ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("delete\n")
	_, err = ll.delete(4)
	if err != nil {
		fmt.Printf("deleteError: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteFront\n")
	_, err = ll.deleteFront()
	if err != nil {
		fmt.Printf("deleteFront Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("length: %d\n", ll.length())

}
