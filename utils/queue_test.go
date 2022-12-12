package utils

import (
	"testing"
)

func TestQueuu(t *testing.T) {
	var q Queue[int] = []int{1, 2}
	q.Enqueue(3)

	v, err := q.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	Assert(t, 1, v, 1)
	v, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	Assert(t, 2, v, 2)
	v, err = q.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	Assert(t, 3, v, 3)

	Assert(t, true, q.Empty(), "is empty")
}
