package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danwhitford/2022adventofcode/utils"
)

type instruction struct {
	Cmd string
	Val int
}

type Part1VM struct {
	X               int
	Cycle           int
	Instructions    []instruction
	IP              int
	Buf             *instruction
	StrengthTracker int
}

func NewPart1Vm(in []instruction) *Part1VM {
	return &Part1VM{
		1,
		1,
		in,
		0,
		nil,
		0,
	}
}

func (vm *Part1VM) Run() {
	for vm.IP < len(vm.Instructions) || vm.Buf != nil {

		if vm.Cycle == 20 || (vm.Cycle-20)%40 == 0 {
			vm.StrengthTracker += (vm.X * vm.Cycle)
		}

		if vm.Buf != nil {
			do := *vm.Buf
			vm.Buf = nil
			vm.X += do.Val
			vm.Cycle++
		} else {
			curr := vm.Instructions[vm.IP]
			switch curr.Cmd {
			case "addx":
				vm.Buf = &curr
				vm.Cycle++
			case "noop":
				vm.Cycle++
			}
			vm.IP++
		}

	}
}

func solvePart1(lines []string) int {
	instructions := parseInput(lines)

	vm := NewPart1Vm(instructions)
	vm.Run()
	return vm.StrengthTracker
}

func parseInput(lines []string) []instruction {
	return utils.Map(
		lines,
		func(line string) instruction {
			var i instruction
			fields := strings.Fields(line)
			i.Cmd = fields[0]
			if len(fields) > 1 {
				d, err := strconv.Atoi(fields[1])
				if err != nil {
					panic(err)
				}
				i.Val = d
			}
			return i
		},
	)
}

type Part2VM struct {
	X            int
	Cycle        int
	Instructions []instruction
	IP           int
	Buf          *instruction
}

func NewPart2Vm(in []instruction) *Part2VM {
	return &Part2VM{
		1,
		0,
		in,
		0,
		nil,
	}
}

func (vm *Part2VM) Run() string {
	var sb strings.Builder
	for vm.IP < len(vm.Instructions) || vm.Buf != nil {
		if vm.Cycle > 0 && vm.Cycle%40 == 0 {
			sb.WriteString("\n")
		}
		willDraw := vm.Cycle % 40
		if utils.Abs(vm.X, willDraw) <= 1 {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}

		if vm.Buf != nil {
			do := *vm.Buf
			vm.Buf = nil
			vm.X += do.Val
			vm.Cycle++
		} else {
			curr := vm.Instructions[vm.IP]
			switch curr.Cmd {
			case "addx":
				vm.Buf = &curr
				vm.Cycle++
			case "noop":
				vm.Cycle++
			}
			vm.IP++
		}

	}
	return sb.String()
}

func solvePart2(lines []string) string {
	instructions := parseInput(lines)
	vm := NewPart2Vm(instructions)
	return vm.Run()
}

func main() {
	lines := utils.Lines("day10/day10.txt")
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
