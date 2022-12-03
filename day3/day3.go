package main

import (
	"fmt"

	"github.com/danwhitford/2022adventofcode/utils"
)

func getMisplaced(items string) string {
	l := len(items)
	a := items[l/2:]
	b := items[:l/2]
	found := make(map[rune]int, 0)
	for _, r := range a {
		found[r] = 1
	}
	for _, r := range b {
		_, prs := found[r]
		if prs {
			return string(r)
		}
	}

	panic("not found")
}

func getPriority(s string) int {
	b := s[0]
	if 'a' <= b && b <= 'z' {
		return int(b % 'a') + 1
	} else {
		return int(b % 'A') + 27
	}
}

func solveDay3(fname string) int {
	lines := utils.Lines(fname)
	total := 0
	for _, line := range lines {
		missing := getMisplaced(line)
		priority := getPriority(missing)
		total += priority
	}
	return total
}

func main() {
	fmt.Println(solveDay3("day3/day3.txt"))
}