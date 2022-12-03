package utils

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func Split[T any](slice []T, split func(t T) bool) [][]T {
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

func Map[T, U any](slice []T, f func(t T) U) []U {
	out := make([]U, 0)
	for _, t := range slice {
		u := f(t)
		out = append(out, u)
	}
	return out
}

func Assert[U any](t *testing.T, expected, actual interface{}, input U) {
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatalf("failed: %v\n(-want +got):\n%s", input, diff)
	}
}

func FoldN[T any](slice []T, n int) [][]T {
	ret := make([][]T, 0)
	buf := make([]T, 0)
	tick := 0
	for _, el := range slice {
		buf = append(buf, el)
		tick++
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
