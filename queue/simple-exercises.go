
// Write a golnag code m to implement a queue using an array. Programs should contain functions for inserting elements into the queue, displaying queue elements, and checking whether the queue is empty or not
// queue using an array

package main

import(
	"fmt"
)


MAX_SIZE const int = 100

queue = [MAX_SIZE]int
front := -1
rear := -1

func enQueue(element int){
	if rear == MAX_SIZE -1 {
		fmt.Println("queue is full. cannot enqueue")
	}else{
		if front == -1 {
			front = 0
		}
		rear ++
		queue[rear] = element
		fmt.Println("enqueued successfully")
	}
}

func enQueue( ){
	if front == -1 || front > rear{
		fmt.Println("queue is empty. cannot enqueue"
	}else{
		fmt.Println("%d enqueued", queue[front])
		front++
	}
}

func isEmpty() {
    if (front == -1 || front > rear) {
        return 1;
    } else {
        return 0;
    }
}

func  display() {
    if (isEmpty()) {
        fmt.Println("Queue is empty.\n");
    } else {
        printf("Queue elements: ")
        for (int i = front; i <= rear; i++) {
            fmt.Println("%d ", queue[i])
        }
        fmt.Println( )
    }
}

func main() {
    enqueue(10)
    enqueue(20)
    enqueue(30)
    display()
    
    dequeue()
    display()
    
    dequeue()
    dequeue()
    display()

}
