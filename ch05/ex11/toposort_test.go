package main

import (
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	// トポロジカルな順序になっていることを確認
	topologicalTestInput := map[string]map[string]string{
		"a": {"b": "", "c": ""},
		"b": {"c": "", "d": ""},
		"c": {"e": ""},
	}
	descr := fmt.Sprintf("topoSort(topologicalTestInput)")

	got, err := topoSort(topologicalTestInput)
	if err != nil {
		t.Errorf("%s is not topological test case", descr)
	}
	seen := make(map[string]bool, len(got))
	for _, g := range got {
		// 一度しか現れていないことを確認
		if seen[g] {
			t.Errorf("%s: %q appears twice", descr, g)
		}
		seen[g] = true
		// 事前条件がすでに登場していることを確認
		if v, ok := topologicalTestInput[g]; ok {
			for key := range v {
				if !seen[key] {
					t.Errorf("%s: %q is not topological!", descr, g)
				}
			}
		}
	}

	// ループが発生する場合のテスト
	loopTestInput := map[string]map[string]string{
		"a": {"b": ""},
		"b": {"a": ""},
	}
	wantErrorMsg1 := fmt.Errorf("loop occurred: %q", []string{"a", "b", "a"}).Error()
	wantErrorMsg2 := fmt.Errorf("loop occurred: %q", []string{"b", "a", "b"}).Error()
	descr = fmt.Sprintf("topoSort(loopTestInput)")

	got, err = topoSort(loopTestInput)
	errMsg := err.Error()
	if errMsg != wantErrorMsg1 && errMsg != wantErrorMsg2 {
		t.Errorf("%s didn't return correct error: %v", descr, errMsg)
	}
}
