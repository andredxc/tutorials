package golang

import (
	"reflect"
	"testing"
)

func TestReverseStringInPlace(t *testing.T) {

	testCases := []struct {
		input []string
		want  []string
	}{
		{input: []string{"a", "b", "c", "d"}, want: []string{"d", "c", "b", "a"}},
		{input: []string{"a", "b", "c"}, want: []string{"c", "b", "a"}},
	}

	var got []string
	for i, testCase := range testCases {
		got = ReverseStringInPlace(testCase.input)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Error on test case %d, want=%v, got %v", i, testCase.want, got)
		}
	}

}

func ReverseStringInPlace(content []string) []string {

	var pos1, pos2 int

	for i := 0; i < len(content)/2; i++ {

		pos1 = i
		pos2 = len(content) - 1 - i

		content[pos1], content[pos2] = content[pos2], content[pos1]
	}

	return content
}
