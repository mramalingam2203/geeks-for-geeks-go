// Write a golnag code m to implement a queue using an array. Programs should contain functions for inserting elements into the queue, displaying queue elements, and checking whether the queue is empty or not
// queue using an array

package main

import (
	"fmt"
)

const MAX_SIZE int = 100

var queue = make([]int, MAX_SIZE)

//var queue [100]int

var front, rear int = -1, -1

func enQueue(element int) {
	if rear == MAX_SIZE-1 {
		fmt.Println("queue is full. cannot enqueue")
	} else {
		if front == -1 {
			front = 0
		}
		rear++
		queue[rear] = element
		fmt.Println("enqueued successfully")
	}
}

func deQueue() {
	if front == -1 || front > rear {
		fmt.Println("queue is empty. cannot enqueue")
	} else {
		fmt.Println("%d enqueued", queue[front])
		front++
	}
}

func isEmpty() bool {
	if front == -1 || front > rear {
		return true
	} else {
		return false
	}
}

func display() {
	if isEmpty() == true {
		fmt.Println("Queue is empty.\n")
	} else {
		printf("Queue elements: ")
		for i := front; i <= rear; i++ {
			fmt.Println("%d ", queue[i])
		}
		fmt.Println()
	}
}

func main() {
	enQueue(10)
	enQueue(20)
	enQueue(30)
	display()

	dequeue()
	display()

	dequeue()
	dequeue()
	display()

}
