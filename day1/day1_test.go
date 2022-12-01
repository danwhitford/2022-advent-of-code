package main

import "testing"

func TestSolve1(t *testing.T) {
	answer := solveDay1("testday1.txt")
	if answer != 45000 {
		t.Fatalf("wanted 45000 but got %d", answer)
	}
}
