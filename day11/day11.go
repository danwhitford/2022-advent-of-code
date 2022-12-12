package main

import (
	"fmt"
	"sort"

	"github.com/danwhitford/2022adventofcode/utils"
	"github.com/jinzhu/copier"
)

type Monkey struct {
	Items           utils.Queue[int]
	Operation       func(int) int
	Test            int
	TrueThrow       int
	FalseThrow      int
	InspectionCount int
}

var myMonkeys = []Monkey{
	{ // Monkey 0
		[]int{84, 72, 58, 51},
		func(i int) int { return i * 3 },
		13,
		1,
		7,
		0,
	},
	{ // Monkey 1
		[]int{88, 58, 58},
		func(i int) int { return i + 8 },
		2,
		7,
		5,
		0,
	},
	{ // Monkey 2
		[]int{93, 82, 71, 77, 83, 53, 71, 89},
		func(i int) int { return i * i },
		7,
		3,
		4,
		0,
	},
	{ // Monkey 3
		[]int{81, 68, 65, 81, 73, 77, 96},
		func(i int) int { return i + 2 },
		17,
		4,
		6,
		0,
	},
	{ // Monkey 4
		[]int{75, 80, 50, 73, 88},
		func(i int) int { return i + 3 },
		5,
		6,
		0,
		0,
	},
	{ // Monkey 5
		[]int{59, 72, 99, 87, 91, 81},
		func(i int) int { return i * 17 },
		11,
		2,
		3,
		0,
	},
	{ // Monkey 6
		[]int{86, 69},
		func(i int) int { return i + 6 },
		3,
		1,
		0,
		0,
	},
	{ // Monkey 7
		[]int{91},
		func(i int) int { return i + 1 },
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
				item /= 3
				// log.Printf("\tWorry level is divided by 3 to %d\n", item)

				if item%monkey.Test == 0 {
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
	var divisor int = 1
	for _, m := range monkeys {
		divisor *= m.Test
	}

	for round := 0; round < 10_000; round++ {
		for i, monkey := range monkeys {
			for !monkey.Items.Empty() {
				item, err := monkey.Items.Dequeue()
				monkey.InspectionCount++
				if err != nil {
					panic(err)
				}
				item = monkey.Operation(item)
				item %= divisor

				if item%monkey.Test == 0 {
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
	var m []Monkey
	copier.Copy(&m, &myMonkeys)
	fmt.Println(solvePart1(m))
	copier.Copy(&m, &myMonkeys)
	fmt.Println(solvePart2(m))
}
