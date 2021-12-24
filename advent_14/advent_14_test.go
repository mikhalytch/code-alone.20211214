package main

import (
	"reflect"
	"testing"
)

const sampleFilename = "advent_14.sample"

func TestReadAdvent14File(t *testing.T) {
	t.Run("read sample", func(t *testing.T) {
		got := readAdvent14File(sampleFilename)
		want := advent14File{2, 2, []simpleRestriction{{1, 'A'}, {2, 'Z'}}}

		assertAdvent14Files(t, got, want)
	})
	t.Run("read real file (some checks)", func(t *testing.T) {
		gotFile := readAdvent14File(realFilename)
		assertInts(t, gotFile.codeLength, 100)
		assertInts(t, gotFile.rulesAmount, 777)
		assertInts(t, len(gotFile.rules), 777)
	})
}

func TestCalcAdvent14Result(t *testing.T) {
	t.Run("sample example calculation", func(t *testing.T) {
		gotResult := calcAdvent14Result(readAdvent14File(sampleFilename))
		want := "NM"
		got := gotResult.answer
		assertStrings(t, got, want)
	})
}

func TestCalcMedian(t *testing.T) {
	tests := []struct {
		name string
		set  []rune
		want rune
	}{
		{"B-Z : N",
			[]rune{'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
			'N'},
		{"A-Y : M",
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y'},
			'M'},
		{"A-D : C", []rune{'A', 'B', 'C', 'D'}, 'B'},
		{"A-E : C", []rune{'A', 'B', 'C', 'D', 'E'}, 'C'},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calcMedian(test.set)
			want := test.want
			assertRunes(t, got, want)
		})
	}
}

func TestCreateLimitedSet(t *testing.T) {
	tests := []struct {
		name       string
		restricted []rune
		want       []rune
	}{
		{"full", []rune{},
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}},
		{"full w/ nil", nil,
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}},
		{"minus A", []rune{'A'},
			[]rune{'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}},
		{"minus Z", []rune{'Z'},
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y'}},
		{"minus J & O", []rune{'J', 'O'},
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := createLimitedSet(test.restricted)
			want := test.want
			assertRuneSlices(t, got, want)
		})
	}
}

func assertRuneSlices(t *testing.T, got []rune, want []rune) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func assertRunes(t *testing.T, got rune, want rune) {
	if got != want {
		t.Fatalf("Got %s, want %s", string(got), string(want))
	}
}

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func assertInts(t *testing.T, got int, want int) {
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}

func assertAdvent14Files(t *testing.T, got advent14File, want advent14File) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
