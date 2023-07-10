// Linked list playground using golang

package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type NodeDLL struct {
	data interface{}
	prev *NodeDLL
	next *NodeDLL
}

type DoublyLinkedList struct {
	head *NodeDLL
	tail *NodeDLL
}

func main() {

	var node1, node2, node3, node4, node5 *Node

	// Creating nodes
	node1 = &Node{data: "A"}
	node2 = &Node{data: "B"}
	node3 = &Node{data: "C"}
	node4 = &Node{data: "D"}
	node5 = &Node{data: "E"}

	// Linking nodes
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	// Traversing the linked list
	currentNode := node1
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

	fmt.Println("Middle node:", findMiddleNode(node1).data)
	//makeMiddleNodeHead(node1)
	deleteAlternateNode(node1)
	deleteList(node1)

	// Creating nodes
	node1 = &Node{data: 1}
	node2 = &Node{data: 2}
	node3 = &Node{data: 3}
	node4 = &Node{data: 4}
	node5 = &Node{data: 5}

	// Linking nodes
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	findSumOfLastNNodes(node1, 4)

}

func deleteList(head *Node) {
	current := head
	for current != nil {
		next := current.next
		current.next = nil
		current = next
	}
}

// To find the middle node of a linked list in Go, you can use the "slow and fast pointer" technique
func findMiddleNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// Given a singly linked list and a key, count the number of occurrences of the given key in the linked list

func countOccurrences(head *Node, key interface{}) int {
	count := 0
	currentNode := head

	for currentNode != nil {
		if currentNode.data == key {
			count++
		}
		currentNode = currentNode.next
	}

	return count
}

// To determine if a given linked list is a circular linked list in Go

func isCircularLinkedList(head *Node) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

// Given a circular linked list, count the number of nodes in it

func countNodesInCircularList(head *Node) int {
	if head == nil {
		return 0
	}

	count := 1
	currentNode := head.next

	for currentNode != head {
		count++
		currentNode = currentNode.next
	}

	return count
}

// To convert a singly linked list into a circular linked list

func convertToCircularList(head *Node) *Node {
	if head == nil {
		return nil
	}

	currentNode := head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// Make the last node point to the head node
	currentNode.next = head

	return head
}

// To exchange the first and last nodes in a circular linked list

func exchangeFirstAndLastNodes(head *Node) *Node {
	if head == nil || head.next == head {
		return head
	}

	// Find the last node
	lastNode := head
	for lastNode.next != head {
		lastNode = lastNode.next
	}

	// Swap the first and last nodes
	lastNode.next = head.next
	head.next = lastNode

	return lastNode
}

// To check if a linked list has a loop (cycle) or not
func hasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return false
}

// Given the head of a linked list. The task is to find if a loop exists in the linked list if yes then return the length of the loop in the linked list else return 0.

func findLoopLength(head *Node) int {
	if head == nil || head.next == nil {
		return 0
	}

	slow := head
	fast := head
	loopExists := false

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			loopExists = true
			break
		}
	}

	if !loopExists {
		return 0
	}

	count := 1
	slow = slow.next

	for slow != fast {
		slow = slow.next
		count++
	}

	return count
}

// To remove duplicate nodes from a sorted linked list in non-decreasing order
func removeDuplicates(head *Node) {
	if head == nil {
		return
	}

	currentNode := head
	for currentNode != nil && currentNode.next != nil {
		if currentNode.data == currentNode.next.data {
			currentNode.next = currentNode.next.next
		} else {
			currentNode = currentNode.next
		}
	}
}

func printList(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

// Make middle node head in a linked list
// https://www.geeksforgeeks.org/make-middle-node-head-linked-list/

func makeMiddleNodeHead(head *Node) *Node {
	if head == nil {
		fmt.Println("empty list")
		return nil
	}

	head = findMiddleNode(head)

	fmt.Println("After making middle node head : ")

	currentNode := head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

	//fmt.Println(head.data)
	return head

}

// Delete alternate nodes of a singly linked list
//https://www.geeksforgeeks.org/delete-alternate-nodes-of-a-linked-list/

// To remove duplicate nodes from a sorted linked list in non-decreasing order
func deleteAlternateNode(head *Node) {
	if head == nil {
		return
	}
	i := 0
	currentNode := head
	for currentNode != nil && currentNode.next != nil {
		if i%2 == 0 {
			currentNode.next = currentNode.next.next
		} else {
			currentNode = currentNode.next
		}
		i++
	}
	printList(head)
}

// https://www.geeksforgeeks.org/add-1-number-represented-linked-list/
// Add 1 to a number represented as linked list

func reverseList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

/*
func addOneToNumber(head *Node) *Node {
	reverseList(head)
	fmt.Println(head.data.(int))
	newNode = &Node{data: 0}

	while(head != NULL) //while both lists exist
	{
		// Calculate value of next digit in resultant list.
		// The next digit is sum of following things
		// (i) Carry
		// (ii) Next digit of head list (if there is a
		//     next digit)
		sum = carry + head.data

		// update carry for next calculation
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		// update sum if it is greater than 10
		sum = sum % 10

		// Create a new node with sum as data
		head.data = sum

		// Move head and second pointers to next nodes
		temp = head
		head = head.next
	}

	return head

}
*/

func findSumOfLastNNodes(head *Node, n int) {

	count := 0
	sum := 0

	if count != n {
		sum += head.data.(int)
		count++
		head = head.next
	}
	//fmt.Println(sum)
}
