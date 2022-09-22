package two_sum

import (
	"reflect"
	"testing"
)

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
