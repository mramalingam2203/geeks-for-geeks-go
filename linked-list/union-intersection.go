// https://www.geeksforgeeks.org/union-and-intersection-of-two-linked-lists/

package main()

type Node struct {
	data interface{}
	next *Node
}


func main() {



}

func union-intersection(head1 *Node, head2 *Node) {

	head_u := *Node
	head_i := *Node

	current_1 := head1
	current_2 := head2

	for{
		if current_1 != nil || current_2 != nil {
			if current_1.data == current_2.data {
				head_i.data = current_1.data
				head_i = head_1.next
				current_1 = current_1.next
				current_2 = current_2.next
			}
		fmt.Println("found union and intersection lists")
		break
			
	}

}