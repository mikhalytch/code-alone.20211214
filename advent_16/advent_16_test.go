package main

import (
	"reflect"
	"testing"
)

const sampleFilename = "advent_16.sample"

func TestReadAdvent16File(t *testing.T) {
	t.Run("read sample file", func(t *testing.T) {
		got := readAdvent16File(sampleFilename)
		want := advent16File{field{[]fieldRow{
			{[]fieldPosition{'.', '.', '.', '.', '.', '.'}},
			{[]fieldPosition{'#', '#', '#', '#', '#', '.'}},
			{[]fieldPosition{'#', 'A', '#', '#', '#', '.'}},
			{[]fieldPosition{'#', '.', '#', 'B', '#', '.'}},
			{[]fieldPosition{'.', '.', '.', '.', '.', '.'}},
		}, point{1, 2}}}

		assertAdvent16Files(t, got, want)
	})
	t.Run("read real file (some checks)", func(t *testing.T) {
		gotFile := readAdvent16File(realFilename)
		assertInts(t, 100, len(gotFile.f.rows[0].positions), "columns")
		assertInts(t, 100, len(gotFile.f.rows), "rows")
	})
}

func assertAdvent16Files(t *testing.T, got advent16File, want advent16File) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func assertInts(t *testing.T, want int, got int, clue interface{}) {
	if want != got {
		t.Fatalf("Got %d %s, want %d", got, clue, want)
	}
}
