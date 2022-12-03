package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestGetMisplaces(t *testing.T) {
	table := []struct {
		in       []string
		expected string
	}{
		{
			[]string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			"r",
		},
		{
			[]string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			"Z",
		},
	}

	for _, test := range table {
		actual := getMisplaced(test.in)
		utils.Assert(t, test.expected, actual, test)
	}
}

func TestGetPriority(t *testing.T) {
	table := []struct {
		in       string
		expected int
	}{
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
		{"t", 20},
		{"s", 19},
		{"r", 18},
		{"Z", 52},
	}

	for _, test := range table {
		actual := getPriority(test.in)
		utils.Assert(t, test.expected, actual, test)
	}
}

func TestSolveDay3(t *testing.T) {
	actual := solveDay3("testday3.txt")
	utils.Assert(t, 70, actual, "testday3")
}
