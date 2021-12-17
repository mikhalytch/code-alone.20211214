package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReadAdvent6File(t *testing.T) {
	t.Run("sample can be read as expected", func(t *testing.T) {
		filename := "advent_6.sample"
		want := []rune{'L', 'R', 'L', 'R', 'R'}
		wantString := "LRLRR"
		gotFile := readAdvent6File(filename)

		assertRuneArray(t, want, gotFile.directions)
		asResult := advent6Result(gotFile)
		assertStrings(t, asResult.print(), wantString)
	})
	t.Run("actual file can be read without changes", func(t *testing.T) {
		filename := "advent_6.test.txt"
		buf, err := os.ReadFile(filename)
		if err != nil {
			t.Fatal(fmt.Errorf("unable to read %s: %w", filename, err))
		}
		want := string(buf)
		got := readAdvent6File(filename)
		result := advent6Result(got)

		assertStrings(t, result.print(), want)
	})
}

func TestCalcAdvent6Result(t *testing.T) {
	filename := "advent_6.sample"
	want := "R"
	inputFile := readAdvent6File(filename)
	result := calcAdvent6Result(&inputFile)
	got := result.print()

	assertStrings(t, got, want)
}

func TestGetAggregatedDirections(t *testing.T) {
	t.Run("real life value printed as expected", func(t *testing.T) {
		value := directionsAggregator(-21 + 9i)
		want := bytes.Runes([]byte("DDDDDDDDDRRRRRRRRRRRRRRRRRRRRR"))
		got := value.getAggregatedDirections()
		assertRuneArray(t, want, got)
	})
}

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}

func assertRuneArray(t *testing.T, want []rune, got []rune) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
