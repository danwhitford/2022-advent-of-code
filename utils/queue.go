package utils

import "fmt"

type Queue[T comparable] []T

func (q *Queue[T]) Enqueue(t T) {
	*q = append(*q, t)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(*q) == 0 {
		var t T
		return t, fmt.Errorf("queue is empty")
	}

	v := (*q)[0]
	*q = (*q)[1:]
	return v, nil
}

func (q *Queue[T]) Empty() bool {
	return len(*q) < 1
}
