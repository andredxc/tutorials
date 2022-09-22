package brackets

import "testing"

func TestCheckBrackets(t *testing.T) {

	var got bool

	testCases := []struct {
		input string
		want  bool
	}{
		{input: "[({})]", want: true},
		{input: "[({}]", want: false},
		{input: "([][]{}()({({[]})}))", want: true},
	}

	for i, testCase := range testCases {
		got = CheckBrackets(testCase.input)
		if got != testCase.want {
			t.Errorf("Error on test case index %d, expected %v got %v", i, testCase.want, got)
		}
	}
}
