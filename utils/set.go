package utils

type Set[T comparable] struct {
	Arr []T
}

func SetFromList[T comparable](li []T) Set[T] {
	s := make([]T, 0)
	found := make(map[T]bool, 0)
	for _, el := range li {
		found[el] = true
	}
	for k := range found {
		s = append(s, k)
	}
	return Set[T]{s}
}

func SetsIntersection[T comparable](sets []Set[T]) Set[T] {
	found := make(map[T]int, 0)
	for _, set := range sets {
		for _, el := range set.Arr {
			found[el]++
		}
	}
	intersected := make([]T, 0)
	for k, v := range found {
		if v == len(sets) {
			intersected = append(intersected, k)
		}
	}
	return Set[T]{intersected}
}

func (set Set[T]) Empty() bool {
	return len(set.Arr) == 0
}
