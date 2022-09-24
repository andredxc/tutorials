package add_two_numbers

import (
	"strconv"
	"testing"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
https://leetcode.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var nodeValue1, nodeValue2, nodeSum int
	var prevNode, resultNode, firstResultNode *ListNode

	carryOver := 0
	for l1 != nil || l2 != nil || carryOver != 0 {

		nodeValue1 = 0
		if l1 != nil {
			nodeValue1 = l1.Val
			l1 = l1.Next
		}

		nodeValue2 = 0
		if l2 != nil {
			nodeValue2 = l2.Val
			l2 = l2.Next
		}

		nodeSum = nodeValue1 + nodeValue2 + carryOver
		carryOver = 0
		if nodeSum >= 10 {
			carryOver = int(nodeSum / 10)
			nodeSum = nodeSum % 10
		}

		resultNode = &ListNode{Val: nodeSum, Next: nil}

		if prevNode != nil {
			prevNode.Next = resultNode
		}

		if firstResultNode == nil {
			firstResultNode = resultNode
		}

		prevNode = resultNode
	}

	return firstResultNode
}

func TestAddTwoNumbers(t *testing.T) {

	testCases := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
		{l1: []int{0}, l2: []int{0}, want: []int{0}},
		{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, testCase := range testCases {
		got := addTwoNumbers(sliceToListNode(testCase.l1), sliceToListNode(testCase.l2))

		want := sliceToListNode(testCase.want)
		if linkedListToStr(want) != linkedListToStr(got) {
			t.Errorf("Error, expected %s, got %s\n", linkedListToStr(want), linkedListToStr(got))
		}
	}
}

func sliceToListNode(number []int) *ListNode {

	var firstNode, curNode, prevNode *ListNode

	for _, num := range number {

		if num >= 10 {
			panic("This algorithm does not support node values greater than 10!")
		}

		curNode = &ListNode{Val: num}

		if firstNode == nil {
			firstNode = curNode
		}

		if prevNode != nil {
			prevNode.Next = curNode
		}

		prevNode = curNode
	}

	return firstNode
}

func linkedListToStr(list *ListNode) string {

	listStr := "["
	for list != nil {

		listStr += strconv.Itoa(list.Val)
		// Add a coma to all except the last element
		if list.Next != nil {
			listStr += ", "
		}

		list = list.Next
	}

	listStr += "]"
	return listStr
}
