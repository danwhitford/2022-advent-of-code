package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/danwhitford/2022adventofcode/utils"
)

func solvePart1(fname string) int {
	intLines := utils.Map(
		utils.Lines(fname),
		func(s string) int {
			if s == "" {
				return -1
			}
			d, err := strconv.Atoi(s)
			utils.TestErr(err)
			return d
		})

	totals := utils.Split(
		intLines,
		func(t int) bool {
			return t == -1
		},
	)
	elves := utils.Map(
		totals,
		func(t []int) int {
			total := 0
			for _, v := range t {
				total += v
			}
			return total
		},
	)
	highest := 0
	for _, v := range elves {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func solvePart2(fname string) int {
	intLines := utils.Map(
		utils.Lines(fname),
		func(s string) int {
			if s == "" {
				return -1
			}
			d, err := strconv.Atoi(s)
			utils.TestErr(err)
			return d
		})

	totals := utils.Split(
		intLines,
		func(t int) bool {
			return t == -1
		},
	)

	elves := utils.Map(
		totals,
		func(t []int) int {
			total := 0
			for _, v := range t {
				total += v
			}
			return total
		},
	)

	sort.Ints(elves)
	var highest = 0
	for _, v := range elves[len(elves)-3:] {
		highest += v
	}

	return highest
}

func main() {
	fmt.Println(solvePart1("day1/day1.txt"))
	fmt.Println(solvePart2("day1/day1.txt"))
}
