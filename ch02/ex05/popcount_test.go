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

func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(testInput)
	}
}

func TestPopCount(t *testing.T) {
	descr := fmt.Sprintf("PopCount(%q)", strconv.FormatUint(testInput, 10))
	if got := PopCount(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}

func TestClearPopCount(t *testing.T) {
	descr := fmt.Sprintf("ClearPopCount(%q)", strconv.FormatUint(testInput, 10))
	if got := ClearPopCount(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}
