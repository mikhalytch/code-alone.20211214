package main

import (
	"reflect"
	"testing"
)

const sampleFilename = "advent_15.sample"

func TestReadAdvent15File(t *testing.T) {
	t.Run("sample file", func(t *testing.T) {
		got := readAdvent15File(sampleFilename)
		want := advent15File{4, []point{{1, 1}, {3, 3}, {1, 2}, {3, 2}}}
		assertAdvent15Files(t, got, want)
	})
	t.Run("real file (some checks)", func(t *testing.T) {
		gotFile := readAdvent15File(realFilename)
		assertInts(t, gotFile.pointsAmount, 223)
		assertInts(t, len(gotFile.points), 223)
		assertPoints(t, gotFile.points[2], point{48182, 64444})
	})
}

func TestCalcAdvent15Result(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"sample", sampleFilename, 3},
		{"real file calculation (for refactoring)", realFilename, 95590},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calcAdvent15Result(readAdvent15File(test.filename))
			assertInts(t, got.answer, test.want)
		})
	}
}
func TestDistanceTo(t *testing.T) {
	t.Run("point{1,2} to single point", func(t *testing.T) {
		tests := []struct {
			name     string
			p        point
			distance uint64
		}{
			{"(1, 1)", point{1, 1}, 1},
			{"(3, 3)", point{3, 3}, 3},
			{"(1, 2)", point{1, 2}, 0},
			{"(3, 2)", point{3, 2}, 2},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := point{1, 2}.sumDistanceTo(test.p)
				want := test.distance
				assertUints(t, got, want)
			})
		}
	})
	t.Run("single point to multiple points", func(t *testing.T) {
		points := []point{{1, 1}, {3, 3}, {1, 2}, {3, 2}}
		tests := []struct {
			name     string
			p        point
			distance uint64
		}{
			{"(1, 2)", point{1, 2}, 6},
			{"(2, 2)", point{2, 2}, 6},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := test.p.sumDistanceTo(points...)
				want := test.distance
				assertUints(t, got, want)
			})
		}
	})
}

func assertInts(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}
func assertUints(t *testing.T, got uint64, want uint64) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}

func assertAdvent15Files(t *testing.T, got advent15File, want advent15File) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
func assertPoints(t *testing.T, got point, want point) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
