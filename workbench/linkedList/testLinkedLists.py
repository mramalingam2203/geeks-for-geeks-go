import unittest
from io import StringIO
import sys

# Assume your class code is in a file named `linkedlist.py`
# from linkedlist import Node, LinkedList

class Node:
    def __init__(self, next=None, data=None):
        self.data = data
        self.next = next

    def getData(self):
        return self.data

    def setData(self, data):
        self.data = data

    def setNext(self, next):
        self.next = next

    def getNext(self):
        return self.next


class LinkedList(object):
    def __init__(self, node=None):
        self.length = 0
        self.head = node

    def length_of_list(self):
        current = self.head
        count = 0
        while current is not None:
            count += 1
            current = current.next
        return count

    def traverseList(self):
        result = []
        curr = self.head
        while curr is not None:
            result.append(curr.getData())
            curr = curr.next
        return result

    def insertAtBeginning(self, node):
        if self.head is None:
            self.head = node
        else:
            node.next = self.head
            self.head = node

    def insertAtEnd(self, node):
        if self.head is None:
            self.head = node
        else:
            current = self.head
            while current.next is not None:
                current = current.next
            current.next = node

    def insertAtMiddle(self, node, pos):
        if pos > self.length | pos <0:
            return None
        elif pos == 0:
            self.insertAtBeginning(node)
        else :
            count = 1
            curr = self.head
            while count < pos-1:
                count += 1
                curr = curr.next
            node.next = curr.next
            curr.next = node
            self.length += 1



class TestLinkedList(unittest.TestCase):

    def test_node_operations(self):
        n = Node(None, 10)
        self.assertEqual(n.getData(), 10)
        n.setData(20)
        self.assertEqual(n.getData(), 20)
        next_node = Node(None, 30)
        n.setNext(next_node)
        self.assertEqual(n.getNext().getData(), 30)

    def test_insert_beginning(self):
        ll = LinkedList()
        n1 = Node(None, 1)
        n2 = Node(None, 2)
        ll.insertAtBeginning(n1)
        ll.insertAtBeginning(n2)
        self.assertEqual(ll.head.getData(), 2)
        self.assertEqual(ll.head.getNext().getData(), 1)

    def test_insert_end(self):
        ll = LinkedList()
        n1 = Node(None, 1)
        n2 = Node(None, 2)
        ll.insertAtEnd(n1)
        ll.insertAtEnd(n2)
        self.assertEqual(ll.head.getNext().getData(), 2)

    def test_insert_middle(self):
        n1 = Node(None, 1)
        n2 = Node(None, 2)
        n3 = Node(None, 3)
        ll = LinkedList(n1)
        ll.insertAtEnd(n2)
        ll.length = 2  # simulate actual length
        ll.insertAtMiddle(n3, 1)
        self.assertEqual(ll.head.getNext().getData(), 3)

    def test_length_of_list(self):
        n1 = Node(None, 1)
        n2 = Node(None, 2)
        n1.setNext(n2)
        ll = LinkedList(n1)
        self.assertEqual(ll.length_of_list(), 2)

    def test_traversal(self):
        n1 = Node(None, 1)
        n2 = Node(None, 2)
        n3 = Node(None, 3)
        n1.setNext(n2)
        n2.setNext(n3)
        ll = LinkedList(n1)
        self.assertEqual(ll.traverseList(), [1, 2, 3])


if __name__ == "__main__":
    unittest.main()
