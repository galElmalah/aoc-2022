package set

import "fmt"

type Set[T comparable] struct {
	values map[T]bool
}

func NewSet[T comparable](values ...T) *Set[T] {
	m := make(map[T]bool, len(values))
	for _, v := range values {
		m[v] = true
	}
	return &Set[T]{
		values: m,
	}
}

func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.values[v] = true
	}
}

func (s *Set[T]) Remove(values ...T) {
	for _, v := range values {
		delete(s.values, v)
	}
}

func (s *Set[T]) Has(values ...T) bool {
	for _, v := range values {
		_, ok := s.values[v]
		if !ok {
			return false
		}
	}
	return true
}

// set2# OMIT
// set3 OMIT

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet(s.Values()...)
	for _, v := range other.Values() {
		if !result.Has(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	// pass smaller set first for optimization
	if s.Size() < other.Size() {
		return intersect(s, other)
	}
	return intersect(other, s)
}

// set3# OMIT
// set4 OMIT

// intersect returns intersection of given sets. It iterates over smaller set for optimization.
func intersect[T comparable](smaller, bigger *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k, _ := range smaller.values {
		if bigger.Has(k) {
			result.Add(k)
		}
	}
	return result
}

func (s *Set[T]) Values() []T {
	return s.toSlice()
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Clear() {
	s.values = map[T]bool{}
}

func (s *Set[T]) String() string {
	return fmt.Sprint(s.toSlice())
}

func (s *Set[T]) toSlice() []T {
	result := make([]T, 0, len(s.values))
	for k := range s.values {
		result = append(result, k)
	}
	return result
}

// set5# OMIT
// set-usage OMIT

func ExampleSet() {
	s1 := NewSet(4, 4, -8, 15)
	s2 := NewSet("foo", "foo", "bar", "baz")
	fmt.Println(s1.Size(), s2.Size()) // 3, 3

	s1.Add(-16)
	s2.Add("hoge")
	fmt.Println(s1.Size(), s2.Size())        // 4, 4
	fmt.Println(s1.Has(-16), s2.Has("hoge")) // true, true

	s1.Remove(15)
	s2.Remove("baz")
	fmt.Println(s1.Size(), s2.Size()) // 3, 3

	fmt.Println(len(s1.Values()), len(s2.Values())) // 3, 3

	s3 := NewSet("hoge", "dragon", "fly")
	fmt.Println(s2.Union(s3).Size()) // 5
	fmt.Println(s2.Intersect(s3))    // [hoge]

	s1.Clear()
	s2.Clear()
	fmt.Println(s1.Size(), s2.Size()) // 0, 0
}

// set-usage# OMIT
