package internal

type StackValue interface {
	int | int64 | float32 | float64 | string
}

type Stack[T StackValue] struct {
	Data []T
}

func (s *Stack[T]) Push(val T) {
	s.Data = append(s.Data, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Data) == 0 {
		var emptyValue T
		return emptyValue, false
	}

	lastIndex := len(s.Data) - 1
	lastVal := s.Data[lastIndex]
	s.Data = s.Data[:lastIndex]

	return lastVal, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.Data) == 0 {
		var emptyValue T
		return emptyValue, false
	}

	return s.Data[len(s.Data)-1], true
}