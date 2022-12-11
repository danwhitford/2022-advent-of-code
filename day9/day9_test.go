package main

import (
	"strings"
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

var in = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var in1 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestSolvePart1(t *testing.T) {
	actual := solvePart1(strings.Split(in, "\n"))
	utils.Assert(t, 13, actual, "test solve day 9 part 1")
}

func TestSolvePart2(t *testing.T) {
	ts := []struct {
		in   string
		want int
	}{
		{in, 1},
		{in1, 36},
	}

	for _, tst := range ts {
		actual := solvePart2(strings.Split(tst.in, "\n"))
		utils.Assert(t, tst.want, actual, tst)
	}
}
