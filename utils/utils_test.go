package utils

import (
	"testing"

)

func TestFoldN(t *testing.T) {
	table := []struct{
		in []int
		n int
		expected [][]int
	}{
		{
			[]int{1,2,3,4,5,6,7,8,9,10},
			3,
			[][]int{{1,2,3},{4,5,6},{7,8,9},{10}},
		},
	}

	for _, test := range table {
		actual := FoldN(test.in, test.n)
		Assert(t, test.expected, actual, test)
	}
}

func TestUniq(t *testing.T) {
	table := []struct{
		in []rune
		expected []rune
	}{
		{
			[]rune("vJrwpWtwJgWrhcsFMMfFFhFp"),
			[]rune("vJrwpWtghcsFMf"),
		},
	}

	for _, test := range table {
		actual := Uniq(test.in)
		Assert(t, test.expected, actual, test)
	}
}