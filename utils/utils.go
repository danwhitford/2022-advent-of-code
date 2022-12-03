package utils

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Lines(fname string) []string {
	f, err := os.Open(fname)
	TestErr(err)
	scanner := bufio.NewScanner(f)

	elves := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		elves = append(elves, line)
	}
	
	return elves
}

func Split[T any] (slice[]T, split func(t T) bool) [][]T {
	elves := make([][]T, 0)
	buf := make([]T, 0)
	for _, v := range slice {
		if split(v) {
			elves = append(elves, buf)
			buf = make([]T, 0)
			continue
		}
		buf = append(buf, v)
	}
	elves = append(elves, buf)
	return elves
}

func Map[T, U any] (slice []T, f func(t T) U) []U {
	out := make([]U, 0)
	for _, t := range slice {
		u := f(t)
		out = append(out, u)
	}
	return out
}

func Assert[T comparable, U any] (t *testing.T, expected, actual T, input U) {
	if actual != expected {
		t.Fatalf("wanted %v but got %v for %v", expected, actual, input)
	}
}

func FoldN[T any](slice []T, n int) [][]T {
	ret := make([][]T, 0)
	buf := make([]T, 0)
	tick := 0
	for _, el := range slice {
		buf = append(buf, el)
		if tick == 3 {
			ret = append(ret, buf)
			tick = 0
			buf = make([]T, 0)
		}
	}
	if len(buf) > 0 {
		ret = append(ret, buf)
	}

	return ret
}