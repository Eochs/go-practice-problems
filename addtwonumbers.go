package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// l1 := ListNode{2,
	// 	&ListNode{4,
	// 		&ListNode{3, nil}}}
	l1 := *convertIntToList(342)

	// l2 := ListNode{5,
	// 	&ListNode{6,
	// 		&ListNode{4, nil}}}
	l2 := *convertIntToList(465)

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Printf(l3.String())
	// fmt.Printf(convertIntToList(243).String())
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func convertListToInt(l *ListNode) int {
	acc := ""
	for l.Next != nil {
		acc = strconv.Itoa(l.Val) + acc
		l = l.Next
	}
	acc = strconv.Itoa(l.Val) + acc
	i, _ := strconv.Atoi(acc)
	return i
}

func convertIntToList(i int) *ListNode {
	s := strconv.Itoa(i) // list representation will be reverse of int
	fmt.Printf("int passed in: %s\n", s)
	headPointer := &ListNode{}

	for ix, c := range strings.Split(s, "") {
		i, _ := strconv.Atoi(c)
		if ix == 0 {
			m := &ListNode{i, nil}
			headPointer = m // moves to head
		} else {
			m := &ListNode{i, headPointer}
			headPointer = m // moves to head
		}
		fmt.Printf("testing: %v", headPointer.String())
	}
	return headPointer
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return convertIntToList(convertListToInt(l1) + convertListToInt(l2))
}

// Test output
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	acc := ""
	for l.Next != nil {
		acc += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}
	return acc + strconv.Itoa(l.Val) + "\n"
}

// func makeArrIntoList(arr []int) *ListNode {

// }
