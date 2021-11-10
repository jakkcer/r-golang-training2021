package main

import (
	"fmt"
	"testing"
)

func TestNoReturn(t *testing.T) {
	want := fmt.Errorf("no return")
	got := noReturn()
	if got.Error() != want.Error() {
		t.Errorf("noReturn() should return %q, but got %q", want, got)
	}
}
