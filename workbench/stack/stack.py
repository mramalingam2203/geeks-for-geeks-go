class Stack:
	def __init__(self, Capacity = -1):
		self.top = -1
		self.Capacity = Capacity
		self.A = [None]*Capacity

	def push(self, data):
		if self.Capacity == self.top + 1:
			print('Stack Overfloe')
			return
		self.top += 1
		self.A[self.top] = data

	def pop(self):
	def peek(self):
	def isFull(self):
		return self.Capacity != -1
	def isEmpty(self):
		return self.Capacity == -1
