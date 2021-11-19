package main

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	// reverseCaseを使ってテストする
	tests := []struct {
		input string
		want  string
	}{
		{"Hello, $hoge", "Hello, HOGE"},
		{"Hello, $HOGE", "Hello, hoge"},
		{"Hello, $HoGe", "Hello, hOgE"},
		{"Hello, $hoge $hoge $fuga", "Hello, HOGE HOGE FUGA"},
		{"Hello, $hoge\n$fuga $PIYO $PIYO,  $HOGEhoge", "Hello, HOGE\nFUGA piyo piyo,  hogeHOGE"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("expand(%q, reverseCase)", test.input)
		got := expand(test.input, reverseCase)
		if got != test.want {
			t.Errorf("%s = %q, but want %q", descr, got, test.want)
		}
	}
}
