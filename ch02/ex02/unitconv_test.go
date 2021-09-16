package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestConvTempature(t *testing.T) {
	var testInput float64 = 100
	testWant := "100°F = 37.77777777777778°C, 100°C = 212°F\n"

	descr := fmt.Sprintf("convTempature(%q)", strconv.FormatFloat(testInput, 'f', -1, 64))
	out = new(bytes.Buffer)
	convTempature(testInput)
	if got := out.(*bytes.Buffer).String(); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestConvLength(t *testing.T) {
	var testInput float64 = 100
	testWant := "100ft = 30.478512648582747m, 100m = 328.1ft\n"
	descr := fmt.Sprintf("convLength(%q)", strconv.FormatFloat(testInput, 'f', -1, 64))
	out = new(bytes.Buffer)
	convLength(testInput)
	if got := out.(*bytes.Buffer).String(); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestConvWeight(t *testing.T) {
	var testInput float64 = 100
	testWant := "100lb = 45.35147392290249kg, 100kg = 220.5lb\n"

	descr := fmt.Sprintf("convWeight(%q)", strconv.FormatFloat(testInput, 'f', -1, 64))
	out = new(bytes.Buffer)
	convWeight(testInput)
	if got := out.(*bytes.Buffer).String(); got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}
