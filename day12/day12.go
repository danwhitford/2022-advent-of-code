package main

import (
	"fmt"
	"math"

	"github.com/danwhitford/2022adventofcode/utils"
)

type Board struct {
	b             []byte
	width, height int
}

type Coord [2]int

type QueueNode struct {
	coord Coord
	depth int
}

func (c Coord) X() int {
	return c[0]
}

func (c Coord) Y() int {
	return c[1]
}

func (board Board) GetAt(x, y int) byte {
	if x < 0 || y < 0 || x > board.width-1 || y > board.width-1 {
		panic("out of bounds")
	}
	i := (y * board.width) + x
	v := board.b[i]
	return v
}

func (board Board) GetDestinations(coord Coord) []Coord {
	sourceHeight := board.GetAt(coord.X(), coord.Y())
	destinations := make([]Coord, 0)
	// left
	if coord.X() > 0 {
		v := board.GetAt(coord.X()-1, coord.Y())
		if v <= sourceHeight+1 {
			destinations = append(destinations, Coord{coord.X() - 1, coord.Y()})
		}
	}
	// right
	if coord.X()+1 < board.width {
		v := board.GetAt(coord.X()+1, coord.Y())
		if v <= sourceHeight+1 {
			destinations = append(destinations, Coord{coord.X() + 1, coord.Y()})
		}
	}
	// up
	if coord.Y() > 0 {
		v := board.GetAt(coord.X(), coord.Y()-1)
		if v <= sourceHeight+1 {
			destinations = append(destinations, Coord{coord.X(), coord.Y() - 1})
		}
	}
	// down
	if coord.Y()+1 < board.height {
		v := board.GetAt(coord.X(), coord.Y()+1)
		if v <= sourceHeight+1 {
			destinations = append(destinations, Coord{coord.X(), coord.Y() + 1})
		}
	}

	return destinations
}

func shortestPath(board Board, startPos, endPos Coord) int {
	seen := utils.EmptySet[Coord]()
	q := make(utils.Queue[QueueNode], 0)
	q.Enqueue(QueueNode{startPos, 0})
	seen.Add(startPos)
	for !q.Empty() {
		n, err := q.Dequeue()
		if err != nil {
			panic(err)
		}
		if n.coord == endPos {
			return n.depth
		}
		for _, destination := range board.GetDestinations(n.coord) {
			if !seen.Contains(destination) {
				q.Enqueue(QueueNode{destination, n.depth + 1})
				seen.Add(destination)
			}
		}
	}

	return -1
}

func solvePart1(input []string) int {
	var board Board
	var startPos, endPos Coord
	board.height = len(input)
	for y, line := range input {
		board.width = len(line)
		for x, b := range []byte(line) {
			if b == 'S' {
				startPos = Coord{x, y}
				board.b = append(board.b, 'a')
			} else if b == 'E' {
				endPos = Coord{x, y}
				board.b = append(board.b, 'z')
			} else {
				board.b = append(board.b, b)
			}
		}
	}
	return shortestPath(board, startPos, endPos)
}

func solvePart2(input []string) int {
	var board Board
	var startPos []Coord
	var endPos Coord
	board.height = len(input)
	for y, line := range input {
		board.width = len(line)
		for x, b := range []byte(line) {
			if b == 'S' || b == 'a' {
				startPos = append(startPos, Coord{x, y})
				board.b = append(board.b, 'a')
			} else if b == 'E' {
				endPos = Coord{x, y}
				board.b = append(board.b, 'z')
			} else {
				board.b = append(board.b, b)
			}
		}
	}

	min := math.MaxInt
	for _, pos := range startPos {
		path := shortestPath(board, pos, endPos)
		if path > 0 && path < min {
			min = path
		}
	}
	return min
}

func main() {
	lines := utils.Lines("day12/day12.txt")
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
