package main

import "testing"

type testobjs struct {
	input          string
	expectedOutput string
}

var tests = []testobjs{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"pwwkew", "ww"},
	{"bb", "bb"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, pair := range tests {
		result := longestPalindrome(pair.input)
		if result != pair.expectedOutput {
			t.Errorf("Expected %v, got %v", pair.expectedOutput, result)
		}
	}
}
