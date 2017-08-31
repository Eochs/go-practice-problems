package main

import "testing"

type testobjs struct {
	heights      []int
	expectedArea int
}

var tests = []testobjs{
	{[]int{2, 2}, 2},
	{[]int{1, 5, 7, 2}, 5},
	{[]int{2, 1, 3, 2, 1, 1, 1, 1, 1}, 8},
	{[]int{2, 1, 2}, 4},
	{[]int{1, 2, 2, 1}, 3},
}

func TestMaxArea(t *testing.T) {
	for _, pair := range tests {
		result := maxArea(pair.heights)
		if result != pair.expectedArea {
			t.Errorf("Expected %v, got %v", pair.expectedArea, result)
		}
	}
}
