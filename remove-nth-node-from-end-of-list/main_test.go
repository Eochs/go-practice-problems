package main

import (
	"strconv"
	"testing"
)

type testobjs struct {
	l        ListNode
	expected string
}

var tests = []testobjs{
	{
		ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		"7->0->8",
	},
	{
		ListNode{0, &ListNode{1, nil}},
		"0->2->2",
	},
	{
		ListNode{1},
		"0->1",
	},
	{
		ListNode{9, &ListNode{9, nil}},
		"0->0->1",
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, stuff := range tests {
		result := AddTwoNumbers(&stuff.l1, &stuff.l2)
		if result.String() != stuff.expected {
			t.Errorf("Expected %s, got %s", stuff.expected, result.String())
		}
	}
}

func (l *ListNode) String() string {
	acc := ""
	for l.Next != nil {
		acc += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}
	return acc + strconv.Itoa(l.Val)
}
