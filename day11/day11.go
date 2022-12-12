package main

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/danwhitford/2022adventofcode/utils"
)

type Monkey struct {
	Items           utils.Queue[*big.Int]
	Operation       func(*big.Int) *big.Int
	Test            int64
	TrueThrow       int
	FalseThrow      int
	InspectionCount int
}

var myMonkeys = []Monkey{
	{ // Monkey 0
		utils.ToBigInts([]int{84, 72, 58, 51}),
		func(i *big.Int) *big.Int { return i.Mul(i, big.NewInt(3)) },
		13,
		1,
		7,
		0,
	},
	{ // Monkey 1
		utils.ToBigInts([]int{88, 58, 58}),
		func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(8)) },
		2,
		7,
		5,
		0,
	},
	{ // Monkey 2
		utils.ToBigInts([]int{93, 82, 71, 77, 83, 53, 71, 89}),
		func(i *big.Int) *big.Int { return i.Mul(i, i) },
		7,
		3,
		4,
		0,
	},
	{ // Monkey 3
		utils.ToBigInts([]int{81, 68, 65, 81, 73, 77, 96}),
		func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(2)) },
		17,
		4,
		6,
		0,
	},
	{ // Monkey 4
		utils.ToBigInts([]int{75, 80, 50, 73, 88}),
		func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(3)) },
		5,
		6,
		0,
		0,
	},
	{ // Monkey 5
		utils.ToBigInts([]int{59, 72, 99, 87, 91, 81}),
		func(i *big.Int) *big.Int { return i.Mul(i, big.NewInt(17)) },
		11,
		2,
		3,
		0,
	},
	{ // Monkey 6
		utils.ToBigInts([]int{86, 69}),
		func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(6)) },
		3,
		1,
		0,
		0,
	},
	{ // Monkey 7
		utils.ToBigInts([]int{91}),
		func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(1)) },
		19,
		2,
		5,
		0,
	},
}

func solvePart1(monkeys []Monkey) int {
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for !monkey.Items.Empty() {
				item, err := monkey.Items.Dequeue()
				monkey.InspectionCount++
				// log.Printf("Monkey %d inspects item %d\n", i, item)
				if err != nil {
					panic(err)
				}
				item = monkey.Operation(item)
				// log.Printf("\tItem is changed to %d\n", item)
				item.Div(item, big.NewInt(3))
				// log.Printf("\tWorry level is divided by 3 to %d\n", item)
				var m big.Int
				m.Set(item)
				if m.Mod(item, big.NewInt(monkey.Test)).Cmp(big.NewInt(0)) == 0 {
					// log.Printf("\tWorry level is divisible by %d\n", monkey.Test)
					// log.Printf("\tItem with worry level %d is thrown to monkey %d\n", item, monkey.TrueThrow)
					monkeys[monkey.TrueThrow].Items.Enqueue(item)
				} else {
					// log.Printf("\tWorry level is not divisible by %d\n", monkey.Test)
					// log.Printf("\tItem with worry level %d is thrown to monkey %d\n", item, monkey.FalseThrow)
					monkeys[monkey.FalseThrow].Items.Enqueue(item)
				}
			}
			monkeys[i] = monkey
		}
	}

	for i, m := range monkeys {
		fmt.Println("Monkey", i, "holding", m.Items)
	}
	for i, m := range monkeys {
		fmt.Println("Monkey", i, "inspected items", m.InspectionCount, "times")
	}
	business := make([]int, 0)
	for _, m := range monkeys {
		business = append(business, m.InspectionCount)
	}
	sort.Ints(business)
	l := len(business)
	return business[l-1] * business[l-2]
}

func solvePart2(monkeys []Monkey) int {
	var divisor int64 = 1
	for _, m := range monkeys {
		divisor *= m.Test
	}

	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for !monkey.Items.Empty() {
				monkey.InspectionCount++
				item, err := monkey.Items.Dequeue()
				if err != nil {
					panic(err)
				}
				item = monkey.Operation(item)				

				item.Mod(item, big.NewInt(divisor))
				// Division test
				var m big.Int
				m.Mod(item, big.NewInt(monkey.Test))			
				if m.Cmp(big.NewInt(0)) == 0 {
					monkeys[monkey.TrueThrow].Items.Enqueue(item)
				} else {
					monkeys[monkey.FalseThrow].Items.Enqueue(item)
				}
			}
			monkeys[i] = monkey
		}

		fmt.Println("After round", round)
		for i, m := range monkeys {
			fmt.Println("Monkey", i, "inspected items", m.InspectionCount, "times")
		}
	}

	business := make([]int, 0)
	for _, m := range monkeys {
		business = append(business, m.InspectionCount)
	}
	sort.Ints(business)
	l := len(business)
	return business[l-1] * business[l-2]
}

func main() {
	fmt.Println(solvePart1(myMonkeys))
}
