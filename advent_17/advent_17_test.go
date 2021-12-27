package main

import "testing"

const sampleFilename = "advent_17.sample"

func TestCalcAdvent17Result(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		tests := []struct {
			name     string
			filename string
			want     int
		}{
			{"sample", sampleFilename, 2},
			{"sample.1", "advent_17.1.sample", 1},
			{"sample.2", "advent_17.2.sample", 0},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := calcAdvent17Result(readAdvent17File(test.filename))
				want := test.want
				assertInts(t, got.answer, want, "answer")
			})
		}
	})
}

func assertInts(t *testing.T, got int, want int, clue interface{}) {
	if got != want {
		t.Fatalf("Got %d %s, want %d", got, clue, want)
	}
}
