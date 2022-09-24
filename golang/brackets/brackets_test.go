package brackets

import (
	"fmt"
	"testing"
)

func CheckBrackets(line string) bool {

	// var stringRune string
	openers := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	bracketStack := []rune{}
	for i, c := range line {
		// stringRune = string(c)
		// fmt.Println(stringRune)

		if c == '(' || c == '{' || c == '[' {
			// Opening character
			bracketStack = append(bracketStack, c)
		} else if closer, ok := openers[c]; ok {
			// Closing character
			if len(bracketStack) == 0 || bracketStack[len(bracketStack)-1] != closer {
				return false
			}
			bracketStack = bracketStack[:len(bracketStack)-1]
		} else {
			// Unknown rune
			fmt.Println(i)
			return false
		}
	}

	return true
}

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
