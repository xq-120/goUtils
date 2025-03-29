package goutils

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Count() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) PushItems(items []T) {
	s.elements = append(s.elements, items...)
}
