package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s, sep string
		want   int
	}{
		{"", "", 0},
		{"", " ", 1},
		{" ", "", 1},
		{" ", " ", 2},
		{"1", "1", 2},
		{"a:b:c", ":", 3},
		{"a b c d e", " ", 5},
		{"abc123abc123abc123", "123", 4},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
