package reverse

import (
	"testing"
)

func TestRevUTF8(t *testing.T) {
	s := []byte("あいうえお")
	got := string(reverseUTF8(s))
	want := "おえういあ"
	if got != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
