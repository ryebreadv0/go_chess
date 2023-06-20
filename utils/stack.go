package utils

type Stack[N any] struct {
	data []N
}

func NewStack[N any]() *Stack[N] {
	return &Stack[N]{make([]N, 0)}
}

func (s *Stack[N]) Push(x N) {
	s.data = append(s.data, x)
}

func (s *Stack[N]) Pop() N {
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}