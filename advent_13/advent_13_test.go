package main

import (
	"reflect"
	"testing"
)

const sampleFilename = "advent_13.sample"

func TestReadAdvent13File(t *testing.T) {
	t.Run("sample file reading", func(t *testing.T) {
		want := advent13File{2, 2, []string{"(800)-555-35-35", "+7-(999)-1-1-1-1-1-1-1"}}
		got := readAdvent13File(sampleFilename)
		assertAdvent13Files(t, got, want)
	})
	t.Run("real file reading (some checks)", func(t *testing.T) {
		gotFile := readAdvent13File(realFilename)

		assertInts(t, gotFile.ascNum, 7493)
		assertInts(t, gotFile.amount, 10000)
		assertInts(t, len(gotFile.phones), 10000)
		assertStrings(t, gotFile.phones[9998], "841(259)61769")
	})
}

func TestCalcAdvent13Result(t *testing.T) {
	gotResult := calcAdvent13Result(readAdvent13File(sampleFilename))
	want := "+7-(999)-1-1-1-1-1-1-1"
	assertStrings(t, gotResult.answer, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %s, want %s", got, want)
	}
}

func assertInts(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}

func assertAdvent13Files(t *testing.T, got advent13File, want advent13File) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
