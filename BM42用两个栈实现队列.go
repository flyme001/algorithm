package main

import (
	"fmt"
)

type Stack struct {
	element []int
}

func NewStack(col int) *Stack {
	return &Stack{
		element: make([]int, 0, col),
	}
}

func (s Stack) Len() int {
	return len(s.element)
}

func (s *Stack) Push(elem int) {
	s.element = append(s.element, elem)
}

func (s *Stack) Pop() int {
	if len(s.element) == 0 {
		return 0
	}

	fmt.Println("--2--", s.element)
	tmp := s.element[len(s.element)-1]
	// 巧妙的地方在这里 这里相当于每次递减取一个值
	s.element = s.element[:len(s.element)-1]
	fmt.Println("----3----", s.element)
	return tmp
}

// 输出6 5
func main() {
	stack := NewStack(10)

	stack.Push(5)
	stack.Push(6)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
