package main

import (
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
)

func TestContains(t *testing.T) {
	table := []struct {
		a        assignment
		b        assignment
		expected bool
	}{
		{
			assignment{2, 4},
			assignment{6, 8},
			false,
		},
		{
			assignment{2, 3},
			assignment{4, 5},
			false,
		},
		{
			assignment{5, 7},
			assignment{7, 9},
			false,
		},
		{
			assignment{2, 8},
			assignment{3, 7},
			true,
		},
		{
			assignment{6, 6},
			assignment{4, 6},
			true,
		},
		{
			assignment{2, 6},
			assignment{4, 8},
			false,
		},
	}

	for _, test := range table {
		actual := test.a.contains(test.b) || test.b.contains(test.a)
		utils.Assert(t, test.expected, actual, test)
	}
}

func TestOverlap(t *testing.T) {
	table := []struct {
		a        assignment
		b        assignment
		expected bool
	}{
		{
			assignment{2, 4},
			assignment{6, 8},
			false,
		},
		{
			assignment{2, 3},
			assignment{4, 5},
			false,
		},
		{
			assignment{5, 7},
			assignment{7, 9},
			true,
		},
		{
			assignment{2, 8},
			assignment{3, 7},
			true,
		},
		{
			assignment{6, 6},
			assignment{4, 6},
			true,
		},
		{
			assignment{2, 6},
			assignment{4, 8},
			true,
		},
		{
			assignment{1, 6},
			assignment{6, 10},
			true,
		},
		{
			assignment{6, 10},
			assignment{1, 6},
			true,
		},
	}

	for _, test := range table {
		actual := test.a.overlaps(test.b)
		utils.Assert(t, test.expected, actual, test)
	}
}

func TestSolvePart1(t *testing.T) {
	actual := solvePart1("testday4.txt")
	utils.Assert(t, 2, actual, "testday4")
}

func TestSolvePart2(t *testing.T) {
	actual := solvePart2("testday4.txt")
	utils.Assert(t, 4, actual, "testday4")
}
