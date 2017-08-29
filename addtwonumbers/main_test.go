package main

import (
	"strconv"
	"testing"
)

type testobjs struct {
	l1       ListNode
	l2       ListNode
	expected string
}

var tests = []testobjs{
	{
		ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		"7->0->8",
	},
	{ // when one list longer than the other
		ListNode{0, &ListNode{1, nil}},
		ListNode{0, &ListNode{1, &ListNode{2, nil}}},
		"0->2->2",
	},
	{ // when one list is empty
		ListNode{},
		ListNode{0, &ListNode{1, nil}},
		"0->1",
	},
	{ // The sum could have an extra carry of one at the end
		ListNode{9, &ListNode{9, nil}},
		ListNode{1, nil},
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
