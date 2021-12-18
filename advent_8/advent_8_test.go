package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadAdvent8File(t *testing.T) {
	filename := "advent_8.sample"
	want := "2233"
	got := readAdvent8File(filename)

	assertStrings(t, string(got), want)
}

func TestCalcAdvent8Result(t *testing.T) {
	t.Run("sample file", func(t *testing.T) {
		filename := "advent_8.sample"
		want := "BE"
		inputFile := readAdvent8File(filename)
		gotResult := calcAdvent8Result(inputFile)

		assertStrings(t, gotResult.value, want)
	})
	t.Run("check real file by calculating to code, then back", func(t *testing.T) {
		filename := "advent_8.test.txt"
		inputFile := readAdvent8File(filename)
		gotResult := calcAdvent8Result(inputFile)
		rev := func(res string) string {
			r := make([]rune, 0)
			for _, s := range res {
				var a string
				switch s {
				case 'A':
					a = "2"
				case 'B':
					a = "22"
				case 'C':
					a = "222"
				case 'D':
					a = "3"
				case 'E':
					a = "33"
				case 'F':
					a = "333"
				case 'G':
					a = "4"
				case 'H':
					a = "44"
				case 'I':
					a = "444"
				case 'J':
					a = "5"
				case 'K':
					a = "55"
				case 'L':
					a = "555"
				case 'M':
					a = "6"
				case 'N':
					a = "66"
				case 'O':
					a = "666"
				case 'P':
					a = "7"
				case 'Q':
					a = "77"
				case 'R':
					a = "777"
				case 'S':
					a = "7777"
				case 'T':
					a = "8"
				case 'U':
					a = "88"
				case 'V':
					a = "888"
				case 'W':
					a = "9"
				case 'X':
					a = "99"
				case 'Y':
					a = "999"
				case 'Z':
					a = "9999"
				}
				r = append(r, bytes.Runes([]byte(a))...)
			}
			return string(r)
		}
		reverse := rev(gotResult.value)
		assertStrings(t, reverse, string(inputFile))
	})
}

func TestCreateCodeVariants(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		variants []string
	}{
		{"sample variant", "2233", []string{ /*"AADD", "AAE", "BDD",*/ "BE"}},
		{"Kevin's variant", "222", []string{ /*"AAA", "AB", "BA",*/ "C"}},
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
