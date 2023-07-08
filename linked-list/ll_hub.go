package main

import "fmt"

func main() {
	//demo()

	ll := LinkedList{}

	ll.insertAtBeginning(1)
	ll.insertAtBeginning(2)
	ll.insertAtEnd(3)
	ll.insert(2, 4)
	ll.insertAtEnd(5)
	ll.insertAtEnd(3)
	ll.insertAtEnd(5)
	ll.insertAtEnd(3)
	ll.insert(3, 5)
	ll.insertAtEnd(3)
	ll.insertAtEnd(3)
	ll.insertAtEnd(3)
	ll.insert(5, 4)

	//ll.findMiddleElement()
	findGivenKey(ll, 3)
	findNthNode(ll, 4)

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

func (ll *LinkedList) findMiddleElement() {
	i := 0
	current := ll.head
	fmt.Println(ll.display())
	for {
		if i == ll.length()/2 {
			fmt.Println(current.data)
			break
		}
		current = current.next
		i++
	}

}

// https://www.geeksforgeeks.org/write-a-function-that-counts-the-number-of-times-a-given-int-occurs-in-a-linked-list/?ref=lbp
// find number of times a given key occurs in a linked list

func findGivenKey(ll LinkedList, key int) {
	count := 0
	current := ll.head
	for i := 0; i < ll.length(); i++ {
		if current.data == key {
			count++
		}
		current = current.next
	}
	fmt.Println(count)
}

// https://www.geeksforgeeks.org/check-if-a-linked-list-is-circular-linked-list/?ref=lbp
// check for a circular list

func isCircularLL(ll LinkedList) bool {

	if ll.head == nil {
		return false
	}

	current := ll.head
	for current != nil {
		current = current.next
	}
	if current == ll.head {
		return true
	}

	return false

}

// https://www.geeksforgeeks.org/data-structures/linked-list/singly-linked-list/?ref=ghm

func checkIfEqual(ll1 LinkedList, ll2 LinkedList) bool {

	if ll1.head == nil || ll2.head == nil {
		return false
	}

	current_1 := ll1.head
	current_2 := ll2.head

	for {
		if current_1 != nil && current_2 != nil {
			if current_1.data == current_2.data {
				current_1 = current_1.next
				current_2 = current_2.next
			} else {
				return false
			}
		}
	}

	return true
}

// https://www.geeksforgeeks.org/write-a-function-to-get-nth-node-in-a-linked-list/

func findNthNode(ll LinkedList, n int) int {
	count := 0
	current := ll.head
	ll.display()

	for i := 0; i < ll.length(); i++ {
		if count == n {
			fmt.Println(current.data)
		}
		count++
		current = current.next
	}

	return -1
}
