package main

import (
	"fmt"
	"github.com/danwhitford/2022adventofcode/utils"
	"strconv"
)

type coord struct {
	x, y, val int
}

func makeGrid(in []string) ([]int, int) {
	grid := make([]int, 0)
	var h, w int
	for _, line := range in {
		for _, r := range line {
			i, err := strconv.Atoi(string(r))
			utils.TestErr(err)
			grid = append(grid, i)
		}
		w = len(line)
		h++
	}
	if h != w {
		panic("not square")
	}
	return grid, h
}

func solvePart1(in []string) int {
	grid, size := makeGrid(in)
	visibles := map[string]utils.Set[coord]{
		"left":   utils.EmptySet[coord](),
		"right":  utils.EmptySet[coord](),
		"top":    utils.EmptySet[coord](),
		"bottom": utils.EmptySet[coord](),
	}

	// From left
	for y := 0; y < size; y++ {
		highestTree := -1
		for x := 0; x < size; x++ {
			tree := GetAt(grid, x, y, size)
			if tree > highestTree {
				u := visibles["left"]
				u.Add(coord{x, y, tree})
				highestTree = tree
			}
		}
	}

	// From top
	for x := 0; x < size; x++ {
		highestTree := -1
		for y := 0; y < size; y++ {
			tree := GetAt(grid, x, y, size)
			if tree > highestTree {
				u := visibles["top"]
				u.Add(coord{x, y, tree})
			}
			if tree > highestTree {
				highestTree = tree
			}
		}
	}

	// From right
	for y := 0; y < size; y++ {
		highestTree := -1
		for x := size - 1; x >= 0; x-- {
			tree := GetAt(grid, x, y, size)
			if tree > highestTree {
				u := visibles["right"]
				u.Add(coord{x, y, tree})
			}
			if tree > highestTree {
				highestTree = tree
			}
		}
	}

	// From bottom
	for x := 0; x < size; x++ {
		highestTree := -1
		for y := size - 1; y >= 0; y-- {
			tree := GetAt(grid, x, y, size)
			if tree > highestTree {
				u := visibles["bottom"]
				u.Add(coord{x, y, tree})
			}
			if tree > highestTree {
				highestTree = tree
			}
		}
	}

	merged := visibles["left"]
	for _, v := range visibles {
		merged = merged.Union(v)
	}

	return merged.Len()
}

func solvePart2(in []string) int {
	grid, size := makeGrid(in)
	best := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			tree := GetAt(grid, x, y, size)
			var xx, yy int

			// Go left
			xx, yy = x-1, y
			leftScore := 0
			for {
				otherTree := GetAt(grid, xx, yy, size)
				if otherTree == 0 {
					break
				}
				leftScore++
				if otherTree >= tree {
					break
				}
				xx--
			}

			// Go down
			xx, yy = x, y+1
			downScore := 0
			for {
				otherTree := GetAt(grid, xx, yy, size)
				if otherTree == 0 {
					break
				}
				downScore++
				if otherTree >= tree {
					break
				}
				yy++
			}

			// Go right
			xx, yy = x+1, y
			rightScore := 0
			for {
				otherTree := GetAt(grid, xx, yy, size)
				if otherTree == 0 {
					break
				}
				rightScore++
				if otherTree >= tree {
					break
				}
				xx++
			}

			// Go up
			xx, yy = x, y-1
			upScore := 0
			for {
				otherTree := GetAt(grid, xx, yy, size)
				if otherTree == 0 {
					break
				}
				upScore++
				if otherTree >= tree {
					break
				}
				yy--
			}

			score := upScore * downScore * leftScore * rightScore
			fmt.Printf("{%d, %d} %d (%d, %d, %d, %d)\n", x, y, score, upScore, downScore, leftScore, rightScore)
			if score > best {
				best = score
			}
		}
	}
	return best
}

func GetAt(arr []int, x, y, size int) int {
	if x < 0 || y < 0 || x > size-1 || y > size-1 {
		return 0
	}
	i := (y * size) + x
	v := arr[i]
	return v
}

func main() {
	lines := utils.Lines("day8/day8.txt")
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines)) // 169344 too low
}
