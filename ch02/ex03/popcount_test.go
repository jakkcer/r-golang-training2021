package popcount

import (
	"fmt"
	"strconv"
	"testing"
)

var testInput uint64 = 926
var testWant int = 7

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testInput)
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(testInput)
	}
}

func TestPopCount(t *testing.T) {
	descr := fmt.Sprintf("PopCount(%q)", strconv.FormatUint(testInput, 10))
	if got := PopCount(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}

func TestLoopPopCount(t *testing.T) {
	descr := fmt.Sprintf("LoopPopCount(%q)", strconv.FormatUint(testInput, 10))
	if got := LoopPopCount(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}
