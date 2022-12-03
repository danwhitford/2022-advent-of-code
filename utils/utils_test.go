package utils

import (
	"encoding/json"
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
		expectedJson, _ := json.Marshal(test.expected)
		actualJson, _ := json.Marshal(actual)
		Assert(t, string(expectedJson), string(actualJson), test)
	}
}