package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input string

		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			input: "",

			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 0,
		},
		{
			input: "a",

			counts:  map[rune]int{'a': 1},
			utflen:  []int{0, 1, 0, 0, 0},
			invalid: 0,
		},
		{
			input: "あ",

			counts:  map[rune]int{'あ': 1},
			utflen:  []int{0, 0, 0, 1, 0},
			invalid: 0,
		},
		{
			input: "日本",

			counts:  map[rune]int{'日': 1, '本': 1},
			utflen:  []int{0, 0, 0, 2, 0},
			invalid: 0,
		},
		{
			input: "あa",

			counts:  map[rune]int{'あ': 1, 'a': 1},
			utflen:  []int{0, 1, 0, 1, 0},
			invalid: 0,
		},
	}

	for _, test := range tests {
		counts, utflen, invalid := charCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("%q counts: got %v, want %v", test.input, counts, test.counts)
		}
		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("%q utflen: got %v, want %v", test.input, utflen, test.utflen)
		}
		if invalid != test.invalid {
			t.Errorf("%q invalid: got %v, want %v", test.input, invalid, test.invalid)
		}
	}
}
