package main

import (
	"github.com/danwhitford/2022adventofcode/utils"
	"testing"
)

var testInput = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestSolvePart1(t *testing.T) {
	table := []struct {
		in       []string
		expected int
	}{
		{testInput, 21},
	}
	for _, tst := range table {
		actual := solvePart1(tst.in)
		utils.Assert(t, tst.expected, actual, "testday8part1")
	}
}
func TestSolvePart2(t *testing.T) {
	table := []struct {
		in       []string
		expected int
	}{
		{testInput, 8},
	}
	for _, tst := range table {
		actual := solvePart2(tst.in)
		utils.Assert(t, tst.expected, actual, "testday8part2")
	}
}
