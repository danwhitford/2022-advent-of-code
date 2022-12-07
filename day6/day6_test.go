package main

import (
	"github.com/danwhitford/2022adventofcode/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	table := []struct {
		in       string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range table {
		actual := solvePart1(test.in)
		utils.Assert(t, test.expected, actual, test)
	}
}

func TestSolvePart2(t *testing.T) {
	table := []struct {
		in       string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range table {
		actual := solvePart2(test.in)
		utils.Assert(t, test.expected, actual, test)
	}
}
