package stack

import "errors"

type Stack struct {
	stack []int
}

var comErr = errors.New("stack is empty")

func NewStack(vals []int) *Stack {
	return &Stack{vals}
}

func (s *Stack) Push(val ...int) {
	s.stack = append(s.stack, val...)
}

func (s *Stack) Pop() (int, error) {
	if !s.Empty() {
		result := s.stack[0]
		s.stack = s.stack[1:]
		return result, nil
	}
	return -1, comErr
}

func (s *Stack) Geek() (int, error) {

	if !s.Empty() {
		result := s.stack[0]
		return result, nil
	}
	return -1, comErr
}

func (s Stack) Empty() bool {
	return len(s.stack) == 0
}
