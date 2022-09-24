package two_sum

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

https://leetcode.com/problems/two-sum/
*/

import (
	"reflect"
	"testing"
)

func twoSum(numbers []int, target int) []int {

	var remainder int

	numberMap := make(map[int]int, len(numbers))

	// Add all numbers to the number map for easy access
	for i, num := range numbers {
		if _, ok := numberMap[num]; !ok {
			numberMap[num] = i
		}
	}

	// Find the remainder in the map
	for i, num := range numbers {
		remainder = target - num
		if j, ok := numberMap[remainder]; ok {
			if i != j {
				// Can't use the same element twice
				if i < j {
					return []int{i, j}
				} else {
					return []int{j, i}
				}
			}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {

	test_cases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{numbers: []int{1, 2, 3, 4, 5}, target: 5, want: []int{0, 3}},
		{numbers: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{numbers: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{numbers: []int{3, 3}, target: 6, want: []int{0, 1}},
		{numbers: []int{3, 3}, target: 7, want: nil},
	}

	for i, testCase := range test_cases {
		got := twoSum(testCase.numbers, testCase.target)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test case index %d error, expected %v, got %v\n", i, testCase.want, got)
		}
	}
}
