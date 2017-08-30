package main

import "testing"

type testobjs struct {
	given          string
	expectedLength int
}

var tests = []testobjs{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, pair := range tests {
		result := LengthOfLongestSubstring(pair.given)
		if result != pair.expectedLength {
			t.Errorf("Expected %v, got %v", pair.expectedLength, result)
		}
	}
}
