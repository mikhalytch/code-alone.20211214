package main

import "testing"

const filename = "advent_2_test.sample"

func TestReadAdvent2File(t *testing.T) {
	want := advent2File{4, 1}
	got := readAdvent2File(filename)

	assertAdvent2Files(t, got, want)
}

func assertAdvent2Files(t *testing.T, got advent2File, want advent2File) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestCalcAdvent2(t *testing.T) {
	want := advent2Result{6}
	got := calcAdvent2(readAdvent2File(filename))

	assertAdvent2Results(t, got, want)
}

func assertAdvent2Results(t *testing.T, got advent2Result, want advent2Result) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"gcd of 5 and 1 = 1", args{5, 1}, 1},
		{"gcd of 12 and 16 = 4", args{12, 16}, 4},
		{"gcd of 261426 and 1630068 = 15378", args{261426, 1630068}, 15378},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
