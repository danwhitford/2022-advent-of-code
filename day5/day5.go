package main

import (
	"fmt"
	"strings"

	"github.com/danwhitford/2022adventofcode/utils"
	"github.com/danwhitford/stacko/stack"
)

func newStacks() []stack.Stack[rune] {
	return []stack.Stack[rune]{
		{'Q', 'S', 'W', 'C', 'Z', 'V', 'F', 'T'},
		{'Q', 'R', 'B'},
		{'B', 'Z', 'T', 'Q', 'P', 'M', 'S'},
		{'D', 'V', 'F', 'R', 'Q', 'H'},
		{'J', 'G', 'L', 'D', 'B', 'S', 'T', 'P'},
		{'W', 'R', 'T', 'Z'},
		{'H', 'Q', 'M', 'N', 'S', 'F', 'R', 'J'},
		{'R', 'N', 'F', 'H', 'W'},
		{'J', 'Z', 'T', 'Q', 'P', 'R', 'B'},
	}
}

type instruction struct {
	quantity, from, to int
}

func getInstructions(fname string) []instruction {
	lines := utils.Lines(fname)
	moves := utils.Filter(lines, func(s string) bool {
		return strings.HasPrefix(s, "move")
	})
	instructions := utils.Map(moves, func(s string) instruction {
		var q, f, t int
		fmt.Sscanf(s, "move %d from %d to %d", &q, &f, &t)
		f--
		t--
		return instruction{q, f, t}
	})
	return instructions
}

func readTops(stacks []stack.Stack[rune]) string {
	tops := make([]rune, 0)
	for _, s := range stacks {
		v, err := s.Peek()
		if err != nil {
			panic(err)
		}
		tops = append(tops, *v)
	}
	return string(tops)
}

func solvePart1(fname string, stacks []stack.Stack[rune]) string {
	instructions := getInstructions(fname)
	for _, instruction := range instructions {
		for i := 0; i < instruction.quantity; i++ {
			v, err := stacks[instruction.from].Pop()
			if err != nil {
				panic(err)
			}
			stacks[instruction.to].Push(v)
		}
	}

	return readTops(stacks)
}

func solvePart2(fname string, stacks []stack.Stack[rune]) string {
	instructions := getInstructions(fname)
	for _, instruction := range instructions {
		tmp := stack.Stack[rune]{}
		for i := 0; i < instruction.quantity; i++ {
			v, err := stacks[instruction.from].Pop()
			if err != nil {
				panic(err)
			}
			tmp.Push(v)
		}
		for i := 0; i < instruction.quantity; i++ {
			v, err := tmp.Pop()
			if err != nil {
				panic(err)
			}
			stacks[instruction.to].Push(v)
		}
	}

	return readTops(stacks)
}

func main() {
	fmt.Println(solvePart1("day5/day5.txt", newStacks()))
	fmt.Println(solvePart2("day5/day5.txt", newStacks()))
}
