package main

import "testing"

func TestIsFit(t *testing.T) {
	tests := []struct {
		name            string
		digitsAmount    int
		valueInQuestion int64
		want            bool
	}{
		{"100 not fit 2 digits", 2, 100, false},
		{"100 fit 3 digits", 3, 100, true},
		{"95 fit 2 digits", 2, 95, true},
		{"14 digits false low corner", 14, 9999999999999, false},
		{"14 digits true low corner", 14, 10000000000000, true},
		{"14 digits true high corner", 14, 99999999999999, true},
		{"14 digits false high corner", 14, 100000000000000, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertBool(t, test.want, newDigitsAmountChecker(test.digitsAmount).isFit(test.valueInQuestion))
		})
	}
}

func assertBool(t *testing.T, want bool, got bool) {
	if want != got {
		t.Errorf("Got %v, want %v", got, want)
	}
}
