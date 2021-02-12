package problems

import "github.com/changjunpyo/golang-leetcode/dataStructures"

// Stack is LIFO data structrue
type Stack = dataStructures.Stack

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	s := &Stack{}

	for i := 0; i < len(T); i++ {
		for top, _ := s.Top(); !s.IsEmpty() && T[top] < T[i]; top, _ = s.Top() {
			result[top] = i - top
			s.Pop()
			if s.IsEmpty() {
				break
			}
		}
		s.Push(i)
	}
	return result
}
