package main

import (
	"reflect"
	"testing"
)

func TestReadAdvent5File(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     advent5File
	}{
		{"reading sample file", "advent_5.sample", advent5File{
			5, 5, [][]rune{{'1'}, {'2'}, {'3', '6', '9'}, {'4'}, {'5', '6'}},
		}},
		{"real life file", "advent_5.test.txt", advent5File{
			10, 1234567,
			[][]rune{
				{'1', '2', '3', '4', '6', '7', '8', '9'},
				{'3', '5', '7', '8', '9'},
				{'1', '2', '5', '7', '8', '9'},
				{'3', '4', '5', '6', '7', '8'},
				{'1', '3', '5', '7', '9'},
				{'3', '5', '6', '7'},
				{'2', '5', '6', '7', '8', '9'},
				{'9'},
				{'4', '8'},
				{'2', '5', '6', '7'},
			},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := readAdvent5File(test.filename)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Got %v, wanted %v", got, test.want)
			}
		})
	}
}

func TestCalcAdvent5Result(t *testing.T) {
	filename := "advent_5.sample"
	want := advent5Result{"12945", "12945"}
	got := calcAdvent5Result(readAdvent5File(filename))
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
