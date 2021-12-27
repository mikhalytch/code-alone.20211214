package main

import "testing"

func TestCalcAdvent1(t *testing.T) {
	want := advent1Result{3}
	filename := "advent_1_test.sample"
	input := readAdvent1File(filename)
	got := calcAdvent1(input)

	assertCalcAdvent1Results(t, got, want)
}

func assertCalcAdvent1Results(t *testing.T, got advent1Result, want advent1Result) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
