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
		assertInts(t, len(gotFile.f.rows[0].positions), 100, "columns")
		assertInts(t, len(gotFile.f.rows), 100, "rows")
	})
}

func TestCalcAdvent16Result(t *testing.T) {
	t.Run("calc sample example", func(t *testing.T) {
		gotResult := calcAdvent16Result(readAdvent16File(sampleFilename))
		got := gotResult.answer
		want := 5
		assertInts(t, got, want, "result")
	})
}

func assertAdvent16Files(t *testing.T, got advent16File, want advent16File) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func assertInts(t *testing.T, got int, want int, clue interface{}) {
	if want != got {
		t.Fatalf("Got %d %s, want %d", got, clue, want)
	}
}
