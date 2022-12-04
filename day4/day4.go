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

func main() {
	lines := utils.Lines("day4/day4.txt")
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
		if pair[0].overlaps(pair[1]) || pair[1].contains(pair[0]) {
			total++
		}
	}

	fmt.Println(total)
}
