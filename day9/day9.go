package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danwhitford/2022adventofcode/utils"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type motion struct {
	Dir   direction
	Steps int
}

type coord struct {
	X, Y int
}

func parseMoves(lines []string) []motion {
	moves := make([]motion, 0)
	for _, line := range lines {
		m := motion{}
		fields := strings.Fields(line)
		switch d := fields[0]; d {
		case "R":
			m.Dir = right
		case "L":
			m.Dir = left
		case "D":
			m.Dir = down
		case "U":
			m.Dir = up
		}
		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		m.Steps = steps
		moves = append(moves, m)
	}
	return moves
}

func solvePart1(lines []string) int {
	moves := parseMoves(lines)

	head := coord{0, 0}
	tail := coord{0, 0}
	visits := utils.EmptySet[coord]()
	visits.Add(tail)

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			head.move(move.Dir)
			tail.moveTowards(head)
			visits.Add(tail)
		}
	}

	return visits.Len()
}

func solvePart2(lines []string) int {
	moves := parseMoves(lines)

	knots := make([]coord, 10)
	for i := range knots {
		knots[i] = coord{0, 0}
	}
	visits := utils.EmptySet[coord]()
	visits.Add(knots[9])

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			knots[0].move(move.Dir)
			for j := 1; j < 10; j++ {
				knots[j].moveTowards(knots[j-1])
			}
			visits.Add(knots[9])
		}
	}

	return visits.Len()
}

func (m *coord) move(dir direction) {
	switch dir {
	case up:
		m.Y++
	case down:
		m.Y--
	case right:
		m.X++
	case left:
		m.X--
	}
}

func (m *coord) moveTowards(other coord) {
	if m.X == other.X && m.Y == other.Y {
		return
	} else if m.X == other.X && utils.Abs(m.Y, other.Y) == 2 {
		if m.Y < other.Y {
			m.Y++
		} else {
			m.Y--
		}
	} else if m.Y == other.Y && utils.Abs(m.X, other.X) == 2 {
		if m.X < other.X {
			m.X++
		} else {
			m.X--
		}
	} else if m.X == other.X || m.Y == other.Y {
		return
	} else if !m.touching(other) {
		if m.Y < other.Y {
			m.Y++
		} else {
			m.Y--
		}

		if m.X < other.X {
			m.X++
		} else {
			m.X--
		}
	}
}

func (m *coord) touching(other coord) bool {
	return utils.Abs(m.X, other.X) < 2 && utils.Abs(m.Y, other.Y) < 2
}

func main() {
	lines := utils.Lines("day9/day9.txt")
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
