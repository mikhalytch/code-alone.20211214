package main

import "testing"

func TestCalcAdvent3(t *testing.T) {
	want := advent3Result{5, 1}
	filename := "advent_3_test.test"
	input := readAdvent3File(filename)
	got := calcAdvent3(input)
	assertAdvent3Results(t, got, want)
}

func assertAdvent3Results(t *testing.T, got advent3Result, want advent3Result) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
