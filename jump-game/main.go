package main

import (
	"fmt"
)

type node struct {
	idx int
	val int
}

type stack []node

func (s stack) Push(n node) stack {
	return append(s, n)
}

func (s stack) Pop() (stack, node) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func canJump(nums []int) bool {
	// edge case
	if len(nums) == 1 {
		return true
	}
	s := make(stack, 0)
	s = s.Push(node{0, nums[0]})
	var currnode node
	for len(s) > 0 {
		// Depth first search for end node
		s, currnode = s.Pop()
		for i := 1; i <= currnode.val; i++ {
			if currnode.idx+i >= len(nums)-1 {
				return true
			}
			s = s.Push(node{currnode.idx + i, nums[currnode.idx+i]})
		}
	}
	return false
}

func main() {
	// a := []int{2, 3, 1, 1, 4}
	// fmt.Println("[2,3,1,1,4] -> ", canJump(a))
	// b := []int{3, 2, 1, 0, 4}
	// fmt.Println("[3,2,1,0,4] -> ", canJump(b))
	c := []int{0}
	fmt.Println("[0] -> ", canJump(c))
	// s := make(stack, 0)
	// s = s.Push(node{0, 2})
	// s = s.Push(node{1, 3})
	// s = s.Push(node{2, 1})

	// s, p := s.Pop()
	// fmt.Println(p)
}
