package utils

type Set[T comparable] struct {
	Arr map[T]struct{}
}

func SetFromList[T comparable](li []T) Set[T] {
	s := make(map[T]struct{}, 0)
	found := make(map[T]bool, 0)
	for _, el := range li {
		found[el] = true
	}
	for k := range found {
		s[k] = struct{}{}
	}
	return Set[T]{s}
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	o := SetFromList([]T{})
	var smaller, bigger Set[T]
	if s.Len() > other.Len() {
		smaller = other
		bigger = s
	} else {
		smaller = s
		bigger = other
	}

	for v := range smaller.Arr {
		_, prs := bigger.Arr[v]
		if prs {
			o.Arr[v] = struct{}{}
		}
	}

	return o
}

func (s Set[T]) GetOne() T {
	for k := range s.Arr {
		return k
	}
	panic("empty set can't get")
}

func (s Set[T]) Len() int {
	return len(s.Arr)
}

func SetsIntersection[T comparable](sets []Set[T]) Set[T] {
	ret := sets[0]
	for _, set := range sets {
		ret = ret.Intersect(set)
	}
	return ret
}

func (set Set[T]) Empty() bool {
	return len(set.Arr) == 0
}
