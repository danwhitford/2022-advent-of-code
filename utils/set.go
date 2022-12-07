package utils

type Set[T comparable] struct {
	Arr map[T]struct{}
}

func EmptySet[T comparable]() Set[T] {
	return SetFromList(make([]T, 0))
}

func SetFromList[T comparable](li []T) Set[T] {
	s := make(map[T]struct{}, 0)
	for _, el := range li {
		s[el] = struct{}{}
	}
	return Set[T]{s}
}

func (set *Set[T]) Intersect(other Set[T]) Set[T] {
	o := SetFromList([]T{})
	var smaller, bigger Set[T]
	if set.Len() > other.Len() {
		smaller = other
		bigger = *set
	} else {
		smaller = *set
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

func (set *Set[T]) GetOne() T {
	for k := range set.Arr {
		return k
	}
	panic("empty set can't get")
}

func (set *Set[T]) Len() int {
	return len(set.Arr)
}

func SetsIntersection[T comparable](sets []Set[T]) Set[T] {
	ret := sets[0]
	for _, set := range sets {
		ret = ret.Intersect(set)
	}
	return ret
}

func (set *Set[T]) Empty() bool {
	return len(set.Arr) == 0
}

func (set *Set[T]) Add(t T) {
	set.Arr[t] = struct{}{}
}

func (set *Set[T]) Contains(t T) bool {
	_, prs := set.Arr[t]
	return prs
}
