package intset

import (
	"testing"
)

func newMapIntSet(a ...int) *MapIntSet {
	mapIntSet := &MapIntSet{}
	for _, x := range a {
		mapIntSet.Add(x)
	}
	return mapIntSet
}

func TestMapIntSetHas(t *testing.T) {
	tests := []struct {
		mapIntSet *MapIntSet
		expected  map[int]bool
	}{
		{
			mapIntSet: newMapIntSet(1, 2, 3, 10),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				5:  false,
				9:  false,
				10: true,
				11: false,
			},
		},
	}
	for _, test := range tests {
		for x, expected := range test.expected {
			if got := test.mapIntSet.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestMapIntSetAdd(t *testing.T) {
	tests := []struct {
		mapIntSet *MapIntSet
		adds      []int
		expected  map[int]bool
	}{
		{
			mapIntSet: newMapIntSet(1, 2, 3, 10),
			adds:      []int{0, 5, 100},
			expected: map[int]bool{
				0:   true,
				1:   true,
				2:   true,
				3:   true,
				4:   false,
				5:   true,
				9:   false,
				10:  true,
				11:  false,
				99:  false,
				100: true,
				101: false,
			},
		},
	}

	for _, test := range tests {
		for _, x := range test.adds {
			test.mapIntSet.Add(x)
		}
		for x, expected := range test.expected {
			if got := test.mapIntSet.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestMapIntSetUnionWith(t *testing.T) {
	tests := []struct {
		a, b     *MapIntSet
		expected map[int]bool
	}{
		{
			a: newMapIntSet(1, 2, 3, 10),
			b: newMapIntSet(50, 55),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				9:  false,
				10: true,
				11: false,
				49: false,
				50: true,
				51: false,
				54: false,
				55: true,
				56: false,
			},
		},
		{
			a: newMapIntSet(1, 2, 3, 10),
			b: newMapIntSet(),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				9:  false,
				10: true,
				11: false,
			},
		},
		{
			a: newMapIntSet(),
			b: newMapIntSet(),
			expected: map[int]bool{
				0:     false,
				1:     false,
				2:     false,
				10:    false,
				100:   false,
				1000:  false,
				10000: false,
			},
		},
	}

	for _, test := range tests {
		test.a.UnionWith(test.b)
		for x, expected := range test.expected {
			if got := test.a.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestMapIntSetString(t *testing.T) {
	tests := []struct {
		mapIntSet *MapIntSet
		expected  string
	}{
		{
			mapIntSet: newMapIntSet(),
			expected:  "{}",
		},
		{
			mapIntSet: newMapIntSet(1),
			expected:  "{1}",
		},
		{
			mapIntSet: newMapIntSet(1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144),
			expected:  "{1 2 3 5 8 13 21 34 55 89 144}",
		},
	}

	for _, test := range tests {
		if got := test.mapIntSet.String(); got != test.expected {
			t.Errorf("mapIntSet.String() expected: %q, but got: %q", test.expected, got)
		}
	}
}
