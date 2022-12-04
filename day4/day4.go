package main

import (
	"fmt"

	"github.com/danwhitford/2022adventofcode/utils"
)

type assignment struct {
	low, high int
}

func (a assignment) contains(b assignment) bool {
	return a.low <= b.low && a.high >= b.high
}

func (a assignment) overlaps(b assignment) bool {
	return a.high >= b.low && b.high >= a.low
}

func solvePart1(fname string) int {
	lines := utils.Lines(fname)
	pairs := utils.Map(
		lines,
		func(s string) [2]assignment {
			var a, b, c, d int
			fmt.Sscanf(s, "%d-%d,%d-%d", &a, &b, &c, &d)
			return [2]assignment{
				{a, b},
				{c, d},
			}
		},
	)

	total := 0
	for _, pair := range pairs {
		if pair[0].contains(pair[1]) || pair[1].contains(pair[0]) {
			total++
		}
	}

	return total
}

func solvePart2(fname string) int {
	lines := utils.Lines(fname)
	pairs := utils.Map(
		lines,
		func(s string) [2]assignment {
			var a, b, c, d int
			fmt.Sscanf(s, "%d-%d,%d-%d", &a, &b, &c, &d)
			return [2]assignment{
				{a, b},
				{c, d},
			}
		},
	)

	total := 0
	for _, pair := range pairs {
		if pair[0].overlaps(pair[1]) {
			total++
		}
	}

	return total
}

func main() {
	fmt.Println(solvePart1("day4/day4.txt"))
	fmt.Println(solvePart2("day4/day4.txt"))
}
