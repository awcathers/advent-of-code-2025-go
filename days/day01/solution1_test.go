package main

import "testing"

func TestSolve(t *testing.T) {
	input := "test_input"
	got := Solve(input)
	want := 3

	if got != want {
		t.Fatalf("Solve(%q) = %d, want %d", input, got, want)
	}
}
