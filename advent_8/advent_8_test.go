package main

import (
	"reflect"
	"testing"
)

func TestCalcAdvent8Result(t *testing.T) {
	filename := "advent_8.sample"
	want := "BE"
	inputFile := readAdvent8File(filename)
	gotResult := calcAdvent8Result(inputFile)

	assertStrings(t, gotResult.value, want)
}

func TestCreateCodeVariants(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		variants []string
	}{
		{"sample variant", "2233", []string{"AADD", "AAE", "BDD", "BE"}},
		{"Kevin's variant", "222", []string{"AAA", "AB", "BA", "C"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.variants
			got := createCodeVariants(test.in)

			assertSlices(t, got, want)
		})
	}
}

func assertSlices(t *testing.T, got []string, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestLexicographicallyMinimalString(t *testing.T) {
	tests := []struct {
		name     string
		variants []string
		want     string
	}{
		{"Kevin's variant", []string{"AC", "CA"}, "AC"},
		{"Kevin's non-told variant", []string{"AC", "C"}, "C"},
		{"sample variant", []string{"BE", "AADD"}, "BE"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lexicographicallyMinimalString(test.variants)
			want := test.want
			assertStrings(t, got, want)
		})
	}
}

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
