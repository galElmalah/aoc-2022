package stack

type Stack[T comparable] struct {
	items []T
}

func (s *Stack[T]) Push(item T) int {
	s.items = append(s.items, item)
	return s.Size()
}

func (s *Stack[T]) Pop() T {
	item := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return item
}

func (s *Stack[T]) Peek() T {
	return s.items[s.Size()-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
