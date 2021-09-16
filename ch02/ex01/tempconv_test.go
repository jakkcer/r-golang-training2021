package tempconv

import (
	"fmt"
	"testing"
)

func TestCToF(t *testing.T) {
	testInput := BoilingC
	testWant := BoilingF

	descr := fmt.Sprintf("CToF(%q)", testInput)
	if got := CToF(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestCToK(t *testing.T) {
	testInput := BoilingC
	testWant := BoilingK

	descr := fmt.Sprintf("CToK(%q)", testInput)
	if got := CToK(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestFToC(t *testing.T) {
	testInput := BoilingF
	testWant := BoilingC

	descr := fmt.Sprintf("FToC(%q)", testInput)
	if got := FToC(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestFToK(t *testing.T) {
	testInput := BoilingF
	testWant := BoilingK

	descr := fmt.Sprintf("FToK(%q)", testInput)
	if got := FToK(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestKToC(t *testing.T) {
	testInput := BoilingK
	testWant := BoilingC

	descr := fmt.Sprintf("KToC(%q)", testInput)
	if got := KToC(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestKToF(t *testing.T) {
	testInput := BoilingK
	testWant := BoilingF

	descr := fmt.Sprintf("KToF(%q)", testInput)
	if got := KToF(testInput); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}
