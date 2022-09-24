package longest_palindromic_substring

/*
Given a string s, return the longest palindromic substring in s.

A string is called a palindrome string if the reverse of that string is the same as the original string.
https://leetcode.com/problems/longest-palindromic-substring/
*/

import (
	"testing"
)

func longestPalindromicSubstring(s string) string {

	startingIndex := 0
	longestPalidrome := ""
	for startingIndex < len(s) {
		for i := startingIndex + 1; i < len(s); i++ {
			if s[startingIndex] == s[i] && startingIndex != i {
				if i-startingIndex > len(longestPalidrome) && isPalidrome(s[startingIndex:i+1]) {
					longestPalidrome = s[startingIndex : i+1]
				}
			}
		}
		startingIndex++
	}

	return longestPalidrome
}

func isPalidrome(s string) bool {

	for i := 0; i < (len(s)/2)-1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestLongestPalindromicSubstring(t *testing.T) {

	testCases := []struct {
		s    string
		want string
	}{
		{s: "babad", want: "bab"},
		{s: "cbbd", want: "bb"},
	}

	for i, testCase := range testCases {
		got := longestPalindromicSubstring(testCase.s)

		if testCase.want != got {
			t.Errorf("Error on test case %d, expected %s got %s", i, testCase.want, got)
		}
	}
}
