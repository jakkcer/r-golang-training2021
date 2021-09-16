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

func BenchmarkPopCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64(testInput)
	}
}

func TestPopCount(t *testing.T) {
	descr := fmt.Sprintf("PopCount(%q)", strconv.FormatUint(testInput, 10))
	if got := PopCount(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}

func TestPopCount64(t *testing.T) {
	descr := fmt.Sprintf("PopCount64(%q)", strconv.FormatUint(testInput, 10))
	if got := PopCount64(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, strconv.Itoa(got), strconv.Itoa(testWant))
	}
}
