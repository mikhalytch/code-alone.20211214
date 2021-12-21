package main

import (
	"reflect"
	"testing"
)

const sampleFilename = "advent_11.sample"

func TestCalcAdvent11Result(t *testing.T) {
	t.Run("sample (example) calculation", func(t *testing.T) {
		want := advent11Result{23}
		got := calcAdvent11Result(readAdvent11File(sampleFilename))
		assertResults(t, got, want)
	})
}

func assertResults(t *testing.T, got advent11Result, want advent11Result) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func TestReadAdvent11File(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		want := advent11File{
			4,
			[]advent11FileLine{
				{2, []int{2, 3}},
				{0, []int{}},
				{1, []int{4}},
				{0, []int{}},
			},
		}
		got := readAdvent11File(sampleFilename)

		assertAdvent11Files(t, got, want)
	})
	t.Run("real file", func(t *testing.T) {
		gotFile := readAdvent11File(realFilename)
		wantSize := 5000
		assertNumbers(t, gotFile.linesAmount, wantSize)
		assertNumbers(t, len(gotFile.lines), wantSize)
		want4998Line := advent11FileLine{5, []int{696, 1702, 2148, 2654, 2997}}
		assertFileLines(t, gotFile.lines[4998], want4998Line)
	})
}

func assertFileLines(t *testing.T, got advent11FileLine, want advent11FileLine) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v line, want %v", got, want)
	}
}

func assertNumbers(t *testing.T, got int, want int) {
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}

func assertAdvent11Files(t *testing.T, got advent11File, want advent11File) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
