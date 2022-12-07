package main

import (
	"fmt"
	"github.com/danwhitford/2022adventofcode/utils"
	"io"
	"os"
)

func solve(input string, n int) int {
	runes := []rune(input)
	for i := range runes {
		if i < n {
			continue
		}
		s := utils.SetFromList(runes[i-n : i])
		if s.Len() == n {
			return i
		}
	}
	return -1
}

func solvePart1(input string) int {
	return solve(input, 4)
}

func solvePart2(input string) int {
	return solve(input, 14)
}

func main() {
	f, err := os.Open("day6/day6.txt")
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePart1(string(bytes)))
	fmt.Println(solvePart2(string(bytes)))
}
