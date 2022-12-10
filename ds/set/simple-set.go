package set

type SimpleSet[T comparable] struct {
	items map[T]bool
}

func NewSimpleSet[T comparable](values ...T) *SimpleSet[T] {
	m := make(map[T]bool)

	return &SimpleSet[T]{
		items: m,
	}
}

func (s *SimpleSet[T]) Add(key T) {
	s.items[key] = true
}

func (s *SimpleSet[T]) Has(key T) bool {
	_, ok := s.items[key]
	return ok
}

func (s *SimpleSet[T]) Size() int {
	return len(s.items)
}
