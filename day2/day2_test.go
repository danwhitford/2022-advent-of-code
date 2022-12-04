package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestGetScore(t *testing.T) {
	table := []struct {
		a   string
		b   string
		out int
	}{
		{"A", "Y", 4},
		{"B", "X", 1},
		{"C", "Z", 7},
	}

	for _, test := range table {
		e := getScore(test.a, test.b)
		utils.Assert(t, test.out, e, test)
	}
}

func TestSolvePart1(t *testing.T) {
	e := solvePart1("testday2.txt")
	utils.Assert(t, 15, e, "testsolveday2")
}

func TestSolvePart2(t *testing.T) {
	e := solvePart2("testday2.txt")
	utils.Assert(t, 12, e, "testsolveday2")
}
