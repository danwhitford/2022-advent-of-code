package main

import (
	"strings"
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

const testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestSolvePart1(t *testing.T) {
	lines := strings.Split(testInput, "\n")
	actual := solvePart1(lines)
	utils.Assert(t, 31, actual, "day 12 part 1")
}

func TestSolvePart2(t *testing.T) {
	lines := strings.Split(testInput, "\n")
	actual := solvePart2(lines)
	utils.Assert(t, 29, actual, "day 12 part 2")
}
