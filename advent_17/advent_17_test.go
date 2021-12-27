package main

import (
	"fmt"
	"testing"
)

const sampleFilename = "advent_17.sample"

func TestCalcAdvent17Result(t *testing.T) {
	t.Run("files", func(t *testing.T) {
		tests := []struct {
			name     string
			filename string
			want     int
		}{
			{"sample", sampleFilename, 2},
			{"sample.1", "advent_17.1.sample", 1},
			{"sample.2", "advent_17.2.sample", 0},
			{"sample.2", "advent_17.3.sample", 0},
			{"sample.2", "advent_17.4.sample", 0},
			{"sample.2", realFilename, 1},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := calcAdvent17Result(readAdvent17File(test.filename))
				want := test.want
				assertInts(t, got.answer, want, "answer")
			})
		}
	})
	t.Run("real file calculation (based on IDE analysis)", func(t *testing.T) {
		got := calcAdvent17Result(readAdvent17File(realFilename))
		want := 1
		assertInts(t, got.answer, want, "answer")
	})
}

func TestReadAdvent17File(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantLength int
	}{
		{"sample", sampleFilename, 16},
		{"real file", realFilename, 100000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := readAdvent17File(test.filename)
			assertInts(t, len(got.symbols), test.wantLength, "length")
		})
	}
}

func TestAppendRunes(t *testing.T) {
	parens := []paren{'(', '{', '}', '[', '{', '{'}
	stack := createEmptyParenStack()
	calc := &zeroLevelBracesCalculator{}
	for _, p := range parens {
		stack = stack.add(p, calc)
	}
	got := fmt.Sprint(stack)
	want := "\"([{{\""
	clue := interface{}("stack")

	assertStrings(t, got, want, clue)
}

func assertStrings(t *testing.T, got string, want string, clue interface{}) {
	if got != want {
		t.Fatalf("Got %q %s, want %q", got, clue, want)
	}
}

func assertInts(t *testing.T, got int, want int, clue interface{}) {
	if got != want {
		t.Fatalf("Got %d %s, want %d", got, clue, want)
	}
}
