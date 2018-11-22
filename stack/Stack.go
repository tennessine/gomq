package stack

import "fmt"

type Node struct {
	value int32
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.head.value, true
}

func (s *Stack) Push(v int32) {
	s.head = &Node{v, s.head}
	s.size++
}

func (s *Stack) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	v := s.head.value
	s.head = s.head.next
	s.size--
	return v, true
}

func (s *Stack) Print() {
	temp := s.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}
