package main

import (
	"reflect"
	"testing"
)

func TestReadAdvent7File(t *testing.T) {
	t.Run("sample read fully", func(t *testing.T) {
		filename := "advent_7.sample"
		want := advent7File{3,
			[]advent7FileLine{{2, []int64{2, 3}}, {0, []int64{}}, {0, []int64{}}}}
		got := readAdvent7File(filename)

		assertAdvent7Files(t, got, want)
	})
	t.Run("real-life parts are read as expected", func(t *testing.T) {
		filename := "advent_7.test.txt"
		gotFile := readAdvent7File(filename)

		assertInt64(t, gotFile.linesAmt, 50000)
		assertInt(t, len(gotFile.lines), 50000)
		assertSlices(t, gotFile.lines[49998],
			advent7FileLine{7, []int64{49510, 812, 4299, 11633, 14751, 4812, 13208}})
	})
}

func assertSlices(t *testing.T, got advent7FileLine, want advent7FileLine) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func assertAdvent7Files(t *testing.T, got advent7File, want advent7File) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestCalcAdvent7Result(t *testing.T) {
	filename := "advent_7.sample"
	want := int64(5)
	inputFile := readAdvent7File(filename)
	gotResult := calcAdvent7File(inputFile)

	got := gotResult.all
	assertInt64(t, got, want)
}

func assertInt(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func assertInt64(t *testing.T, got int64, want int64) {
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
