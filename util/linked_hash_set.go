package util

// Simulate Java LinkedHashSet
type LinkedHashSet[T comparable] struct {
	items  []T
	exists map[T]struct{}
}

func NewLinkedHashSet[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{
		exists: make(map[T]struct{}),
	}
}

func (s *LinkedHashSet[T]) Add(item T) bool {
	if _, ok := s.exists[item]; ok {
		return false
	}
	s.exists[item] = struct{}{}
	s.items = append(s.items, item)
	return true
}

func (s *LinkedHashSet[T]) AddAll(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *LinkedHashSet[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

func (s *LinkedHashSet[T]) Merge(other *LinkedHashSet[T]) {
	for _, item := range other.items {
		s.Add(item)
	}
}
