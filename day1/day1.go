package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

)

func solveDay1(fname string) int {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	
	var buf int
	elves := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, buf)
			buf = 0
			continue
		}
		lineD, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		buf += lineD
	}
	elves = append(elves, buf)

	sort.Ints(elves)
	var highest = 0
	for _, v := range elves[len(elves)-3:] {
		highest += v
	}

	return highest
}

func main() {
	answer := solveDay1("day1/day1.txt")
	fmt.Println(answer)
}
