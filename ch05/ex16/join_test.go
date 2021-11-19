package main

import "testing"

func TestJoin(t *testing.T) {
	testCases := []struct {
		inSep  string
		inVals []string
		want   string
	}{
		{" ", []string{}, ""},
		{" ", []string{"a"}, "a"},
		{" ", []string{"a", "b", "c"}, "a b c"},
		{",", []string{"a", "b", "c"}, "a,b,c"},
		{" = ", []string{"a", "b", "c"}, "a = b = c"},
	}

	for _, tc := range testCases {
		got := join(tc.inSep, tc.inVals...)
		if got != tc.want {
			t.Errorf("join(%q, %v...) = %q, but want %q", tc.inSep, tc.inVals, got, tc.want)
		}
	}
}
