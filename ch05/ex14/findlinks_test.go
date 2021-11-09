package main

import "testing"

func TestCountClass(t *testing.T) {
	want := map[string]int{
		"data structures":       5,
		"computer organization": 3,
		"discrete math":         2,
		"operating systems":     1,
		"linear algebra":        1,
		"formal languages":      1,
		"intro to programming":  1,
	}

	got := countClass()
	for gotKey, gotVal := range got {
		if wantVal, ok := want[gotKey]; !ok {
			t.Errorf("got count for %q, but don't want", gotKey)
		} else {
			if gotVal != wantVal {
				t.Errorf("%q appears %d times, but want to be %d times", gotKey, gotVal, wantVal)
			}
		}
	}
}
