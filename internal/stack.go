package internal

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	lastIndex := len(s.data) - 1
	lastVal := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return lastVal, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	return s.data[len(s.data)-1], true
}