package main

import "testing"

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
