package main

import (
	"math/big"
	"testing"

	"github.com/danwhitford/2022adventofcode/utils"
	"github.com/jinzhu/copier"
)

var monkeys = []Monkey{
	{
		utils.ToBigInts([]int{79, 98}),
		func (i *big.Int) *big.Int { return i.Mul(i, big.NewInt(19)) },
		23,
		2,
		3,
		0,
	},
	{
		utils.ToBigInts([]int{54, 65, 75, 74}),
		func (i *big.Int) *big.Int { return i.Add(i, big.NewInt(6)) },
		19,
		2,
		0,
		0,
	},
	{
		utils.ToBigInts([]int{79, 60, 97}),
		func (i *big.Int) *big.Int { return i.Mul(i, i) },
		13,
		1,
		3,
		0,
	},
	{
		utils.ToBigInts([]int{74}),
		func (i *big.Int) *big.Int { return i.Add(i, big.NewInt(3)) },
		17,
		0,
		1,
		0,
	},
}

func TestSolvePart1(t *testing.T) {
	var m []Monkey
	copier.Copy(&m, &monkeys)
	actual := solvePart1(m)
	utils.Assert(t, 10605, actual, "day 11 part 1")
}

func TestSolvePart2(t *testing.T) {
	var m []Monkey
	copier.Copy(&m, &monkeys)
	actual := solvePart2(m)
	utils.Assert(t, 2713310158, actual, "day 11 part 2")
}
