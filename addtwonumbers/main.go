package main

// ListNode represents a non-negative number.
// You are given two linked lists representing two non-negative numbers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers ...
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result
	pnode := l1
	qnode := l2

	carry := 0

	for pnode != nil || qnode != nil {

		tmp := carry

		if pnode != nil {
			tmp = tmp + pnode.Val
			pnode = pnode.Next
		}

		if qnode != nil {
			tmp = tmp + qnode.Val
			qnode = qnode.Next
		}

		if tmp >= 10 {
			carry = 1
			tmp -= 10
		} else {
			carry = 0
		}

		curr.Val = tmp

		if pnode != nil || qnode != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = 1
	}

	return result
}

func main() {
	
}