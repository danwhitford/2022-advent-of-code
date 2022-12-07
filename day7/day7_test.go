package main

import (
	"github.com/danwhitford/2022adventofcode/utils"
	"strings"
	"testing"
)

var input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestSolveDay1(t *testing.T) {
	actual := solveDay1(strings.Split(input, "\n"))
	utils.Assert(t, 95437, actual, "testpart1")
}
