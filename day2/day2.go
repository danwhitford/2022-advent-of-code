package main

import (
	"fmt"
	"strings"

	"github.com/danwhitford/2022adventofcode/utils"
)

type move int

const (
	Rock move = iota
	Paper
	Scissors
)

var scoreFer = map[move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func getMove(s string) move {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("not found" + s)
	}
}

func beatingMove(m move) move {
	switch m {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("what is this")
	}
}

func losingMove(m move) move {
	switch m {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("what is this")
	}
}

func getScore(a, b string) int {
	theirMove := getMove(a)

	var outcome int
	if b == "Z" {
		outcome += 6
		yourMove := beatingMove(theirMove)
		outcome += scoreFer[yourMove]
	} else if b == "Y" {
		outcome += 3
		yourMove := theirMove
		outcome += scoreFer[yourMove]
	} else {
		yourMove := losingMove(theirMove)
		outcome += scoreFer[yourMove]
	}

	return outcome
}

func solvePart1(fname string) int {
	lines := utils.Lines(fname)
	turns := utils.Map(
		lines,
		func(s string) []string {
			return strings.Split(s, " ")
		},
	)
	total := 0
	for _, turn := range turns {
		theirs := getMove(turn[0])
		ours := getMove(turn[1])
		total += scoreFer[ours]
		if theirs == ours {
			total += (3)
		} else if (ours == Rock && theirs == Scissors) ||
			(ours == Scissors && theirs == Paper) ||
			(ours == Paper && theirs == Rock) {
			total += 6
		}
	}
	return total
}

func solvePart2(fname string) int {
	lines := utils.Lines(fname)
	turns := utils.Map(
		lines,
		func(s string) []string {
			return strings.Split(s, " ")
		},
	)

	var total int
	for _, turn := range turns {
		score := getScore(turn[0], turn[1])
		total += score
	}
	return total
}

func main() {
	fmt.Println(solvePart1("day2/day2.txt"))
	fmt.Println(solvePart2("day2/day2.txt"))
}
