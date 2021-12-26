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

func TestGetPossibleMoves(t *testing.T) {
	t.Run("sample field possible moves", func(t *testing.T) {
		tests := []struct {
			name      string
			atPoint   point
			wantMoves []point
		}{
			{"at (5,4)", point{5, 4}, []point{{4, 4}, {5, 3}}},
			{"at (3,4)", point{3, 4}, []point{{2, 4}, {4, 4}, {3, 3}}},
		}
		inputField := readAdvent16File(sampleFilename).f
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := inputField.getPossibleMoves(test.atPoint)
				want := test.wantMoves
				clue := "moves"
				assertSlices(t, got, want, clue)
			})
		}
	})
}

func assertSlices(t *testing.T, got []point, want []point, clue interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v %s, want %v", got, clue, want)
	}
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
