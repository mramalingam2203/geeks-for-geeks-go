# LINKED LISTS

class Node:
	def __init__(self, next = None, data = None):
		self.data = data
		self.next = next

	def getData(self):
		return self.data

	def setData(self,data):
		self.data = data

	def setNext(self, next):
		self.next = next

	def getNext(self):
		return self.next


class LinkedList(object):
	def __init__(self, node = None):
		self.length = 0
		self.head = node

	def lenth(self):
		current = self.head
		count = 0
		while current != None:
			count = count + 1
			current = current.next
		return count


	def traverseList(self):
		curr = self.head
		while curr != None:
			print(curr.getData())
			curr  = curr.next


	def insertAtBeginning(self, node):
		if self.head == None:
			self.head = node
		else:
			node.next = self.head
			self.head = node

	def printList(self):
		curr = self.head
		while curr != None:
			print(curr.data)
			curr = curr.next

	def insertAtEnd(self, node):
		if self.head == None:
			self.head = node
		else:
			current = self.head
			while current.next != None:
				current = current.next
			current.next = node

	def insertAtMiddle(self, node):
		if self.head == None:
			self.head = node
		else:
			pos = 0

			curr = self.head
			while pos <= self.length//2 :
				pos += 1
				curr = curr.next
			temp = curr.next
			curr.next = node
			node.next = temp
		self.length += 1 
		print(self.length)

	
	"""
	def insertAtGivenPosition(n, node):
		if self.head == None:
			self.head = none
		else:
			current = self.head
			while current.next != None
				if  k < n:
					current = current.next
					k += 1
				current.next = node
		pass
	"""
	
	def deleteFromBeginning(self):
		if self.head == None:
			return
		else:
			self.head = self.head.next
			self.length -= 1

	
	def deleteFromEnd():
		if self.head == None:
			return
		else:
			current = self.head
			while current.next != None:
				current = current.next

	
	def deleteAnIntermediateNode():
	
		pass
	
	def deleteWholeList():
		pass
	
class DLLNOde():
	def __init__(self, prev,  next, data):
		self.prev = prev
		self.next = next
		self.data = data
	def setPrev(self, prev):
		self.prev = prev
	def setNext(self, next):
		self.next = next
	def setData(self,data):
		self.data = data
	def getPrev(self):
		return self.prev
	def getNext(self):
		return self.next
	def getData(self):
		return self.data
	def hasPrev(self):
		if self.prev != None:
			return true
		return false

	def hasNext(self):
		if self.next != None:
			return true
		return false



class DoublyLinkedList(object):
	def __init__(self, node):
		self.head = node

	def insertAtBeginning(self, node):
		if self.head == None:
			self.head =  node
		else:
			curr = self.head
			self.head = node
			self.head.next = curr
			
	def insertAtEnd(self, node):
		pass
	def isnertAtSpot(self, mode, k):
		pass
	def deleteFromBeginning(self, node):
		pass
	def deleteFromEnd(self, end):
		pass




def main():
	node1 = Node(None,1)
	node2 = Node(None,2)
	node3 = Node(None, 3)
	node4 = Node(None, 4)

	print(node1.getData(), node2.getData(), node3.getData())

	node1.setNext(node2)
	node2.setNext(node3)
	node3.setNext(node4)

	ll = LinkedList(node1)
	print(ll.lenth())

	node5 = Node(None, 5)
	node6 = Node(None, 6)

	#ll.insertAtBeginning(node5)
	#ll.insertAtEnd(node6)
	#ll.traverseList()
	ll.insertAtMiddle(node3)
	print(ll.lenth())
	#ll.traverseList()


	#print(ll.head.getData())
	#print(ll.lenth())

	dll = DoublyLinkedList(node1)
	dll.insertAtBeginning(node2)


if __name__ == main():
	main()


