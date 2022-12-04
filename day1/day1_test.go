package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestSolvePart1(t *testing.T) {
	answer := solvePart1("testday1.txt")
	utils.Assert(t, 24000, answer, "test day 1")
}
func TestSolvePart2(t *testing.T) {
	answer := solvePart2("testday1.txt")
	utils.Assert(t, 45000, answer, "test day 1")
}

