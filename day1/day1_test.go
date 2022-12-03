package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestSolve1(t *testing.T) {
	answer := solveDay1("testday1.txt")
	utils.Assert(t, 45000, answer, "test day 1")
}
