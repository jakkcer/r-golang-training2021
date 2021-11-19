package main

import (
	"fmt"
	"testing"
)

// トポロジカルな順序になっていることを確認
func TestTopoSort(t *testing.T) {
	descr := fmt.Sprintf("topoSort(testInput)")
	testInput := prereqs

	got := topoSort(testInput)
	seen := make(map[string]bool, len(got))
	for _, g := range got {
		// 一度しか現れていないことを確認
		if seen[g] {
			t.Errorf("%s: %q appears twice", descr, g)
		}
		seen[g] = true
		// 事前条件がすでに登場していることを確認
		if v, ok := testInput[g]; ok {
			for key := range v {
				if !seen[key] {
					t.Errorf("%s: %q is not topological!", descr, g)
				}
			}
		}
	}
}
