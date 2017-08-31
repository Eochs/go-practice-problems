package main

import (
	"fmt"
	"strconv"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	runnerToFindEnd := dummyHead
	// move n Nodes in
	for ith := 1; ith <= n+1; ith++ {
		runnerToFindEnd = runnerToFindEnd.Next
	}
	runnerTrailingByN := dummyHead // follow n steps behind first runner
	for runnerToFindEnd != nil {
		runnerToFindEnd = runnerToFindEnd.Next
		runnerTrailingByN = runnerTrailingByN.Next
	}
	// next node is one to be removed

	runnerTrailingByN.Next = runnerTrailingByN.Next.Next

	return dummyHead.Next
}

func main() {
	// Given linked list: 1->2->3->4->5, and n = 2.
	// After removing the second node from the end, the linked list becomes 1->2->3->5.
	// l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	l := &ListNode{1, &ListNode{2, nil}}
	newlist := removeNthFromEnd(l, 2)
	fmt.Println(newlist.String())
}

func (l *ListNode) String() string {
	acc := ""
	for l.Next != nil {
		acc += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}
	return acc + strconv.Itoa(l.Val)
}
