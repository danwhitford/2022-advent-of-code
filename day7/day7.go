package main

import (
	"fmt"
	"github.com/danwhitford/2022adventofcode/utils"
	"math"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	directories map[string]*directory
	files       []file
	parent      *directory
}

func (d directory) getSize() int {
	total := 0
	for _, f := range d.files {
		total += f.size
	}
	for _, dd := range d.directories {
		total += dd.getSize()
	}
	return total
}

func parseInput(input []string) *directory {
	root := &directory{
		name:        "/",
		directories: make(map[string]*directory, 0),
		files:       make([]file, 0),
		parent:      nil,
	}
	cdir := root
	for _, line := range input {
		tokens := strings.Fields(line)
		if strings.HasPrefix(line, "$") {
			if tokens[1] == "cd" {
				if tokens[2] == "/" {
					continue
				} else if tokens[2] == ".." {
					cdir = cdir.parent
				} else {
					cdir = cdir.directories[tokens[2]]
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			cdir.directories[tokens[1]] = &directory{
				name:        tokens[1],
				directories: make(map[string]*directory, 0),
				files:       make([]file, 0),
				parent:      cdir,
			}
		} else {
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			f := file{tokens[1], size}
			cdir.files = append(cdir.files, f)
		}
	}
	return root
}

func getDirs(root *directory, out *[]*directory) {
	for _, d := range root.directories {
		*out = append(*out, d)
		getDirs(d, out)
	}
}

func solveDay1(i []string) int {
	root := parseInput(i)
	out := make([]*directory, 0)
	out = append(out, root)
	getDirs(root, &out)
	total := 0
	for _, i := range out {
		size := i.getSize()
		if size <= 100_000 {
			total += size
		}
	}
	return total
}

func solveDay2(i []string) int {
	root := parseInput(i)
	rootSize := root.getSize()
	available := 70000000
	neededFree := 30000000
	unUsedSpace := available - rootSize
	needToDelete := neededFree - unUsedSpace
	smallestEligible := math.MaxInt

	out := make([]*directory, 0)
	out = append(out, root)
	getDirs(root, &out)
	for _, i := range out {
		size := i.getSize()
		if size >= needToDelete && size < smallestEligible {
			smallestEligible = size
		}
	}
	return smallestEligible
}

func main() {
	lines := utils.Lines("day7/day7.txt")
	fmt.Println(solveDay1(lines))
	fmt.Println(solveDay2(lines))
}
