// Linked list playground using golang

package main

import "fmt"

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

	// Creating nodes
	node1 := &Node{data: "A"}
	node2 := &Node{data: "B"}
	node3 := &Node{data: "C"}

	// Linking nodes
	node1.next = node2
	node2.next = node3

	// Traversing the linked list
	currentNode := node1
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
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
