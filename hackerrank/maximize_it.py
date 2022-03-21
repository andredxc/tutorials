#!/usr/bin/python3
# Enter your code here. Read input from STDIN. Print output to STDOUT

"""
TODO:
- Unit tests

Both can be acomplished by inheriting the InputReader class and overriding key methods.
"""
import abc

def main():

	inputReader = StdinReader()
	result = maximize_it(inputReader)
	print('Result: %d\n' % result)

def maximize_it(inputReader): 

	k, m = inputReader.readKMValues()
	kLists = inputReader.readKLists(k)
	print("m=%d; k=%d\n" % (m, k))
	for i in range(len(kLists)):
		print("%d %s\n" % (len(kLists[i]), kLists[i]))

	xValues = findHighestValuesInEachList(kLists)
	result = s(xValues, m)
	return result

def findHighestValuesInEachList(valuesLists):

	highestModules = []
	for values in valuesLists:
		highest = 0
		for value in values:
			if abs(value) > highest:
				highest = abs(value)
		highestModules.append(highest)
	return highestModules

def s(xValues, m):
	return sum([f(x) for x in xValues]) % m

def f(x):
    return x**2

class InputReader(metaclass=abc.ABCMeta):

	def readKMValues(self):
		print('K M:')
		line = self.readLine()
		line = line.split(' ')
		k = int(line[0])
		m = int(line[1])
		return k, m

	def readKLists(self, k):
		print('K values:')
		kValues = []
		for i in range(k):
			line = self.readLine()
			strValues = line.strip().split(' ')
			if StdinReader.isKLineValid(strValues):
				values = [int(j) for j in strValues[1:]]
				kValues.append(values)
			else:
				exit('line=\'%s\' is invalid' % line)
		return kValues

	@abc.abstractmethod
	def readLine(self):
		"""Reads a line from some input source"""

	@staticmethod
	def isKLineValid(values):
		length = int(values[0])
		if length != len(values)-1:
			return False
		else:
			return True

class StdinReader(InputReader):

	def readLine(self):
		# This allows the class to be unit tested by overriding this method
		return input()
		
class CsvReader(StdinReader):

	def __init__(self):
		super(StdinReader, self)
		self.inputFile = './testcases/input.csv'
		self.nextLineToRead = 0

	def readLine(self):
		file = open(self.inputFile, 'r')
		lines = file.readlines()
		line = ''
		try:
			if self.nextLineToRead > len(lines):
				raise Exception()
			
			line = lines[self.nextLineToRead]
			self.nextLineToRead += 1
		finally:
			file.close()
		return line

	def readExpectedResult(self):
		file = open(self.inputFile, 'r')
		lines = file.readlines()
		resultLine = lines[-1]
		return int(resultLine.strip())

if __name__ == "__main__":
    main()

