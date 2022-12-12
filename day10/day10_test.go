package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestRunVm(t *testing.T) {
	instructions := []instruction{
		{Cmd: "noop"},
		{"addx", 3},
		{"addx", -5},
	}

	vm := NewPart1Vm(instructions)
	vm.Run()
	actual := vm.X
	utils.Assert(t, -1, actual, "test mini program")
}

func TestSolvePart1(t *testing.T) {
	lines := utils.Lines("testday10.txt")
	actual := solvePart1(lines)
	utils.Assert(t, 13140, actual, "day10part1")
}

func TestSolvePart2(t *testing.T) {
	lines := utils.Lines("testday10.txt")
	actual := solvePart2(lines)
	expected := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	utils.Assert(t, expected, actual, "day10part1")
}
