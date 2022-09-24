package longest_substring

/*
Given a string s, find the length of the longest substring without repeating characters.
https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

import (
	"testing"
)

func longestSubstring(s string) int {

	var longestSubstring string

	done := false
	startingIndex := 0
	for !done {

		charMap := make(map[byte]int)
		curSubstring := ""
		for i := startingIndex; i < len(s); i++ {
			if _, ok := charMap[s[i]]; !ok {
				// Character is not yet in the substring
				curSubstring += string(s[i])
				charMap[s[i]] = i
			} else {
				// Found repeating character, end of substring
				startingIndex++
				if len(curSubstring) > len(longestSubstring) {
					longestSubstring = curSubstring
				}
				if len(s)-startingIndex <= len(longestSubstring) {
					done = true
				}
				break
			}
		}
	}
	return len(longestSubstring)
}

func TestLongestSubstring(t *testing.T) {

	testCases := []struct {
		s    string
		want int
	}{
		{s: "abcabcbb", want: 3},
		{s: "bbbbb", want: 1},
		{s: "pwwkew", want: 3},
	}

	for i, testCase := range testCases {
		got := longestSubstring(testCase.s)

		if got != testCase.want {
			t.Errorf("Error on test index %d, expected %d and got %d\n", i, testCase.want, got)
		}
	}

}
