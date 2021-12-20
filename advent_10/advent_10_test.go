package main

import (
	"reflect"
	"testing"
)

const (
	sampleFilename = "advent_10.sample"
	realFilename   = "advent_10.test.txt"
)

func TestCalcAdvent10Result(t *testing.T) {
	t.Run("sample file", func(t *testing.T) {
		want := 8
		gotResult := calcAdvent10Result(readAdvent10File(sampleFilename))

		got := gotResult.answer
		assertNumbers(t, got, want)
	})
}

func assertNumbers(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
func TestReadAdvent10File(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		want := advent10File{7, []int{2, 1, 4, 5, 1, 3, 3}}
		got := readAdvent10File(sampleFilename)
		assertAdvent10Files(t, got, want)
	})
	t.Run("real file", func(t *testing.T) {
		want := 1000
		gotFile := readAdvent10File(realFilename)

		assertNumbers(t, gotFile.amount, want)
		assertNumbers(t, len(gotFile.sizes), want)
	})
}

func assertAdvent10Files(t *testing.T, got advent10File, want advent10File) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
