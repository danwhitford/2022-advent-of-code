package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
	"github.com/danwhitford/stacko/stack"
)

func TestSolvePart1(t *testing.T) {
	var testStacks []stack.Stack[rune] = []stack.Stack[rune]{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}

	actual := solvePart1("day5test.txt", testStacks)
	utils.Assert(t, "CMZ", actual, "testday5part1")
}

func TestSolvePart2(t *testing.T) {
	var testStacks []stack.Stack[rune] = []stack.Stack[rune]{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	
	actual := solvePart2("day5test.txt", testStacks)
	utils.Assert(t, "MCD", actual, "testday5part2")
}
