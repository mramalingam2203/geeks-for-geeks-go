import math
from functools import reduce
from operator import mul
import operator

a = [2, 4, 8, 3]
def subarray(a):
	n = len(a)
	subarrays = []
	for i in range(n):
		for j in range(i, n):
			subarray = []
			for k in range(i, j+1):
				subarray.append(a[k])
			print(subarray)


			subarrays.append(subarray)
	return subarrays




def subarray1(a):
	n = len(a)
	res = []

	for i in range(n):
		for j in range(i + 1, n + 1):
			res.append(a[i:j])
			print(res)

	#print(res)




def subsequence(a):
	res = [[]]

	for x in a:
		res += [s + [x] for s in res]
		print(res)
	


def subset(a):
	res = [[]]

	for x in a:
		res += [s + [x] for s in res]

	print(res)
	

# https://www.geeksforgeeks.org/dsa/check-if-subarray-with-given-product-exists-in-an-array/

# Check if subarray with given product exists in an array
def subarrayOfGivenProduct(a, PROD):
	n = len(a)
	subarrays = []
	for i in range(n):
		for j in range(i, n):
			subarray = []
			for k in range(i, j+1):
				subarray.append(a[k])

			if reduce(mul, subarray) == PROD:
				return True
	return False


# https://www.geeksforgeeks.org/dsa/array-subarray-subsequence-and-subset//
def subarrayOfGivenSizeAndSum(a, size, sum):
	n = len(a)
	for i in range(n):
		for j in range(i,n):
			subarray = []
			for k in range(i, j+1):
				subarray.append(a[k])
			if len(subarray) == size and reduce(operator.add, subarray) == sum:
				return True

	return False




"""
array = [1, 2, 3]
print(subarrayOfGivenProduct(array, 10))
array2 = [1, 4, 2, 10, 2, 3, 1, 0, 20]
print(subarrayOfGivenSizeAndSum(array2, 4, 18))
"""

array = [1, 1, 3, 3, 2, 4, 2, 3, 5, 6,5,6, 3, 4]
findFreq(array)