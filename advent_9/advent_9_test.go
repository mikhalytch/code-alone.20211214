package main

import (
	"reflect"
	"testing"
)

const (
	sampleFilename = "advent_9.sample"
	realFilename   = "advent_9.test.txt"
)

func TestReadAdvent9File(t *testing.T) {
	t.Run("read sample as expected", func(t *testing.T) {
		want := advent9File{4, []int{3, 1, 2, 4}}
		got := readAdvent9File(sampleFilename)

		assertAdvent9Files(t, got, want)
	})
	t.Run("read real file as expected", func(t *testing.T) {
		gotFile := readAdvent9File(realFilename)

		assertIntegers(t, gotFile.amount, 10000)
		assertIntegers(t, len(gotFile.numbers), 10000)

		m := map[int]bool{}
		for _, n := range gotFile.numbers {
			m[n] = true
		}
		assertIntegers(t, len(m), 10000)
	})
}
func TestCalcAdvent9Result(t *testing.T) {
	t.Run("calculate sample answer", func(t *testing.T) {
		inputFile := readAdvent9File(sampleFilename)
		gotResult := calcAdvent9Result(inputFile)
		assertIntegers(t, gotResult.answer, 3)
		assertIntegers(t, gotResult.amount, 3)
	})
}

func assertIntegers(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func assertAdvent9Files(t *testing.T, got advent9File, want advent9File) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
