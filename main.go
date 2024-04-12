package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(val interface{}) {
	fmt.Println("[push]:", val)
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	lastIndex := len(s.data) - 1
	lastVal := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	fmt.Println("[pop]: ", lastVal)

	return lastVal, true
}

func main() {
	fmt.Println("[ STACK TEST ]")

	stk := &Stack{}

	stk.Push("First")
	stk.Push("Second")
	stk.Push("Last")

	fmt.Println(stk.data)

	lastValue, isEmpty := stk.Pop()

	if isEmpty {
		fmt.Println("Current stack it's empty!")
	}

	fmt.Println("Last value popped from stack:", lastValue)
}