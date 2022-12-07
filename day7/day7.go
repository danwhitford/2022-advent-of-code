package main

import (
	"fmt"
	"github.com/danwhitford/2022adventofcode/utils"
	"github.com/danwhitford/stacko/stack"
	path2 "path"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	directories []string
	files       []file
	parent      string
}

func (d directory) getSize(m map[string]directory) int {
	total := 0
	ds := stack.Stack[directory]{}
	ds.Push(d)
	for !ds.Empty() {
		top, err := ds.Pop()
		if err != nil {
			panic(err)
		}
		for _, f := range top.files {
			total += f.size
		}
		for _, dd := range top.directories {
			ds.Push(m[dd])
		}
		//total += ddd.getSize(m)
	}
	return total
}

func parseInput(input []string) map[string]directory {
	dirs := make(map[string]directory)
	cwd := ""
	var cdir directory
	for _, line := range input {
		tokens := strings.Fields(line)
		if strings.HasPrefix(line, "$") {
			if tokens[1] == "cd" {
				path := path2.Join(cdir.parent, cwd)
				dirs[path] = cdir

				if tokens[2] == ".." {
					cwd = cdir.parent
					cdir = dirs[cwd]
				} else {
					//parent := cwd
					cwd = tokens[2]
					cdir = directory{parent: path}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			dirpath := path2.Join(cdir.parent, cwd, tokens[1])
			cdir.directories = append(cdir.directories, dirpath)
		} else {
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			f := file{tokens[1], size}
			cdir.files = append(cdir.files, f)
		}
	}
	path := path2.Join(cdir.parent, cwd)
	dirs[path] = cdir
	return dirs
}

func solveDay1(i []string) int {
	tree := parseInput(i)
	total := 0
	for k := range tree {
		fmt.Printf("%s\t%+v\t%d\n", k, tree[k], tree[k].getSize(tree))

		d := tree[k]
		if d.getSize(tree) <= 100000 {
			total += d.getSize(tree)
		}
	}
	return total
}

func main() {
	lines := utils.Lines("day7/day7.txt")
	fmt.Println(solveDay1(lines))
}
