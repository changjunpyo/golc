package dataStructures

import "errors"

type Stack struct {
	val  []int
	size int
}

func (s *Stack) Push(item int) {
	s.val = append(s.val[:s.size], item)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	s.size--

	return s.val[s.size], nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	return s.val[s.size-1], nil
}
