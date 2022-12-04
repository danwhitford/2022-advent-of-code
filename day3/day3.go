package main

import (
	"fmt"

	"github.com/danwhitford/2022adventofcode/utils"
)

func getMisplaced(rucksacks []string) string {
	rucksacksets := make([]utils.Set[rune], 0)
	for _, sack := range rucksacks {
		rucksacksets = append(rucksacksets, utils.SetFromList([]rune(sack)))
	}

	intersect := utils.SetsIntersection(rucksacksets)
	if !intersect.Empty() {
		return string(intersect.GetOne())
	}

	panic("not found")
}

func getPriority(s string) int {
	b := s[0]
	if 'a' <= b && b <= 'z' {
		return int(b%'a') + 1
	} else {
		return int(b%'A') + 27
	}
}

func solvePart1(fname string) int {
	lines := utils.Lines(fname)
	total := 0
	for _, line := range lines {
		firstHalf := utils.SetFromList([]rune(line)[len(line)/2:])
		secondHalf := utils.SetFromList([]rune(line)[:len(line)/2])
		intersection := firstHalf.Intersect(secondHalf)
		total += getPriority(string(intersection.GetOne()))
	}
	return total
}

func solvePart2(fname string) int {
	lines := utils.Lines(fname)
	folded := utils.FoldN(lines, 3)
	total := 0
	for _, line := range folded {
		missing := getMisplaced(line)
		fmt.Printf("missing was %s\n", missing)
		priority := getPriority(missing)
		total += priority
	}
	return total
}

func main() {
	fmt.Println(solvePart1("day3/day3.txt"))
	fmt.Println(solvePart2("day3/day3.txt"))
}
