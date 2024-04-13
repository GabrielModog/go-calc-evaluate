package internal

type Stack struct {
	Data []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.Data = append(s.Data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.Data) == 0 {
		return nil, false
	}

	lastIndex := len(s.Data) - 1
	lastVal := s.Data[lastIndex]
	s.Data = s.Data[:lastIndex]

	return lastVal, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.Data) == 0 {
		return nil, false
	}

	return s.Data[len(s.Data)-1], true
}